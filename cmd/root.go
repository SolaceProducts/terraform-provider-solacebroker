// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	terraform "terraform-provider-solacebroker/cmd/command"
)

var rootCmd = &cobra.Command{
	Use:   "terraform-provider-solacebroker",
	Short: "Generates a Terraform configuration file for a specified PubSubPlus Broker object and all child objects known to the provider",
	Long: `The generate command on the provider binary generates a Terraform configuration file for the specified object and all child objects known to the provider.
This is not a Terraform command. One can download the provider binary and can execute that binary with the "generate" command to generate a Terraform configuration file from the current configuration of a PubSubPlus broker..

 <binary> generate <terraform resource address> <provider-specific identifier> <filename>

 where;
	<binary> is the broker provider binary,
	<terraform resource address> is the terraform resource address, for example https://mybroker.example.org:1943/,
	<provider-specific identifier> are the similar to the Terraform Import command,this is the resource name and possible values to find a specific resource,
	<filename> is the desirable name of the generated filename.

For example:
  terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn.mq default my-messagevpn.tf

This command would create a file my-messagevpn.tf that contains a resource definition for the default message VPN and any child objects, assuming the appropriate broker credentials were set in environment variables.`,
}

func Execute() error {
	terraform.CreateBrokerObjectRelationships()
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
