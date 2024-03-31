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
package terraform

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
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

	// Loop through entities and build database
	resourcesPathSignatureMap := map[string]string{}
	e := internalbroker.Entities
	for _, ds := range e {
		// Create new entry for each resource
		BrokerObjectRelationship[BrokerObjectType(ds.TerraformName)] = []BrokerObjectType{}
		// Build a signature for each resource
		rex := regexp.MustCompile(`{[^\/]*}`)
		signature := strings.TrimSuffix(strings.Replace(rex.ReplaceAllString(ds.PathTemplate, ""), "//", "/", -1),"/") // Find all parameters in path template enclosed in {} including multiple ones
		if signature != "" {
			resourcesPathSignatureMap[signature] = ds.TerraformName
		}
	}
	// Loop through entities again and add children to parents
	for _, ds := range e {
		// Parent signature for each resource and add
		rex := regexp.MustCompile(`{[^\/]*}`)
		signature := strings.TrimSuffix(strings.Replace(rex.ReplaceAllString(ds.PathTemplate, ""), "//", "/", -1),"/")
		// get parentSignature by removing the part of signature after the last /
		parentSignature := path.Dir(signature)
		if parentSignature != "." && parentSignature != "/" {
			parentResource := resourcesPathSignatureMap[parentSignature]
			BrokerObjectRelationship[BrokerObjectType(parentResource)] = append(BrokerObjectRelationship[BrokerObjectType(parentResource)], BrokerObjectType(ds.TerraformName))	
		}
	}
}

func ParseTerraformObject(ctx context.Context, client semp.Client, resourceName string, brokerObjectTerraformName string, providerSpecificIdentifier string, parentBrokerResourceAttributesRelationship map[string]string, parentResult map[string]any) GeneratorTerraformOutput {
	var objectName string
	tfObject := map[string]ResourceConfig{}
	tfObjectSempDataResponse := map[string]map[string]any{}
	entityToRead := internalbroker.EntityInputs{}
	// TODO: potentially expensive
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
				ExitWithError("SEMP call failed. " + err.Error() + " on path " + path)
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

	// TODO: this should be optimized

	resourceName := GenerateRandomString(6) //use generated if not able to identify

	resourceTerraformName = strings.Split(resourceTerraformName, " ")[0]
	resourceTerraformName = strings.ReplaceAll(strings.ToLower(resourceTerraformName), "solacebroker_", "")

	//Get identifying attribute name to differentiate from multiples
	// TODO: potentially expensive
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
