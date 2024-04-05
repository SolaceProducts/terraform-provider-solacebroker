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
	"context"
	// "errors"
	// "fmt"
	// "maps"
	// "net/http"
	"os"
	"path"
	"regexp"
	"strings"
	internalbroker "terraform-provider-solacebroker/internal/broker"
	// "terraform-provider-solacebroker/internal/broker/generated"
	"terraform-provider-solacebroker/internal/semp"
)

type BrokerObjectType string

type GeneratorTerraformOutput struct {
	TerraformOutput  map[string]ResourceConfig
	SEMPDataResponse map[string]map[string]any
}

var BrokerObjectRelationship = map[BrokerObjectType][]BrokerObjectType{}
var DSLookup = map[BrokerObjectType]int{}

type BrokerRelationParameterPath struct {
	path          string
	terraformName string
}

var ObjectNamesCount = map[string]int{}

func GenerateAll(brokerURL string, context context.Context, cliClient *semp.Client, brokerResourceTerraformName string, brokerResourceName string, providerSpecificIdentifier string, fileName string) {
	// generatedResource := make(map[string]GeneratorTerraformOutput)

	// This will iterate all resources and genarete config for each

	// TODO: evaluate returning error from this function
	brokerResources, _ := fetchBrokerConfig(context, *cliClient, BrokerObjectType(brokerResourceTerraformName), brokerResourceName, providerSpecificIdentifier)
	// fetchBrokerConfig(context, *cliClient, BrokerObjectType(brokerResourceTerraformName), brokerResourceName, providerSpecificIdentifier)

	// // get all resources to be generated for
	// var resourcesToGenerate []BrokerObjectType
	// resourcesToGenerate = append(resourcesToGenerate, BrokerObjectType(brokerResourceTerraformName))
	// resourcesToGenerate = append(resourcesToGenerate, BrokerObjectRelationship[BrokerObjectType(brokerResourceTerraformName)]...)
	// for _, resource := range resourcesToGenerate {
	// 	generatedResults, generatedResourceChildren := generateForParentAndChildren(context, *cliClient, string(resource), brokerResourceName, providerSpecificIdentifier, generatedResource)
	// 	brokerResources = append(brokerResources, generatedResults...)
	// 	maps.Copy(generatedResource, generatedResourceChildren)
	// }

	LogCLIInfo("Replacing hardcoded names of inter-object dependencies by references where required")
	fixInterObjectDependencies(brokerResources)

	// Format the results
	object := &ObjectInfo{}
	object.BrokerResources = ToFormattedHCL(brokerResources)

	registry, ok := os.LookupEnv("SOLACEBROKER_REGISTRY_OVERRIDE")
	if !ok {
		registry = "registry.terraform.io"
	}
	object.Registry = registry
	object.BrokerURL = brokerURL
	object.Username = StringWithDefaultFromEnv("username", true, "")
	object.Password = StringWithDefaultFromEnv("password", false, "")
	if len(object.Password) == 0 {
		object.BearerToken = StringWithDefaultFromEnv("bearer_token", true, "")
	} else {
		object.BearerToken = StringWithDefaultFromEnv("bearer_token", false, "")
	}
	object.FileName = fileName

	LogCLIInfo("Found all resources. Writing file " + fileName)
	_ = GenerateTerraformFile(object)
	LogCLIInfo(fileName + " created successfully.\n")
}

func generateForParentAndChildren(context context.Context, client semp.Client, parentTerraformName string, brokerObjectInstanceName string, providerSpecificIdentifier string, generatedResources map[string]GeneratorTerraformOutput) ([]map[string]ResourceConfig, map[string]GeneratorTerraformOutput) {
	var brokerResources []map[string]ResourceConfig
	var generatorTerraformOutputForParent GeneratorTerraformOutput

	//get for parent
	_, alreadyGenerated := generatedResources[parentTerraformName]

	if !alreadyGenerated {
		generatorTerraformOutputForParent = ParseTerraformObject(context, client, brokerObjectInstanceName, parentTerraformName, providerSpecificIdentifier, map[string]string{}, map[string]any{})
		if len(generatorTerraformOutputForParent.TerraformOutput) > 0 {
			LogCLIInfo("Generating terraform config for " + parentTerraformName)
			resource := generatorTerraformOutputForParent.TerraformOutput
			brokerResources = append(brokerResources, resource)
			generatedResources[parentTerraformName] = generatorTerraformOutputForParent
		}
	} else {
		//pick output for generated data
		generatorTerraformOutputForParent = generatedResources[parentTerraformName]
	}

	childBrokerObjects := BrokerObjectRelationship[BrokerObjectType(parentTerraformName)]
	//get all children resources

	for _, childBrokerObject := range childBrokerObjects {

		_, alreadyGeneratedChild := generatedResources[string(childBrokerObject)]

		if !alreadyGeneratedChild {

			LogCLIInfo("Generating terraform config for " + string(childBrokerObject) + " as related to " + parentTerraformName)

			for key, parentBrokerResource := range generatorTerraformOutputForParent.TerraformOutput {

				parentResourceAttributes := map[string]ResourceConfig{}

				//use object name to build relationship
				parentResourceAttributes[key] = parentBrokerResource

				parentBrokerResourceAttributeRelationship := GetParentResourceAttributes(key, parentResourceAttributes)

				brokerResourcesToAppend := map[string]ResourceConfig{}

				//use parent semp response data to build semp request for children
				generatorTerraformOutputForChild := ParseTerraformObject(context, client, brokerObjectInstanceName,
					string(childBrokerObject),
					providerSpecificIdentifier,
					parentBrokerResourceAttributeRelationship,
					generatorTerraformOutputForParent.SEMPDataResponse[key])

				if len(generatorTerraformOutputForChild.TerraformOutput) > 0 {
					generatedResources[string(childBrokerObject)] = generatorTerraformOutputForChild
					for childBrokerResourceKey, childBrokerResourceValue := range generatorTerraformOutputForChild.TerraformOutput {
						if len(generatorTerraformOutputForChild.SEMPDataResponse[childBrokerResourceKey]) > 0 {
							//remove blanks
							if generatorTerraformOutputForChild.TerraformOutput[childBrokerResourceKey].ResourceAttributes != nil {
								brokerResourcesToAppend[childBrokerResourceKey] = childBrokerResourceValue
							}
						}
					}
					print("..")
					brokerResources = append(brokerResources, brokerResourcesToAppend)
				}
			}
		}
	}
	return brokerResources, generatedResources
}

func fixInterObjectDependencies(brokerResources []map[string]ResourceConfig) {
	// this will modify the passed brokerResources object

	//temporal hard coding dependency graph fix not available in SEMP API
	InterObjectDependencies := map[string][]string{"solacebroker_msg_vpn_authorization_group": {"solacebroker_msg_vpn_client_profile", "solacebroker_msg_vpn_acl_profile"},
		"solacebroker_msg_vpn_client_username":                            {"solacebroker_msg_vpn_client_profile", "solacebroker_msg_vpn_acl_profile"},
		"solacebroker_msg_vpn_rest_delivery_point":                        {"solacebroker_msg_vpn_client_profile"},
		"solacebroker_msg_vpn_acl_profile_client_connect_exception":       {"solacebroker_msg_vpn_acl_profile"},
		"solacebroker_msg_vpn_acl_profile_publish_topic_exception":        {"solacebroker_msg_vpn_acl_profile"},
		"solacebroker_msg_vpn_acl_profile_subscribe_share_name_exception": {"solacebroker_msg_vpn_acl_profile"},
		"solacebroker_msg_vpn_acl_profile_subscribe_topic_exception":      {"solacebroker_msg_vpn_acl_profile"}}

	ObjectNameAttributes := map[string]string{"solacebroker_msg_vpn_client_profile": "client_profile_name", "solacebroker_msg_vpn_acl_profile": "acl_profile_name"}

	// Post-process brokerResources for dependencies

	// For each resource check if there is any dependency
	for _, resources := range brokerResources {
		var resourceType string
		// var resourceConfig ResourceConfig
		for resourceKey := range resources {
			resourceType = strings.Split(resourceKey, " ")[0]
			resourceDependencies, exists := InterObjectDependencies[resourceType]
			if !exists {
				continue
			}
			// Found a resource that has inter-object relationship
			// fmt.Print("Found " + resourceKey + " with dependencies ")
			// fmt.Println(resourceDependencies)
			for _, dependency := range resourceDependencies {
				nameAttribute := ObjectNameAttributes[dependency]
				dependencyName := strings.Trim(resources[resourceKey].ResourceAttributes[nameAttribute].AttributeValue, "\"")
				if dependencyName != "" {
					// fmt.Println("   Dependency " + dependency + " name is " + dependencyName)
					// Look up key for dependency with dependencyName - iterate all brokerResources
					found := false
					for _, r := range brokerResources {
						for k := range r {
							rName := strings.Split(k, " ")[0]
							if rName != dependency {
								continue
							}
							// Check the name of the found resource
							if strings.Trim(r[k].ResourceAttributes[nameAttribute].AttributeValue, "\"") == dependencyName {
								// fmt.Println("         Found " + k + " as suitable dependency")
								// Replace hardcoded name by reference
								newInfo := ResourceAttributeInfo{
									AttributeValue: strings.Replace(k, " ", ".", -1) + "." + nameAttribute,
									Comment:        resources[resourceKey].ResourceAttributes[nameAttribute].Comment,
								}
								resources[resourceKey].ResourceAttributes[nameAttribute] = newInfo
								found = true
								break
							}
						}
						if found {
							break
						}
					}
				}
			}
		}
	}
}

func CreateBrokerObjectRelationships() {

	// Loop through entities and build database
	resourcesPathSignatureMap := map[string]string{}
	e := internalbroker.Entities
	for i, ds := range e {
		// Create new entry for each resource
		BrokerObjectRelationship[BrokerObjectType(ds.TerraformName)] = []BrokerObjectType{}
		DSLookup[BrokerObjectType(ds.TerraformName)] = i
		//// path := e[DSLookup[BrokerObjectType(ds.TerraformName)]].PathTemplate
		// Build a signature for each resource
		rex := regexp.MustCompile(`{[^\/]*}`)
		signature := strings.TrimSuffix(strings.Replace(rex.ReplaceAllString(ds.PathTemplate, ""), "//", "/", -1), "/") // Find all parameters in path template enclosed in {} including multiple ones
		if signature != "" {
			resourcesPathSignatureMap[signature] = ds.TerraformName
		}
	}
	// Loop through entities again and add children to parents
	for _, ds := range e {
		// Parent signature for each resource and add
		rex := regexp.MustCompile(`{[^\/]*}`)
		signature := strings.TrimSuffix(strings.Replace(rex.ReplaceAllString(ds.PathTemplate, ""), "//", "/", -1), "/")
		// get parentSignature by removing the part of signature after the last /
		parentSignature := path.Dir(signature)
		if parentSignature != "." && parentSignature != "/" {
			parentResource := resourcesPathSignatureMap[parentSignature]
			BrokerObjectRelationship[BrokerObjectType(parentResource)] = append(BrokerObjectRelationship[BrokerObjectType(parentResource)], BrokerObjectType(ds.TerraformName))
		}
	}
}

func ParseTerraformObject(ctx context.Context, client semp.Client, resourceName string, brokerObjectTerraformName string, providerSpecificIdentifier string, parentBrokerResourceAttributesRelationship map[string]string, parentResult map[string]any) GeneratorTerraformOutput {
	// var objectName string
	// tfObject := map[string]ResourceConfig{}
	// tfObjectSempDataResponse := map[string]map[string]any{}
	// entityToRead := internalbroker.EntityInputs{}
	// // TODO: potentially expensive
	// for _, ds := range internalbroker.Entities {
	// 	if strings.ToLower(ds.TerraformName) == strings.ToLower(brokerObjectTerraformName) {
	// 		entityToRead = ds
	// 		break
	// 	}
	// }
	// var path string

	// if len(parentResult) > 0 {
	// 	path, _ = ResolveSempPathWithParent(entityToRead.PathTemplate, parentResult)
	// } else {
	// 	path, _ = ResolveSempPath(entityToRead.PathTemplate, providerSpecificIdentifier)
	// }

	// if len(path) > 0 {

	// 	sempData, err := client.RequestWithoutBodyForGenerator(ctx, generated.BasePath, http.MethodGet, path, []map[string]any{})
	// 	if err != nil {
	// 		if err == semp.ErrResourceNotFound {
	// 			// continue if error is resource not found
	// 			if len(parentResult) > 0 {
	// 				print("..")
	// 			}
	// 			sempData = []map[string]any{}
	// 		} else if errors.Is(err, semp.ErrBadRequest) {
	// 			// continue if error is also bad request
	// 			if len(parentResult) > 0 {
	// 				print("..")
	// 			}
	// 			sempData = []map[string]any{}
	// 		} else {
	// 			ExitWithError("SEMP call failed. " + err.Error() + " on path " + path)
	// 		}
	// 	}

	// 	resourceKey := "solacebroker_" + brokerObjectTerraformName + " " + resourceName

	// 	resourceValues, err := GenerateTerraformString(entityToRead.Attributes, sempData, parentBrokerResourceAttributesRelationship, brokerObjectTerraformName)

	// 	//check resource names used and deduplicate to avoid collision
	// 	for i := range resourceValues {
	// 		totalOccurrence := 1
	// 		objectName = strings.ToLower(resourceKey) + GetNameForResource(strings.ToLower(resourceKey), resourceValues[i])
	// 		count, objectNameExists := ObjectNamesCount[objectName]
	// 		if objectNameExists {
	// 			totalOccurrence = count + 1
	// 		}
	// 		ObjectNamesCount[objectName] = totalOccurrence
	// 		objectName = objectName + "_" + fmt.Sprint(totalOccurrence)
	// 		tfObject[objectName] = resourceValues[i]
	// 		tfObjectSempDataResponse[objectName] = sempData[i]
	// 	}
	// }
	// return GeneratorTerraformOutput{
	// 	TerraformOutput:  tfObject,
	// 	SEMPDataResponse: tfObjectSempDataResponse,
	// }
	return GeneratorTerraformOutput{}
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
