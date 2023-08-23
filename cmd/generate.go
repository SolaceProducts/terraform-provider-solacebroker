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
package cmd

import (
	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strings"
	"terraform-provider-solacebroker/cmd/broker"
	command "terraform-provider-solacebroker/cmd/command"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a Terraform configuration file for a specified object and all child objects known to the provider",
	Long: `The generate command on the provider binary generates a Terraform configuration file for the specified object and all child objects known to the provider.
Generate is not a Terraform command (Terraform does not have the concept).  It is a broker provider command.  If a user has a copy of the provider binary (from either an automatic Terraform download or by downloading it directly), they can execute that binary with the generate command to generate a Terraform configuration file from the current configuration of a broker.
<binary> generate [options] <terraform resource address> <provider-specific identifier> <filename>
where <binary> is the broker provider binary, <terraform resource address> and <provider-specific identifier> are the same as for the Terraform Import command, <filename> is the desirable name of the generated filename (would general end with the standard Terraform extension of .tf), and [options] are the supported options, which mirror the configuration options for the provider object (for example -url=https://f93.soltestlab.ca:1943 and -retry_wait_max=90s) and can be set via environment variables in the same way and with the same validations.
For example: 
terraform-provider-solacebroker generate -url=https://f93.soltestlab.ca:1943 solacebroker_msg_vpn_rest_delivery_point.my_rdp default/my-rdp my-rdp.tf`,
	Run: func(cmd *cobra.Command, args []string) {
		brokerURL, _ := cmd.Flags().GetString("url")
		command.LogCLIInfo("Connecting to Broker : " + brokerURL)

		client := broker.CliClient(brokerURL)
		if client == nil {
			command.LogCLIError("Error creating SEMP Client")
			os.Exit(1)
		}
		command.LogCLIInfo("Connection successful")

		brokerObjectType := cmd.Flags().Arg(0)

		if len(brokerObjectType) == 0 {
			command.LogCLIError("Terraform resource name not provided")
			cmd.Help()
			os.Exit(1)
		}
		providerSpecificIdentifier := cmd.Flags().Arg(1)
		if len(providerSpecificIdentifier) == 0 {
			command.LogCLIError("Broker object  not provided")
			cmd.Help()
			os.Exit(1)
		}

		fileName := cmd.Flags().Arg(2)
		if len(fileName) == 0 {
			command.LogCLIError("\nError: Terraform file name not specified.\n\n")
			cmd.Help()
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
		brokerResources := []map[string]string{}

		//get parent resource
		brokerResources = append(brokerResources, command.ParseTerraformObject(cmd.Context(), *client, brokerObjectInstanceName, brokerObjectTerraformName, providerSpecificIdentifier))

		//get all children resources
		childBrokerObjects := command.BrokerObjectRelationship[command.BrokerObjectType(brokerObjectTerraformName)]
		for _, childBrokerObject := range childBrokerObjects {
			brokerResources = append(brokerResources, command.ParseTerraformObject(cmd.Context(), *client, brokerObjectInstanceName, string(childBrokerObject), providerSpecificIdentifier))
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
		command.GenerateTerraformFile(object)
		command.LogCLIInfo(fileName + " created successfully")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("url", "http://localhost:8080", "Broker URL")
}
