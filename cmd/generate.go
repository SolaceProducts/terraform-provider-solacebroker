// Package cmd terraform-provider-solacebroker
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
package cmd

import (
	"context"
	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"net/http"
	"os"
	"strings"
	"terraform-provider-solacebroker/cmd/broker"
	command "terraform-provider-solacebroker/cmd/command"
	"terraform-provider-solacebroker/internal/semp"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate --url=<terraform resource address> <provider-specific identifier> <filename>",
	Short: "Generates a Terraform configuration file for a specified PubSubPlus Broker object and all child objects known to the provider",
	Long: `The generate command on the provider binary generates a Terraform configuration file for the specified object and all child objects known to the provider.
This is not a Terraform command. One can download the provider binary and can execute that binary with the "generate" command to generate a Terraform configuration file from the current configuration of a PubSubPlus event broker.

 <binary> generate <terraform resource address> <provider-specific identifier> <filename>

 where;
	<binary> is the broker provider binary,
	<terraform resource address> is the terraform resource address, for example http://localhost:8080,
	<provider-specific identifier> are the similar to the Terraform Import command,this is the resource name and possible values to find a specific resource,
	<filename> is the desirable name of the generated filename.

For example:
  terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn.mq default my-messagevpn.tf

This command would create a file my-messagevpn.tf that contains a resource definition for the default message VPN and any child objects, assuming the appropriate broker credentials were set in environment variables.`,

	Run: func(cmd *cobra.Command, args []string) {
		brokerURL, _ := cmd.Flags().GetString("url")
		command.LogCLIInfo("Connecting to Broker : " + brokerURL)

		client := broker.CliClient(brokerURL)
		if client == nil {
			command.LogCLIError("Error creating SEMP Client")
			os.Exit(1)
		}

		brokerObjectType := cmd.Flags().Arg(0)

		if len(brokerObjectType) == 0 {
			command.LogCLIError("Terraform resource name not provided")
			_ = cmd.Help()
			os.Exit(1)
		}
		providerSpecificIdentifier := cmd.Flags().Arg(1)
		if len(providerSpecificIdentifier) == 0 {
			command.LogCLIError("Broker object  not provided")
			_ = cmd.Help()
			os.Exit(1)
		}

		fileName := cmd.Flags().Arg(2)
		if len(fileName) == 0 {
			command.LogCLIError("\nError: Terraform file name not specified.\n\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		if !strings.HasSuffix(fileName, ".tf") {
			fileName = fileName + ".tf"
		}

		//Confirm SEMP version and connection via client
		aboutPath := "/about/api"
		result, err := client.RequestWithoutBody(cmd.Context(), http.MethodGet, aboutPath)
		if err != nil {
			command.LogCLIError("SEMP call failed. " + err.Error())
			os.Exit(1)
		}
		brokerSempVersion, err := version.NewVersion(result["sempVersion"].(string))
		if err != nil {
			command.LogCLIError("Unable to parse SEMP version from API")
			os.Exit(1)
		}
		command.LogCLIInfo("Connection successful")
		command.LogCLIInfo("Broker SEMP version is " + brokerSempVersion.String())

		command.LogCLIInfo("Attempt generation for broker object: " + brokerObjectType + " of " + providerSpecificIdentifier + " in file " + fileName)

		object := &command.ObjectInfo{}

		brokerObjectTypeName := brokerObjectType
		brokerObjectInstanceName := strings.ToLower(brokerObjectType)
		if strings.Contains(brokerObjectType, ".") {
			brokerObjectTypeName = strings.Split(brokerObjectType, ".")[0]
			brokerObjectInstanceName = strings.Split(brokerObjectType, ".")[1]
		}

		brokerObjectTerraformName := strings.ReplaceAll(brokerObjectTypeName, "solacebroker_", "")

		_, found := command.BrokerObjectRelationship[command.BrokerObjectType(brokerObjectTerraformName)]
		if !found {
			command.LogCLIError("\nError: Broker resource not found by terraform name : " + brokerObjectTerraformName + "\n\n")
			os.Exit(1)
		}
		generatedResource := make(map[string]command.GeneratorTerraformOutput)
		var brokerResources []map[string]string

		// get all resources to be generated for
		var resourcesToGenerate []command.BrokerObjectType
		resourcesToGenerate = append(resourcesToGenerate, command.BrokerObjectType(brokerObjectTerraformName))
		resourcesToGenerate = append(resourcesToGenerate, command.BrokerObjectRelationship[command.BrokerObjectType(brokerObjectTerraformName)]...)
		for _, resource := range resourcesToGenerate {
			generatedResults, generatedResourceChildren := generateForParentAndChildren(cmd.Context(), *client, string(resource), brokerObjectInstanceName, providerSpecificIdentifier, generatedResource)
			brokerResources = append(brokerResources, generatedResults...)
			maps.Copy(generatedResource, generatedResourceChildren)
		}

		object.BrokerResources = brokerResources

		registry, ok := os.LookupEnv("SOLACEBROKER_REGISTRY_OVERRIDE")
		if !ok {
			registry = "registry.terraform.io"
		}
		object.Registry = registry
		object.BrokerURL = brokerURL
		object.Username = command.StringWithDefaultFromEnv("username", true, "")
		object.Password = command.StringWithDefaultFromEnv("password", false, "")
		if len(object.Password) == 0 {
			object.BearerToken = command.StringWithDefaultFromEnv("bearer_token", true, "")
		} else {
			object.BearerToken = command.StringWithDefaultFromEnv("bearer_token", false, "")
		}
		object.FileName = fileName

		command.LogCLIInfo("Found all resources. Generation started for file " + fileName)
		_ = command.GenerateTerraformFile(object)
		command.LogCLIInfo(fileName + " created successfully.")
		os.Exit(0)
	},
}

func generateForParentAndChildren(context context.Context, client semp.Client, parentTerraformName string, brokerObjectInstanceName string, providerSpecificIdentifier string, generatedResources map[string]command.GeneratorTerraformOutput) ([]map[string]string, map[string]command.GeneratorTerraformOutput) {
	var brokerResources []map[string]string
	var generatorTerraformOutputForParent command.GeneratorTerraformOutput

	//get for parent
	_, alreadyGenerated := generatedResources[parentTerraformName]

	if !alreadyGenerated {
		generatorTerraformOutputForParent = command.ParseTerraformObject(context, client, brokerObjectInstanceName, parentTerraformName, providerSpecificIdentifier, map[string]string{}, map[string]any{})
		if len(generatorTerraformOutputForParent.TerraformOutput) > 0 {
			command.LogCLIInfo("Generating terraform config for " + parentTerraformName)
			brokerResources = append(brokerResources, generatorTerraformOutputForParent.TerraformOutput)
			generatedResources[parentTerraformName] = generatorTerraformOutputForParent
		}
	} else {
		//pick output for generated data
		generatorTerraformOutputForParent = generatedResources[parentTerraformName]
	}

	childBrokerObjects := command.BrokerObjectRelationship[command.BrokerObjectType(parentTerraformName)]
	//get all children resources

	for _, childBrokerObject := range childBrokerObjects {

		_, alreadyGeneratedChild := generatedResources[string(childBrokerObject)]

		if !alreadyGeneratedChild {

			command.LogCLIInfo("Generating terraform config for " + string(childBrokerObject) + " as related to " + parentTerraformName)

			for key, parentBrokerResource := range generatorTerraformOutputForParent.TerraformOutput {

				parentResourceAttributes := map[string]string{}

				//use object name to build relationship
				parentResourceAttributes[key] = parentBrokerResource

				parentBrokerResourceAttributeRelationship := command.GetParentResourceAttributes(key, parentResourceAttributes)

				brokerResourcesToAppend := map[string]string{}

				//use parent semp response data to build semp request for children
				generatorTerraformOutputForChild := command.ParseTerraformObject(context, client, brokerObjectInstanceName,
					string(childBrokerObject),
					providerSpecificIdentifier,
					parentBrokerResourceAttributeRelationship,
					generatorTerraformOutputForParent.SEMPDataResponse[key])

				if len(generatorTerraformOutputForChild.TerraformOutput) > 0 {
					generatedResources[string(childBrokerObject)] = generatorTerraformOutputForChild
					for childBrokerResourceKey, childBrokerResourceValue := range generatorTerraformOutputForChild.TerraformOutput {
						if len(generatorTerraformOutputForChild.SEMPDataResponse[childBrokerResourceKey]) > 0 {
							//remove blanks
							if len(generatorTerraformOutputForChild.TerraformOutput[childBrokerResourceKey]) > 0 {
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

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("url", "http://localhost:8080", "Broker URL")
}
