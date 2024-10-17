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
	"terraform-provider-solacebroker/internal/semp"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate <terraform resource address> <provider-specific identifier> <filename>",
	Short: "Generates a Terraform configuration file for a specified PubSub+ event broker object and all child objects known to the provider",
	Long: `The generate command on the provider binary generates a Terraform configuration file for the specified object and all child objects known to the provider.
This is not a Terraform command. You can download the provider binary and execute that binary with the "generate" command to generate a Terraform configuration file from the current configuration of a PubSub+ event broker.

  <binary> generate [flags] <terraform resource address> <provider-specific identifier> <filename>

  where:
		<binary> is the broker provider binary
		[flags] are the supported options, which mirror the configuration options for the provider object (for example --url=https://localhost:1943 and --retry_wait_max=90s) and can also be set via environment variables in the same way.
		<terraform resource address> how to address the specified object instance in the generated configuration, in the form of <resource_type>.<resource_name>
		<provider-specific identifier> the import identifier of the specified object instance, refer to the resource type of the object in the provider documentation
		<filename> is the name of the generated file

Example:
  SOLACEBROKER_USERNAME=adminuser SOLACEBROKER_PASSWORD=pass \
	terraform-provider-solacebroker generate --url=http://localhost:8080 solacebroker_msg_vpn.myvpn test vpn-config.tf

This command will create a file vpn-config.tf that contains a resource definition for the 'test' message VPN and any child objects on the broker, assuming the appropriate broker credentials were set in environment variables.
The message VPN resource address in the generated configuration will be 'solacebroker_msg_vpn.myvpn'.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			// Print the help message if the required arguments are not provided
			_ = cmd.Help()
			os.Exit(1)
		}

		flags := cmd.Flags()
		cliParams := generator.CliParams{}
		if flags.Changed("url") {
			if url, err := flags.GetString("url"); err == nil {
				cliParams.Url = &url
			}
		}
		if flags.Changed("username") {
			if username, err := flags.GetString("username"); err == nil {
				cliParams.Username = &username
			}
		}
		if flags.Changed("password") {
			if password, err := flags.GetString("password"); err == nil {
				cliParams.Password = &password
			}
		}
		if flags.Changed("bearer_token") {
			if bearerToken, err := flags.GetString("bearer_token"); err == nil {
				cliParams.Bearer_token = &bearerToken
			}
		}
		if flags.Changed("retries") {
			if retries, err := flags.GetInt64("retries"); err == nil {
				cliParams.Retries = &retries
			}
		}
		if flags.Changed("retry_min_interval") {
			if retryMinInterval, err := flags.GetDuration("retry_min_interval"); err == nil {
				cliParams.Request_min_interval = &retryMinInterval
			}
		}
		if flags.Changed("retry_max_interval") {
			if retryMaxInterval, err := flags.GetDuration("retry_max_interval"); err == nil {
				cliParams.Retry_max_interval = &retryMaxInterval
			}
		}
		if flags.Changed("request_timeout_duration") {
			if requestTimeoutDuration, err := flags.GetDuration("request_timeout_duration"); err == nil {
				cliParams.Request_timeout_duration = &requestTimeoutDuration
			}
		}
		if flags.Changed("request_min_interval") {
			if requestMinInterval, err := flags.GetDuration("request_min_interval"); err == nil {
				cliParams.Request_min_interval = &requestMinInterval
			}
		}
		if flags.Changed("insecure_skip_verify") {
			if insecureSkipVerify, err := flags.GetBool("insecure_skip_verify"); err == nil {
				cliParams.Insecure_skip_verify = &insecureSkipVerify
			}
		}
		if flags.Changed("skip_api_check") {
			if skipApiCheck, err := flags.GetBool("skip_api_check"); err == nil {
				cliParams.Skip_api_check = &skipApiCheck
			}
		}
		// Complement params with env as required, also ensure valid values for all
		cliParams = generator.UpdateCliParamsWithEnv(cliParams)

		cliClient := client.CliClient(cliParams)
		if cliClient == nil {
			generator.ExitWithError("Error creating SEMP Client")
		}

		brokerObjectType := flags.Arg(0)

		if len(brokerObjectType) == 0 {
			generator.LogCLIError("Terraform resource name not provided")
			_ = cmd.Help()
			os.Exit(1)
		}
		providerSpecificIdentifier := flags.Arg(1)
		if len(providerSpecificIdentifier) == 0 {
			generator.LogCLIError("Broker object not provided")
			_ = cmd.Help()
			os.Exit(1)
		}

		fileName := flags.Arg(2)
		if len(fileName) == 0 {
			generator.LogCLIError("\nError: Terraform file name not specified.\n\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		if !strings.HasSuffix(fileName, ".tf") {
			fileName = fileName + ".tf"
		}

		skipApiCheck := *cliParams.Skip_api_check
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
		generator.GenerateAll(cliParams, cmd.Context(), cliClient, brokerResourceTerraformName, brokerResourceName, providerSpecificIdentifier, fileName)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("url", "http://localhost:8080", "Broker base URL, for example https://mybroker.example.org:<semp-service-port>")
	generateCmd.PersistentFlags().String("username", "", "Basic authentication username")
	generateCmd.PersistentFlags().String("password", "", "Basic authentication password")
	generateCmd.PersistentFlags().String("bearer_token", "", "Bearer token for authentication")
	generateCmd.PersistentFlags().Int64("retries", semp.DefaultRetries, "Retries")
	generateCmd.PersistentFlags().Duration("retry_min_interval", semp.DefaultRetryMinInterval, "Minimum retry interval")
	generateCmd.PersistentFlags().Duration("retry_max_interval", semp.DefaultRetryMaxInterval, "Maximum retry interval")
	generateCmd.PersistentFlags().Duration("request_timeout_duration", semp.DefaultRequestTimeout, "Request timeout duration")
	generateCmd.PersistentFlags().Duration("request_min_interval", semp.DefaultRequestInterval, "Minimum request interval")
	generateCmd.PersistentFlags().Bool("insecure_skip_verify", false, "Disable validation of server SSL certificates")
	generateCmd.PersistentFlags().Bool("skip_api_check", false, "Disable validation of the broker SEMP API")
}
