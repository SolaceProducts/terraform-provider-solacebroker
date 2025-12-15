// terraform-provider-solacebroker
//
// Copyright 2024 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package broker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-solacebroker/internal/semp"
)

type brokerResource brokerEntity[schema.Schema]

const (
	defaults                        = "defaults"
	defaultObjectName               = "default"
	minRequiredBrokerSempApiVersion = "2.33" // Shipped with broker version 10.3
)

var (
	ErrDeleteSingletonOrDefaultsNotAllowed = errors.New("deleting singleton or default objects are not allowed from the broker")
	BrokerPlatformName                     = map[string]string{
		"VMR":       "Software Event Broker",
		"Appliance": "Appliance",
	}
)

var (
	_ resource.ResourceWithConfigure        = &brokerResource{}
	_ resource.ResourceWithConfigValidators = &brokerResource{}
	_ resource.ResourceWithImportState      = &brokerResource{}
	_ resource.ResourceWithUpgradeState     = &brokerResource{}
)

var (
	skipApiCheck      = false
	apiAlreadyChecked = false
	lock              sync.Mutex
)

func newBrokerResource(inputs EntityInputs) brokerEntity[schema.Schema] {
	return newBrokerEntity(inputs, true)
}

func newBrokerResourceGenerator(inputs EntityInputs) func() resource.Resource {
	return newBrokerResourceClosure(newBrokerResource(inputs))
}

func newBrokerResourceClosure(templateEntity brokerEntity[schema.Schema]) func() resource.Resource {
	return func() resource.Resource {
		var r = brokerResource(templateEntity)
		return &r
	}
}

func forceBrokerRequirementsCheck() {
	apiAlreadyChecked = false
}

func checkBrokerRequirements(ctx context.Context, client *semp.Client) error {
	if !skipApiCheck && !apiAlreadyChecked {
		lock.Lock()
		defer lock.Unlock()
		if apiAlreadyChecked {
			return nil
		}
		path := "/about/api"
		result, err := client.RequestWithoutBody(ctx, http.MethodGet, path)
		if err != nil {
			return err
		}
		// To support broker developer versions ignore "+" in the returned version
		brokerSempVersion, err := version.NewVersion(strings.Replace(result["sempVersion"].(string), "+", "", -1))
		if err != nil {
			return err
		}
		minSempVersion, _ := version.NewVersion(minRequiredBrokerSempApiVersion)
		if brokerSempVersion.LessThan(minSempVersion) {
			return fmt.Errorf("broker SEMP API version %s does not meet provider required minimum SEMP API version: %s", brokerSempVersion, minSempVersion)
		}
		brokerPlatform := result["platform"].(string)
		if brokerPlatform != SempDetail.Platform {
			return fmt.Errorf("broker platform \"%s\" does not match provider supported platform: %s", BrokerPlatformName[brokerPlatform], BrokerPlatformName[SempDetail.Platform])
		}
		apiAlreadyChecked = true
	}
	return nil
}

// Compares the value with the attribute default value. Must take care of type conversions.
func isValueEqualsAttrDefault(attr *AttributeInfo, response tftypes.Value, brokerDefault tftypes.Value) (bool, error) {
	responseValue, err := attr.Converter.FromTerraform(response)
	if err != nil {
		return false, err
	}
	if attr.Default == nil {
		if brokerDefault.IsNull() {
			// No broker default
			return false, nil
		}
		// Analyze broker default
		brokerDefaultValue, err := attr.Converter.FromTerraform(brokerDefault)
		if err != nil {
			return false, err
		}
		// compare
		return responseValue == brokerDefaultValue, nil
	}
	tfDefault, err := attr.Converter.ToTerraform(attr.Default)
	if err != nil {
		return false, err
	}
	attrDefaultValue, err := attr.Converter.FromTerraform(tfDefault)
	if err != nil {
		return false, err
	}
	return responseValue == attrDefaultValue, nil
}

func toId(path string) string {
	return filepath.Base(path)
}

func (r *brokerResource) resetResponse(attributes []*AttributeInfo, response tftypes.Value, brokerDefaults tftypes.Value, state tftypes.Value, isObject bool) (tftypes.Value, error) {
	responseValues := map[string]tftypes.Value{}
	err := response.As(&responseValues)
	if err != nil {
		return tftypes.Value{}, err
	}
	stateValues := map[string]tftypes.Value{}
	err = state.As(&stateValues)
	if err != nil {
		return tftypes.Value{}, err
	}
	brokerDefaultValues := map[string]tftypes.Value{}
	err = brokerDefaults.As(&brokerDefaultValues)
	if err != nil {
		return tftypes.Value{}, err
	}
	for _, attr := range attributes {
		name := attr.TerraformName
		// Skip attributes that were filtered out from the converter/schema
		// (read-only non-identifying attributes for resources)
		if !attr.Identifying && attr.ReadOnly {
			continue
		}
		response, responseExists := responseValues[name]
		state, stateExists := stateValues[name]
		if responseExists && response.IsKnown() && !response.IsNull() {
			if len(attr.Attributes) != 0 {
				// This case is an object, typically threshold attributes
				v, err := r.resetResponse(attr.Attributes, response, tftypes.NewValue(attr.TerraformType, nil), state, true)
				if err != nil {
					return tftypes.Value{}, err
				}
				responseValuesMap := map[string]tftypes.Value{}
				err = v.As(&responseValuesMap)
				if err != nil {
					return tftypes.Value{}, err
				}
				allDefaults := true
				for _, attr := range responseValuesMap {
					if !attr.IsNull() {
						allDefaults = false
						break
					}
				}
				if allDefaults {
					// Set the whole object to null
					responseValues[name] = tftypes.NewValue(attr.TerraformType, nil)
				} else {
					// Keep the object with individual attributes
					responseValues[name] = v
				}
			} else {
				isResponseValueDefault, err := isValueEqualsAttrDefault(attr, response, brokerDefaultValues[name])
				if err != nil {
					return tftypes.Value{}, err
				}
				if !isResponseValueDefault {
					continue // do not change response for this attr if it was non-default
				}
				if !stateExists && isObject {
					responseValues[name] = tftypes.NewValue(attr.TerraformType, nil)
				} else if stateExists && state.IsNull() {
					responseValues[name] = state
				} // else leave attr response unchanged
			}
		} else if stateExists && attr.Sensitive {
			responseValues[name] = state
		} else {
			responseValues[name] = tftypes.NewValue(attr.TerraformType, nil)
		}
	}
	return tftypes.NewValue(response.Type(), responseValues), nil
}

func (r *brokerResource) findBrokerDefaults(attributes []*AttributeInfo, response tftypes.Value, request tftypes.Value) (any, error) {
	defaultValues := map[string]tftypes.Value{}
	requestValues := map[string]tftypes.Value{}
	err := request.As(&requestValues)
	if err != nil {
		return nil, err
	}
	responseValues := map[string]tftypes.Value{}
	err = response.As(&responseValues)
	if err != nil {
		return nil, err
	}
	for _, attr := range attributes {
		// Set obtained default values and null for any other attributes
		name := attr.TerraformName
		if !attr.Identifying && attr.ReadOnly {
			continue
		}
		if attr.Default == nil && requestValues[name].IsNull() && attr.BaseType != Struct {
			defaultValues[name] = responseValues[name]
		} else {
			defaultValues[name] = tftypes.NewValue(attr.TerraformType, nil)
		}
	}
	return r.converter.FromTerraform(tftypes.NewValue(request.Type(), defaultValues))
}

func (r *brokerResource) Schema(_ context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	// Overwrite the schema version with the provider major version
	providerMajorVersion := getProviderMajorVersion(ProviderVersion)
	r.schema.Version = providerMajorVersion
	response.Schema = r.schema
}

func (r *brokerResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_" + r.terraformName
}

func (r *brokerResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	client, ok := request.ProviderData.(*semp.Client)
	if !ok {
		response.Diagnostics.AddError(
			"Unexpected resource configuration",
			fmt.Sprintf("Unexpected type %T for provider data; expected %T.", request.ProviderData, client),
		)
		return
	}
	r.client = client
}

func (r *brokerResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	client := r.client
	if err := checkBrokerRequirements(ctx, client); err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Broker check failed", err)
		return
	}

	sempData, err := r.converter.FromTerraform(request.Plan.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Error converting data", err)
		return
	}

	var sempPath string
	method := http.MethodPut
	if r.postPathTemplate != "" {
		method = http.MethodPost
		sempPath, err = resolveSempPath(r.postPathTemplate, r.identifyingAttributes, request.Plan.Raw)
	} else {
		sempPath, err = resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.Plan.Raw)
	}
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Error generating SEMP path", err)
		return
	}
	if r.objectType == SingletonObject {
		// if the object is a singleton, PATCH rather than PUT
		method = http.MethodPatch
	}
	jsonResponseData, err := client.RequestWithBody(ctx, method, sempPath, sempData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "SEMP call failed", err)
		return
	}
	// Determine broker defaults as each attribute response, where request was set to null and it didn't have a default
	//   then store it as private data
	tfResponseData, err := r.converter.ToTerraform(jsonResponseData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "SEMP response conversion failed", err)
		return
	}
	brokerDefaultsData, err := r.findBrokerDefaults(r.attributes, tfResponseData, request.Plan.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Response postprocessing failed", err)
		return
	}
	privatData, err := json.Marshal(brokerDefaultsData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Response postprocessing failed", err)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("Create: determined following broker-defined defaults:\n%v", brokerDefaultsData))
	response.Private.SetKey(ctx, defaults, privatData)
	// Set the response
	response.State.Raw = request.Plan.Raw
}

func (r *brokerResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	client := r.client
	if err := checkBrokerRequirements(ctx, client); err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Broker check failed", err)
		return
	}
	sempPath, err := resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.State.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Error generating SEMP path", err)
		return
	}
	sempData, err := client.RequestWithoutBody(ctx, http.MethodGet, sempPath)
	if err != nil {
		if errors.Is(err, semp.ErrResourceNotFound) {
			tflog.Info(ctx, fmt.Sprintf("Detected missing resource %v, removing from state", sempPath))
			response.State.RemoveResource(ctx)
		} else {
			addErrorToDiagnostics(&response.Diagnostics, "SEMP call failed", err)
		}
		return
	}
	responseData, err := r.converter.ToTerraform(sempData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "SEMP response conversion failed", err)
		return
	}
	defaultsJson, diags := request.Private.GetKey(ctx, defaults)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}
	if defaultsJson == nil {
		defaultsJson = []byte("{}")
	}
	brokerDefaultsData := map[string]any{}
	err = json.Unmarshal(defaultsJson, &brokerDefaultsData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Retrieve of defaults failed", err)
		return
	}
	defaultsData, err := r.converter.ToTerraform(brokerDefaultsData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Retrieve of defaults failed", err)
		return
	}
	// Replace default values in response to null
	responseData, err = r.resetResponse(r.attributes, responseData, defaultsData, request.State.Raw, false)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Read response postprocessing failed", err)
		return
	}
	response.State.Raw = responseData
}

func (r *brokerResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	client := r.client
	if err := checkBrokerRequirements(ctx, client); err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Broker check failed", err)
		return
	}
	sempData, err := r.converter.FromTerraform(request.Plan.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Error converting data", err)
		return
	}
	sempPath, err := resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.Plan.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Error generating SEMP path", err)
		return
	}
	method := http.MethodPut
	if r.objectType == SingletonObject {
		method = http.MethodPatch
	}
	jsonResponseData, err := client.RequestWithBody(ctx, method, sempPath, sempData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "SEMP call failed", err)
		return
	}
	// Determine broker defaults as each attribute response, where request was set to null and it didn't have a default
	//   then store it as private data
	tfResponseData, err := r.converter.ToTerraform(jsonResponseData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "SEMP response conversion failed", err)
		return
	}
	brokerDefaultsData, err := r.findBrokerDefaults(r.attributes, tfResponseData, request.Plan.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Response postprocessing failed", err)
		return
	}
	privatData, err := json.Marshal(brokerDefaultsData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Response postprocessing failed", err)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("Update: determined following broker-defined defaults:\n%v", brokerDefaultsData))
	response.Private.SetKey(ctx, defaults, privatData)
	// Set the response
	response.State.Raw = request.Plan.Raw
}

func (r *brokerResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	client := r.client
	if err := checkBrokerRequirements(ctx, client); err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Broker check failed", err)
		return
	}
	// don't actually do anything if the object is a singleton
	if r.objectType == SingletonObject {
		addWarningToDiagnostics(&response.Diagnostics, fmt.Sprintf("Associated state will be removed but singleton object %s cannot be deleted", r.terraformName), ErrDeleteSingletonOrDefaultsNotAllowed)
		return
	}
	path, err := resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.State.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Error generating SEMP path", err)
		return
	}
	// don't actually do anything if the object is a default object
	if toId(path) == defaultObjectName {
		switch r.terraformName {
		case
			"msg_vpn",
			"msg_vpn_client_profile",
			"msg_vpn_acl_profile",
			"msg_vpn_client_username":
			addWarningToDiagnostics(&response.Diagnostics, fmt.Sprintf("Associated state will be removed but default object %s, \"%s\" cannot be deleted", r.terraformName, toId(path)), ErrDeleteSingletonOrDefaultsNotAllowed)
			return
		}
	}
	// request delete
	_, err = client.RequestWithoutBody(ctx, http.MethodDelete, path)
	if err != nil {
		if !errors.Is(err, semp.ErrResourceNotFound) {
			addErrorToDiagnostics(&response.Diagnostics, "SEMP call failed", err)
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Detected object %s, \"%s\" was already missing from the broker, removing from state", r.terraformName, toId(path)))
		// Let destroy finish normally if the error was Resource Not Found - only means that the resource has already been removed from the broker.
	}
}

func (r *brokerResource) ImportState(_ context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {

	if len(r.identifyingAttributes) == 0 {
		if request.ID != "" {
			response.Diagnostics.AddError(
				"singleton object requires empty identifier for import",
				"singleton object requires empty identifier for import",
			)
		}
		response.State.Raw = tftypes.NewValue(tftypes.Object{}, nil)
		return
	}
	split := strings.Split(strings.ReplaceAll(request.ID, ",", "/"), "/")
	if len(split) != len(r.identifyingAttributes) {
		r.addIdentifierErrorToDiagnostics(&response.Diagnostics, request.ID)
		return
	}

	identifierData := map[string]any{}
	for i, attr := range r.identifyingAttributes {
		v, err := url.PathUnescape(split[i])
		if err != nil {
			r.addIdentifierErrorToDiagnostics(&response.Diagnostics, request.ID)
		}
		identifierData[attr.SempName] = v
	}
	identifierState, err := r.converter.ToTerraform(identifierData)
	if err != nil {
		r.addIdentifierErrorToDiagnostics(&response.Diagnostics, request.ID)
		return
	}
	response.State.Raw = identifierState
}

func addErrorToDiagnostics(diags *diag.Diagnostics, summary string, err error) {
	for err != nil {
		diags.AddError(summary, err.Error())
		err = errors.Unwrap(err)
	}
}

func addWarningToDiagnostics(diags *diag.Diagnostics, summary string, err error) {
	for err != nil {
		diags.AddWarning(summary, err.Error())
		err = errors.Unwrap(err)
	}
}

func (r *brokerResource) addIdentifierErrorToDiagnostics(diags *diag.Diagnostics, id string) {
	var identifiers []string
	for _, attr := range r.identifyingAttributes {
		identifiers = append(identifiers, attr.TerraformName)
	}
	addErrorToDiagnostics(
		diags,
		"invalid identifier",
		fmt.Errorf("invalid identifier %v, identifier must be of the form %v with each segment URL-encoded as necessary", id, strings.Join(identifiers, "/")))
}

func (r *brokerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return nil
}

func (r *brokerResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	schema := r.schema
	converter := r.converter
	version := getProviderMajorVersion(ProviderVersion)
	upgraders := make(map[int64]resource.StateUpgrader)
	// new code will add upgraders for each version, starting from 0
	// note that upgraders are the same for each version
	for i := int64(0); i < version; i++ {
		upgraders[i] = resource.StateUpgrader{
			PriorSchema: &schema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// On one side, the raw state
				var oldStateAttributes map[string]interface{}
				if err := json.Unmarshal(req.RawState.JSON, &oldStateAttributes); err != nil {
					resp.Diagnostics.AddError("State conversion failed", err.Error())
				}
				// On the other side, read same through interpreting it in the new schema - this will keep attributes that are included in the new schema
				rawState := req.State.Raw
				resourceData, err := converter.FromTerraform(rawState)
				if err != nil {
					resp.Diagnostics.AddError("State conversion failed", err.Error())
				}
				conversionResults, err := converter.ToTerraform(resourceData)
				if err != nil {
					resp.Diagnostics.AddError("State conversion failed", err.Error())
				}
				var resultsDataMap map[string]tftypes.Value
				err = conversionResults.As(&resultsDataMap)
				if err != nil {
					resp.Diagnostics.AddError("State conversion failed", err.Error())
				}
				// iterate old state attributes and if a value is not null and the attribute is to be removed then fail the upgrade
				for key, value := range oldStateAttributes {
					if value != nil {
						// try to find the key in resourceDataMap
						data, ok := resultsDataMap[key]
						if !ok {
							resp.Diagnostics.AddError("State upgrade failed", fmt.Sprintf("Found deprecated state key '%s', unable to upgrade state if value is not null", key))
						} else {
							// if the type of value is map[string]interface{} then it is a nested object
							if _, ok := value.(map[string]interface{}); ok {
								var resultsDataMap2 map[string]tftypes.Value
								err = data.As(&resultsDataMap2)
								if err != nil {
									resp.Diagnostics.AddError("State conversion failed", err.Error())
								}
								for nestedKey, val := range value.(map[string]interface{}) {
									if val != nil {
										_, ok := resultsDataMap2[nestedKey]
										if !ok {
											resp.Diagnostics.AddError("State upgrade failed", fmt.Sprintf("Found deprecated state key '%s' in nested object '%s', unable to upgrade state if value is not null", nestedKey, key))
										}
									}
								}
							}
						}
					}
				}
				resp.State.Raw = conversionResults
			},
		}
	}
	return upgraders
}
