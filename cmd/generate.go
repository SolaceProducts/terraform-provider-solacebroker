// Package cmd terraform-provider-solacebroker
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
package cmd

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"os"
	"strings"
	"terraform-provider-solacebroker/cmd/client"
	"terraform-provider-solacebroker/cmd/generator"
	"terraform-provider-solacebroker/internal/broker/generated"
	"terraform-provider-solacebroker/internal/semp"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [options] <terraform resource address> <provider-specific identifier> <filename>",
	Short: "Generates a Terraform configuration file for a specified PubSub+ event broker object and all child objects known to the provider",
	Long: `The generate command on the provider binary generates a Terraform configuration file for the specified object and all child objects known to the provider.
This is not a Terraform generator. One can download the provider binary and can execute that binary with the "generate" command to generate a Terraform configuration file from the current configuration of a PubSub+ event broker.

  <binary> generate [options] <terraform resource address> <provider-specific identifier> <filename>

  where;
		<binary> is the broker provider binary
		[options] are the supported options, which mirror the configuration options for the provider object (for example -url=https://f93.soltestlab.ca:1943 and -retry_wait_max=90s) and can be set via environment variables in the same way.
		<terraform resource address> addresses a specific resource instance in the form of <resource_type>.<resource_name>
		<provider-specific identifier> the identifier of the broker object, the same as for the Terraform Import generator.
		<filename> is the name of the generated file

Example:
  SOLACEBROKER_USERNAME=adminuser SOLACEBROKER_PASSWORD=pass \
	terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn.mq default my-messagevpn.tf

This command will create a file my-messagevpn.tf that contains a resource definition for the default message VPN and any child objects, assuming the appropriate broker credentials were set in environment variables.`,

	Run: func(cmd *cobra.Command, args []string) {
		brokerURL, _ := cmd.Flags().GetString("url")
		generator.LogCLIInfo("Connecting to Broker : " + brokerURL)

		cliClient := client.CliClient(brokerURL)
		if cliClient == nil {
			generator.ExitWithError("Error creating SEMP Client")
		}

		brokerObjectType := cmd.Flags().Arg(0)

		if len(brokerObjectType) == 0 {
			generator.LogCLIError("Terraform resource name not provided")
			_ = cmd.Help()
			os.Exit(1)
		}
		providerSpecificIdentifier := cmd.Flags().Arg(1)
		if len(providerSpecificIdentifier) == 0 {
			generator.LogCLIError("Broker object not provided")
			_ = cmd.Help()
			os.Exit(1)
		}

		fileName := cmd.Flags().Arg(2)
		if len(fileName) == 0 {
			generator.LogCLIError("\nError: Terraform file name not specified.\n\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		if !strings.HasSuffix(fileName, ".tf") {
			fileName = fileName + ".tf"
		}

		skipApiCheck, err := generator.BooleanWithDefaultFromEnv("skip_api_check", false, false)
		if err != nil {
			generator.ExitWithError("\nError: Unable to parse provider attribute. " + err.Error())
		}
		//Confirm SEMP version and connection via client
		aboutPath := "/about/api"
		result, err := cliClient.RequestWithoutBody(cmd.Context(), http.MethodGet, aboutPath)
		if err != nil {
			generator.ExitWithError("SEMP call failed. " + err.Error())
		}
		brokerSempVersion := result["sempVersion"].(string)
		brokerPlatform := result["platform"].(string)
		if !skipApiCheck && brokerPlatform != generated.Platform {
			generator.ExitWithError(fmt.Sprintf("Broker platform \"%s\" does not match generator supported platform: %s", BrokerPlatformName[brokerPlatform], BrokerPlatformName[generated.Platform]))
		}
		generator.LogCLIInfo("Connection successful.")
		generator.LogCLIInfo(fmt.Sprintf("Broker SEMP version is %s, Generator SEMP version is %s", brokerSempVersion, generated.SempVersion))

		generator.LogCLIInfo("Attempting config generation for object and its child-objects: " + brokerObjectType + ", identifier: " + providerSpecificIdentifier + ", destination file: " + fileName)

		object := &generator.ObjectInfo{}
		// Extract and verify parameters
		if strings.Count(brokerObjectType, ".") != 1 {
			generator.ExitWithError("\nError: Terraform resource address is not in correct format. Should be in the format <resource_type>.<resource_name>\n\n")
		}
		brokerResourceType := strings.Split(brokerObjectType, ".")[0]
		brokerResourceName := strings.Split(brokerObjectType, ".")[1]
		if !generator.IsValidTerraformIdentifier(brokerResourceName) {
			generator.ExitWithError(fmt.Sprintf("\nError: Resource name %s in the Terraform resource address is not a valid Terraform identifier\n\n", brokerResourceName))
		}

		brokerResourceTerraformName := strings.ReplaceAll(brokerResourceType, "solacebroker_", "")

		_, found := generator.BrokerObjectRelationship[generator.BrokerObjectType(brokerResourceTerraformName)]
		if !found {
			generator.ExitWithError("\nError: Broker resource not found by terraform name : " + brokerResourceTerraformName + "\n\n")
		}
		generatedResource := make(map[string]generator.GeneratorTerraformOutput)
		var brokerResources []map[string]generator.ResourceConfig

		// This will iterate all resources and genarete config for each

		// TODO: evaluate returning error from this function
		generateConfigForObjectInstances(cmd.Context(), *cliClient, generator.BrokerObjectType(brokerResourceTerraformName), providerSpecificIdentifier, nil)

		// get all resources to be generated for
		var resourcesToGenerate []generator.BrokerObjectType
		resourcesToGenerate = append(resourcesToGenerate, generator.BrokerObjectType(brokerResourceTerraformName))
		resourcesToGenerate = append(resourcesToGenerate, generator.BrokerObjectRelationship[generator.BrokerObjectType(brokerResourceTerraformName)]...)
		for _, resource := range resourcesToGenerate {
			generatedResults, generatedResourceChildren := generateForParentAndChildren(cmd.Context(), *cliClient, string(resource), brokerResourceName, providerSpecificIdentifier, generatedResource)
			brokerResources = append(brokerResources, generatedResults...)
			maps.Copy(generatedResource, generatedResourceChildren)
		}

		generator.LogCLIInfo("Replacing hardcoded names of inter-object dependencies by references where required")
		fixInterObjectDependencies(brokerResources)

		// Format the results
		object.BrokerResources = generator.ToFormattedHCL(brokerResources)

		registry, ok := os.LookupEnv("SOLACEBROKER_REGISTRY_OVERRIDE")
		if !ok {
			registry = "registry.terraform.io"
		}
		object.Registry = registry
		object.BrokerURL = brokerURL
		object.Username = generator.StringWithDefaultFromEnv("username", true, "")
		object.Password = generator.StringWithDefaultFromEnv("password", false, "")
		if len(object.Password) == 0 {
			object.BearerToken = generator.StringWithDefaultFromEnv("bearer_token", true, "")
		} else {
			object.BearerToken = generator.StringWithDefaultFromEnv("bearer_token", false, "")
		}
		object.FileName = fileName

		generator.LogCLIInfo("Found all resources. Writing file " + fileName)
		_ = generator.GenerateTerraformFile(object)
		generator.LogCLIInfo(fileName + " created successfully.\n")
		os.Exit(0)
	},
}

func generateForParentAndChildren(context context.Context, client semp.Client, parentTerraformName string, brokerObjectInstanceName string, providerSpecificIdentifier string, generatedResources map[string]generator.GeneratorTerraformOutput) ([]map[string]generator.ResourceConfig, map[string]generator.GeneratorTerraformOutput) {
	var brokerResources []map[string]generator.ResourceConfig
	var generatorTerraformOutputForParent generator.GeneratorTerraformOutput

	//get for parent
	_, alreadyGenerated := generatedResources[parentTerraformName]

	if !alreadyGenerated {
		generatorTerraformOutputForParent = generator.ParseTerraformObject(context, client, brokerObjectInstanceName, parentTerraformName, providerSpecificIdentifier, map[string]string{}, map[string]any{})
		if len(generatorTerraformOutputForParent.TerraformOutput) > 0 {
			generator.LogCLIInfo("Generating terraform config for " + parentTerraformName)
			resource := generatorTerraformOutputForParent.TerraformOutput
			brokerResources = append(brokerResources, resource)
			generatedResources[parentTerraformName] = generatorTerraformOutputForParent
		}
	} else {
		//pick output for generated data
		generatorTerraformOutputForParent = generatedResources[parentTerraformName]
	}

	childBrokerObjects := generator.BrokerObjectRelationship[generator.BrokerObjectType(parentTerraformName)]
	//get all children resources

	for _, childBrokerObject := range childBrokerObjects {

		_, alreadyGeneratedChild := generatedResources[string(childBrokerObject)]

		if !alreadyGeneratedChild {

			generator.LogCLIInfo("Generating terraform config for " + string(childBrokerObject) + " as related to " + parentTerraformName)

			for key, parentBrokerResource := range generatorTerraformOutputForParent.TerraformOutput {

				parentResourceAttributes := map[string]generator.ResourceConfig{}

				//use object name to build relationship
				parentResourceAttributes[key] = parentBrokerResource

				parentBrokerResourceAttributeRelationship := generator.GetParentResourceAttributes(key, parentResourceAttributes)

				brokerResourcesToAppend := map[string]generator.ResourceConfig{}

				//use parent semp response data to build semp request for children
				generatorTerraformOutputForChild := generator.ParseTerraformObject(context, client, brokerObjectInstanceName,
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

func fixInterObjectDependencies(brokerResources []map[string]generator.ResourceConfig) {
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
		// var resourceConfig generator.ResourceConfig
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
								newInfo := generator.ResourceAttributeInfo{
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

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("url", "http://localhost:8080", "Broker URL")
}
