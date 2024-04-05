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
	"fmt"
	"net/http"
	"os"
	"strings"
	"terraform-provider-solacebroker/cmd/client"
	"terraform-provider-solacebroker/cmd/generator"
	"terraform-provider-solacebroker/internal/broker/generated"

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

		generator.LogCLIInfo(fmt.Sprintf("Attempting config generation for object and its child-objects: %s, identifier: %s, destination file: %s\n", brokerObjectType, providerSpecificIdentifier, fileName))

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

		generator.GenerateAll(brokerURL, cmd.Context(), cliClient, brokerResourceTerraformName, brokerResourceName, providerSpecificIdentifier, fileName)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("url", "http://localhost:8080", "Broker URL")
}
