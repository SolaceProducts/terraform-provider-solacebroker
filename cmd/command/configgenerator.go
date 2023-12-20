// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
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
package terraform

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"net/http"
	"os"
	"regexp"
	"strings"
	internalbroker "terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/broker/generated"
	"terraform-provider-solacebroker/internal/semp"
)

type BrokerObjectType string

type GeneratorTerraformOutput struct {
	TerraformOutput  map[string]ResourceConfig
	SEMPDataResponse map[string]map[string]any
}

var BrokerObjectRelationship = map[BrokerObjectType][]BrokerObjectType{}

type BrokerRelationParameterPath struct {
	path          string
	terraformName string
}

var ObjectNamesCount = map[string]int{}

func CreateBrokerObjectRelationships() {

	//loop through entities
	terraformNamePathMap := map[string]BrokerRelationParameterPath{}
	for _, ds := range internalbroker.Entities {
		rex := regexp.MustCompile(`{[^{}]*}`)
		matches := rex.FindAllStringSubmatch(ds.PathTemplate, -1)

		BrokerObjectRelationship[BrokerObjectType(ds.TerraformName)] = []BrokerObjectType{}

		for i := range matches {

			if i == 0 || len(matches) <= 1 {
				var firstParameter string
				firstParameter = strings.TrimPrefix(matches[0][0], "{")
				firstParameter = strings.TrimSuffix(firstParameter, "}")

				_, ok := terraformNamePathMap[firstParameter]
				if !ok {
					terraformNamePathMap[firstParameter] = BrokerRelationParameterPath{
						ds.PathTemplate,
						ds.TerraformName,
					}
				}
			} else {
				firstParameter := strings.TrimPrefix(matches[0][0], "{")
				firstParameter = strings.TrimSuffix(firstParameter, "}")

				//check if parent path is part of child path before we add
				firstParameterRelationship, firstParameterRelationshipExist := terraformNamePathMap[firstParameter]
				if firstParameterRelationshipExist && strings.Contains(ds.PathTemplate, firstParameterRelationship.path) {

					children, ok := BrokerObjectRelationship[BrokerObjectType(firstParameterRelationship.terraformName)]
					if !ok {
						BrokerObjectRelationship[BrokerObjectType(firstParameterRelationship.terraformName)] = []BrokerObjectType{}
						children = []BrokerObjectType{}
					} else {
						if !slices.Contains(children, BrokerObjectType(ds.TerraformName)) && ds.TerraformName != firstParameterRelationship.terraformName {
							children = append(children, BrokerObjectType(ds.TerraformName))
						}
						//confirm if this should be child of child
						for k := range children {
							childrenOfChild, childrenOfChildExists := BrokerObjectRelationship[children[k]]
							if childrenOfChildExists {
								if !slices.Contains(childrenOfChild, BrokerObjectType(ds.TerraformName)) &&
									strings.Contains(ds.TerraformName, string(children[k])) && ds.TerraformName != string(children[k]) {
									childrenOfChild = append(childrenOfChild, BrokerObjectType(ds.TerraformName))
								}
							}
							BrokerObjectRelationship[children[k]] = childrenOfChild
						}
					}

					BrokerObjectRelationship[BrokerObjectType(firstParameterRelationship.terraformName)] = children
				} else {

					terraformNamePathMap[firstParameter] = BrokerRelationParameterPath{
						ds.PathTemplate,
						ds.TerraformName,
					}

				}
			}
		}
	}
}

func ParseTerraformObject(ctx context.Context, client semp.Client, resourceName string, brokerObjectTerraformName string, providerSpecificIdentifier string, parentBrokerResourceAttributesRelationship map[string]string, parentResult map[string]any) GeneratorTerraformOutput {
	var objectName string
	tfObject := map[string]ResourceConfig{}
	tfObjectSempDataResponse := map[string]map[string]any{}
	entityToRead := internalbroker.EntityInputs{}
	for _, ds := range internalbroker.Entities {
		if strings.ToLower(ds.TerraformName) == strings.ToLower(brokerObjectTerraformName) {
			entityToRead = ds
			break
		}
	}
	var path string

	if len(parentResult) > 0 {
		path, _ = ResolveSempPathWithParent(entityToRead.PathTemplate, parentResult)
	} else {
		path, _ = ResolveSempPath(entityToRead.PathTemplate, providerSpecificIdentifier)
	}

	if len(path) > 0 {

		sempData, err := client.RequestWithoutBodyForGenerator(ctx, generated.BasePath, http.MethodGet, path, []map[string]any{})
		if err != nil {
			if err == semp.ErrResourceNotFound {
				// continue if error is resource not found
				if len(parentResult) > 0 {
					print("..")
				}
				sempData = []map[string]any{}
			} else if errors.Is(err, semp.ErrBadRequest) {
				// continue if error is also bad request
				if len(parentResult) > 0 {
					print("..")
				}
				sempData = []map[string]any{}
			} else {
				LogCLIError("SEMP call failed. " + err.Error() + " on path " + path)
				os.Exit(1)
			}
		}

		resourceKey := "solacebroker_" + brokerObjectTerraformName + " " + resourceName

		resourceValues, err := GenerateTerraformString(entityToRead.Attributes, sempData, parentBrokerResourceAttributesRelationship, brokerObjectTerraformName)

		//check resource names used and deduplicate to avoid collision
		for i := range resourceValues {
			totalOccurrence := 1
			objectName = strings.ToLower(resourceKey) + GetNameForResource(strings.ToLower(resourceKey), resourceValues[i])
			count, objectNameExists := ObjectNamesCount[objectName]
			if objectNameExists {
				totalOccurrence = count + 1
			}
			ObjectNamesCount[objectName] = totalOccurrence
			objectName = objectName + "_" + fmt.Sprint(totalOccurrence)
			tfObject[objectName] = resourceValues[i]
			tfObjectSempDataResponse[objectName] = sempData[i]
		}
	}
	return GeneratorTerraformOutput{
		TerraformOutput:  tfObject,
		SEMPDataResponse: tfObjectSempDataResponse,
	}
}

func GetNameForResource(resourceTerraformName string, attributeResourceTerraform ResourceConfig) string {

	resourceName := GenerateRandomString(6) //use generated if not able to identify

	resourceTerraformName = strings.Split(resourceTerraformName, " ")[0]
	resourceTerraformName = strings.ReplaceAll(strings.ToLower(resourceTerraformName), "solacebroker_", "")

	//Get identifying attribute name to differentiate from multiples
	for _, ds := range internalbroker.Entities {
		if ds.TerraformName == resourceTerraformName {
			for _, attr := range ds.Attributes {
				if attr.Identifying &&
					(strings.Contains(strings.ToLower(attr.TerraformName), "name") ||
						strings.Contains(strings.ToLower(attr.TerraformName), "topic")) {
					// intentionally continue looping till we get the best name
					attr, found := attributeResourceTerraform.ResourceAttributes[attr.TerraformName]
					value := attr.AttributeValue
					if strings.Contains(value, ".") {
						continue
					}
					if found {
						//sanitize name
						resourceName = "_" + value
					}
				}
			}
		}
	}
	return SanitizeHclIdentifierName(resourceName)
}
