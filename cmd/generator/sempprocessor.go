// terraform-provider-solacebroker
//
// Copyright 2024 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"terraform-provider-solacebroker/internal/broker"
)

func processSempResults(resourceTypeAndName string, attributes []*broker.AttributeInfo, values []map[string]interface{}, parentInfo BrokerObjectInstanceInfo) ([]ResourceConfig, map[string]VariableConfig, error) {
	var tfBrokerObjects []ResourceConfig
	var tfVariables = map[string]VariableConfig{}
	var attributeLookup = map[string]int{}
	for k := range values {
		resourceConfig := ResourceConfig{
			ResourceAttributes: map[string]ResourceAttributeInfo{},
		}
		systemProvisioned := false
		attributesWithDefaultValue := map[string]*string{} // list of attributes with default values
		linkedAttributes := map[string][]string{}

		for i, attr := range attributes {
			attributeLookup[attr.TerraformName] = i
			if len(attr.Requires) > 0 {
				linkedAttributes[attr.TerraformName] = append(linkedAttributes[attr.TerraformName], attr.Requires...)
			}
			if attr.Sensitive {
				// write-only attributes can't be retrieved, so we don't expose them
				attributesWithDefaultValue[attr.TerraformName] = nil
				continue
			}
			if !attr.Identifying && attr.ReadOnly {
				// read-only attributes should only be in the datasource
				continue
			}
			valuesRes := values[k][attr.SempName]
			if attr.Identifying {
				// iterate parentInfo.identifyingAttributes
				if parentInfo.identifyingAttributes != nil {
					foundReference := false
					for _, identifyingAttribute := range parentInfo.identifyingAttributes {
						if identifyingAttribute.key == attr.SempName {
							reference := strings.ReplaceAll(parentInfo.resourceTypeAndName, " ", ".") + "." + attr.TerraformName
							resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(reference)
							// this means that the attribute value is a reference to a parent attribute, no need to process further
							foundReference = true
							break
						}
					}
					if foundReference {
						continue
					}
				}
			}
			// TODO: revisit, likely not needed
			// } else if attr.TerraformName == "client_profile_name" && attributeExistInParent {
			// 	//peculiar use case where client_profile is not identifying for msg_vpn_client_username but it is dependent
			// 	resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(attributeParentNameAndValue)
			// 	continue
			// }

			switch attr.BaseType {
			case broker.String:
				if reflect.TypeOf(valuesRes) == nil {
					continue
				}
				if !attr.Identifying && valuesRes == "" {
					// non-identifying attributes with empty values will be skipped
					continue
				}
				// TODO: review if this covers all cases of system provisioned attributes
				if len(valuesRes.(string)) > 0 {
					systemProvisioned = isSystemProvisionedAttribute(valuesRes.(string))
				}
				val := "\"" + SanitizeHclStringValue(valuesRes.(string)) + "\""
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(valuesRes) {
					//attributes with default values will be skipped
					// WORKAROUND: Except if attribute is "authentication_basic_type" in "msg_vpn"
					// TODO: review if this is still needed
					// if brokerObjectTerraformName != "msg_vpn" || attr.TerraformName != "authentication_basic_type" {
					// 	attributesWithDefaultValue[attr.TerraformName] = &val
					// 	continue
					// } else {
					// 	LogCLIInfo("Applying workaround: not ignoring default for `msg_vpn` attribute `authentication_basic_type`")
					// }
					attributesWithDefaultValue[attr.TerraformName] = &val
					continue
				}
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			case broker.Int64:
				if valuesRes == nil {
					continue
				}
				intValue := valuesRes
				val := fmt.Sprintf("%v", intValue)
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(intValue) {
					//attributes with default values will be skipped
					attributesWithDefaultValue[attr.TerraformName] = &val
					continue
				}
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			case broker.Bool:
				if valuesRes == nil {
					continue
				}
				boolValue := valuesRes.(bool)
				val := strconv.FormatBool(boolValue)
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(boolValue) {
					//attributes with default values will be skipped
					attributesWithDefaultValue[attr.TerraformName] = &val
					continue
				}
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			case broker.Struct:
				valueJson, err := json.Marshal(valuesRes)
				if err != nil {
					continue
				}
				output := strings.ReplaceAll(string(valueJson), "clearPercent", "clear_percent")
				output = strings.ReplaceAll(output, "setPercent", "set_percent")
				output = strings.ReplaceAll(output, "clearValue", "clear_value")
				output = strings.ReplaceAll(output, "setValue", "set_value")
				val := output
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(valuesRes) {
					//attributes with default values will be skipped
					attributesWithDefaultValue[attr.TerraformName] = &val
					continue
				}
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			}
			// TODO: review here for proper handling of ancillary objects
			if attr.Deprecated && systemProvisioned {
				resourceConfig.ResourceAttributes[attr.TerraformName] = addCommentToAttributeInfo(resourceConfig.ResourceAttributes[attr.TerraformName],
					" # Note: This attribute is deprecated and may also be system provisioned.")
			} else if attr.Deprecated && !systemProvisioned {
				resourceConfig.ResourceAttributes[attr.TerraformName] = addCommentToAttributeInfo(resourceConfig.ResourceAttributes[attr.TerraformName],
					" # Note: This attribute is deprecated.")
			} else if !attr.Deprecated && systemProvisioned {
				resourceConfig.ResourceAttributes[attr.TerraformName] = addCommentToAttributeInfo(resourceConfig.ResourceAttributes[attr.TerraformName],
					" # Note: This attribute may be system provisioned and a \"depends_on\" meta-argument may be required to the parent object of this attribute to ensure proper order of creation")
			}
			systemProvisioned = false // reset systemProvisioned
		}
		// Iterate linkedAttributes
		// for each attribute check if there is an entry in resourceConfig
		// if so, iterate all strings as attributes and add them to resourceConfig if not already there, with value from attributesWithDefaultValue
		for attrName := range linkedAttributes {
			if _, ok := resourceConfig.ResourceAttributes[attrName]; ok {
				sensitive := false
				for _, linkedAttrName := range linkedAttributes[attrName] {
					if _, ok := resourceConfig.ResourceAttributes[linkedAttrName]; !ok {
						if val, ok := attributesWithDefaultValue[linkedAttrName]; ok {
							if val == nil {
								sensitive = true
								break
							}
							resourceConfig.ResourceAttributes[linkedAttrName] = newAttributeInfo(*val)
						}
					}
				}
				if sensitive {
					resourceName := strings.Split(resourceTypeAndName, " ")[1]
					// if any linked attribute is sensitive, use variables
					// first add the current attribute to variables
					variableName := resourceName + "_" + attrName
					attrInfo := attributes[attributeLookup[attrName]]
					attrSensitive := attrInfo.Sensitive
					attrType, attrDefault, _ := GetBaseTypeAndDefault(attrInfo)
					val, ok := resourceConfig.ResourceAttributes[attrName]
					if ok && val.AttributeValue != "" {
						attrDefault = val.AttributeValue
					}
					newVariable := VariableConfig{
						Type:      attrType,
						Default:   attrDefault,
						Sensitive: attrSensitive, // note: expecting not sensitive since this was a defined attribute
					}
					tfVariables[variableName] = newVariable
					resourceConfig.ResourceAttributes[attrName] = newAttributeInfo("var." + variableName)
					for _, linkedAttrName := range linkedAttributes[attrName] {
						// then add the linked attributes to variables
						attrInfo := attributes[attributeLookup[linkedAttrName]]
						attrSensitive := attrInfo.Sensitive
						attrType, attrDefault, _ := GetBaseTypeAndDefault(attrInfo)
						val, ok := resourceConfig.ResourceAttributes[linkedAttrName]
						if ok && val.AttributeValue != "" {
							attrDefault = val.AttributeValue
						}
						variableName := resourceName + "_" + linkedAttrName
						newVariable := VariableConfig{
							Type:      attrType,
							Default:   attrDefault,
							Sensitive: attrSensitive,
						}
						tfVariables[variableName] = newVariable
						resourceConfig.ResourceAttributes[linkedAttrName] = newAttributeInfo("var." + variableName)
					}
				}
			}
		}
		tfBrokerObjects = append(tfBrokerObjects, resourceConfig)
	}
	return tfBrokerObjects, tfVariables, nil
}

func newAttributeInfo(value string) ResourceAttributeInfo {
	return ResourceAttributeInfo{
		AttributeValue: value,
		Comment:        "",
	}
}

func addCommentToAttributeInfo(info ResourceAttributeInfo, comment string) ResourceAttributeInfo {
	return ResourceAttributeInfo{
		AttributeValue: info.AttributeValue,
		Comment:        comment,
	}
}

func GetBaseTypeAndDefault(attrInfo *broker.AttributeInfo) (string, string, error) {
	defaultValue := ""
	switch attrInfo.BaseType {
	case broker.String:
		defaultValue = "\"" + fmt.Sprint(attrInfo.Default) + "\""
		return "string", defaultValue, nil
	case broker.Int64:
		defaultValue = fmt.Sprint(attrInfo.Default)
		return "number", defaultValue, nil
	case broker.Bool:
		defaultValue = strconv.FormatBool(attrInfo.Default.(bool))
		return "bool", defaultValue, nil
	case broker.Struct:
		// TODO: implement this
		return "object", "", nil
	default:
		return "", "", errors.New("unknown base type")
	}
}
