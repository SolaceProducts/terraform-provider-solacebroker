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
	"path"
	"regexp"
	"strings"
	internalbroker "terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/semp"
)

type BrokerObjectType string

type ResourceAttributeInfo struct {
	AttributeValue string
	Comment        string
}

type ResourceConfig struct {
	ResourceAttributes map[string]ResourceAttributeInfo // indexed by resource attribute name
}

type VariableConfig struct {
	Type      string
	Default   string
	Sensitive bool
}

type ObjectInfo struct {
	BasicAuthentication bool
	FileName            string
	BrokerResources     []map[string]string
	Variables           map[string]VariableConfig
}

var BrokerObjectRelationship = map[BrokerObjectType][]BrokerObjectType{}
var DSLookup = map[BrokerObjectType]int{} // Helper to easily lookup an entity in internalbroker.Entities by name

var ObjectNamesCount = map[string]int{}

func GenerateAll(cliParams CliParams, context context.Context, cliClient *semp.Client, brokerResourceTerraformName string, brokerResourceName string, providerSpecificIdentifier string, fileName string) {
	// First build the parent-child relationship between broker objects
	CreateBrokerObjectRelationships()

	// Check if the broker resource is supported
	_, found := BrokerObjectRelationship[BrokerObjectType(brokerResourceTerraformName)]
	if !found {
		ExitWithError("\nError: Broker resource not found by terraform name : " + brokerResourceTerraformName + "\n\n")
	}

	// This will iterate all resources starting at brokerResourceTerraformName and genarete brokerResources and variables config for that and children
	brokerResources, variables, err := fetchBrokerConfig(context, *cliClient, BrokerObjectType(brokerResourceTerraformName), brokerResourceName, providerSpecificIdentifier)
	if err != nil {
		ExitWithError("Failed to fetch broker config, " + err.Error())
	}

	// Postprocess brokerResources for dependencies
	LogCLIInfo("Replacing hardcoded names of inter-object dependencies by references where required")
	fixInterObjectDependencies(brokerResources)

	// Prep to generate the Terraform file
	object := &ObjectInfo{}
	object.BrokerResources = resourcesToFormattedHCL(brokerResources)
	object.Variables = variables
	object.BasicAuthentication = (*cliParams.Username != "" && *cliParams.Bearer_token == "")
	object.FileName = fileName
	LogCLIInfo("Found all resources. Writing file " + fileName)

	// Generate the Terraform file
	err = GenerateTerraformFile(object)
	if err != nil {
		ExitWithError("Failed to write file, " + err.Error())
	}
	LogCLIInfo(fileName + " created successfully.\n")
}

func CreateBrokerObjectRelationships() {
	// Loop through entities and build database
	resourcesPathSignatureMap := map[string]string{}
	e := internalbroker.Entities
	for i, ds := range e {
		// Create new entry for each resource
		BrokerObjectRelationship[BrokerObjectType(ds.TerraformName)] = []BrokerObjectType{}
		DSLookup[BrokerObjectType(ds.TerraformName)] = i
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
