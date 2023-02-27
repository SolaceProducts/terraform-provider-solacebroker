// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
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
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"net/http"
	"net/url"
	"strings"
)

const (
	applied = "applied"
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

var (
	_ resource.ResourceWithConfigure        = &brokerResource{}
	_ resource.ResourceWithConfigValidators = &brokerResource{}
	_ resource.ResourceWithImportState      = &brokerResource{}
)

type brokerResource brokerEntity[schema.Schema]

func (r *brokerResource) resetResponse(attributes []*AttributeInfo, response tftypes.Value, state tftypes.Value) (tftypes.Value, error) {
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
	for _, attr := range attributes {
		name := attr.TerraformName
		response, responseExists := responseValues[name]
		state, stateExists := stateValues[name]
		if responseExists && response.IsKnown() && !response.IsNull() {
			if stateExists && state.IsNull() {
				responseValues[name] = state
			} else {
				if len(attr.Attributes) != 0 {
					v, err := r.resetResponse(attr.Attributes, response, state)
					if err != nil {
						return tftypes.Value{}, err
					}
					responseValues[name] = v
				}
			}
		} else if stateExists && attr.Sensitive {
			responseValues[name] = state
		} else {
			responseValues[name] = tftypes.NewValue(attr.TerraformType, nil)
		}
	}
	return tftypes.NewValue(response.Type(), responseValues), nil
}

func (r *brokerResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = r.schema
}

func (r *brokerResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_" + r.terraformName
}

func (r *brokerResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	config, ok := request.ProviderData.(*providerData)
	if !ok {
		response.Diagnostics = diag.Diagnostics{diag.NewErrorDiagnostic("Unexpected resource configuration", fmt.Sprintf("Unexpected type %T for provider data; expected %T.", request.ProviderData, config))}
		return
	}
	r.providerData = config
}

func generateDiagnostics(summary string, err error) diag.Diagnostics {
	diags := &diag.Diagnostics{}
	for err != nil {
		diags.AddError(summary, err.Error())
		err = errors.Unwrap(err)
	}
	return *diags
}

func (r *brokerResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	client, d := client(r.providerData)
	if d != nil {
		response.Diagnostics.Append(d)
		if response.Diagnostics.HasError() {
			return
		}
	}

	sempData, err := r.converter.FromTerraform(request.Plan.Raw)
	if err != nil {
		response.Diagnostics = generateDiagnostics("Error converting data", err)
		return
	}

	var path string
	method := http.MethodPut
	if r.postPathTemplate != "" {
		method = http.MethodPost
		path, err = resolveSempPath(r.postPathTemplate, r.identifyingAttributes, request.Plan.Raw)
	} else {
		path, err = resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.Plan.Raw)
	}
	if err != nil {
		response.Diagnostics = generateDiagnostics("Error generating SEMP path", err)
		return
	}
	if r.objectType == SingletonObject {
		// if the object is a singleton, PATCH rather than PUT
		method = http.MethodPatch
	}
	_, err = client.RequestWithBody(method, path, sempData)
	if err != nil {
		response.Diagnostics = generateDiagnostics("SEMP call failed", err)
		return
	}

	response.State.Raw = request.Plan.Raw
	response.Private.SetKey(ctx, applied, []byte("true"))
}

func (r *brokerResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	client, d := client(r.providerData)
	if d != nil {
		response.Diagnostics.Append(d)
		if response.Diagnostics.HasError() {
			return
		}
	}

	path, err := resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.State.Raw)
	if err != nil {
		response.Diagnostics = generateDiagnostics("Error generating SEMP path", err)
		return
	}
	sempData, err := client.RequestWithoutBody(http.MethodGet, path)
	if err != nil {
		response.Diagnostics = generateDiagnostics("SEMP call failed", err)
		return
	}

	responseData, err := r.converter.ToTerraform(sempData)
	if err != nil {
		response.Diagnostics = generateDiagnostics("SEMP response conversion failed", err)
		return
	}

	applied, diags := request.Private.GetKey(ctx, applied)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}
	if string(applied) == "true" {
		responseData, err = r.resetResponse(r.attributes, responseData, request.State.Raw)
		if err != nil {
			response.Diagnostics = generateDiagnostics("Response postprocessing failed", err)
			return
		}
	}

	response.State.Raw = responseData
}

func (r *brokerResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	client, d := client(r.providerData)
	if d != nil {
		response.Diagnostics.Append(d)
		if response.Diagnostics.HasError() {
			return
		}
	}

	sempData, err := r.converter.FromTerraform(request.Plan.Raw)
	if err != nil {
		response.Diagnostics = generateDiagnostics("Error converting data", err)
		return
	}

	path, err := resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.Plan.Raw)
	if err != nil {
		response.Diagnostics = generateDiagnostics("Error generating SEMP path", err)
		return
	}
	method := http.MethodPut
	if r.objectType == SingletonObject {
		method = http.MethodPatch
	}
	_, err = client.RequestWithBody(method, path, sempData)
	if err != nil {
		response.Diagnostics = generateDiagnostics("SEMP call failed", err)
		return
	}

	response.State.Raw = request.Plan.Raw
	response.Private.SetKey(ctx, applied, []byte("true"))
}

func (r *brokerResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	if r.objectType == SingletonObject {
		// don't actually do anything if the object is a singleton
		return
	}

	client, d := client(r.providerData)
	if d != nil {
		response.Diagnostics.Append(d)
		if response.Diagnostics.HasError() {
			return
		}
	}

	path, err := resolveSempPath(r.pathTemplate, r.identifyingAttributes, request.State.Raw)
	if err != nil {
		response.Diagnostics = generateDiagnostics("Error generating SEMP path", err)
		return
	}
	_, err = client.RequestWithoutBody(http.MethodDelete, path)
	if err != nil {
		response.Diagnostics = generateDiagnostics("SEMP call failed", err)
		return
	}
}

func (r *brokerResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {

	if len(r.identifyingAttributes) == 0 {
		if request.ID != "" {
			response.Diagnostics.AddError(
				"singleton object requires empty identifier for import",
				"singleton object requires empty identifier for import",
			)
		}
		return
	}
	split := strings.Split(strings.ReplaceAll(request.ID, ",", "/"), "/")
	if len(split) != len(r.identifyingAttributes) {
		response.Diagnostics = r.generateIdentifierDiagnostic(request.ID)
		return
	}

	identifierData := map[string]any{}
	for i, attr := range r.identifyingAttributes {
		v, err := url.PathUnescape(split[i])
		if err != nil {
			response.Diagnostics = r.generateIdentifierDiagnostic(request.ID)
		}
		identifierData[attr.SempName] = v
	}
	identifierState, err := r.converter.ToTerraform(identifierData)
	if err != nil {
		response.Diagnostics = r.generateIdentifierDiagnostic(request.ID)
		return
	}
	response.State.Raw = identifierState
}

func (r *brokerResource) generateIdentifierDiagnostic(id string) diag.Diagnostics {
	var identifiers []string
	for _, id := range r.identifyingAttributes {
		identifiers = append(identifiers, id.TerraformName)
	}
	return generateDiagnostics(
		"invalid identifier",
		fmt.Errorf("invalid identifier %v, identifier must be of the form %v with each segment URL-encoded as necessary", id, strings.Join(identifiers, "/")))
}

func (r *brokerResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return nil
}
