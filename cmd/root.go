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
This is not a Terraform command.  It is a broker provider command.  If a user has a copy of the provider binary (from either an automatic Terraform download or by downloading it directly), they can execute that binary with the generate command to generate a Terraform configuration file from the current configuration of a broker.

<binary> generate [options] <terraform resource address> <provider-specific identifier> <filename>
where <binary> is the broker provider binary, <terraform resource address> and <provider-specific identifier> are the same as for the Terraform Import command, <filename> is the desirable name of the generated filename (would general end with the standard Terraform extension of .tf), and [options] are the supported options, which mirror the configuration options for the provider object (for example -url=https://f93.soltestlab.ca:1943 and -retry_wait_max=90s) and can be set via environment variables in the same way and with the same validations.
For example:

  terraform-provider-solacebroker generate -url=https://localhost:8080 solacebroker_msg_vpn.my_rdp default/my-rdp my-rdp.tf

This command would create a file my-rdp.tf that contained a resource definition for the my-rdp RDP and any child objects (probably a REST consumer and a queue binding), assuming the appropriate broker credentials were set in environment variables.`,
}

func Execute() error {
	terraform.CreateBrokerObjectRelationships()
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
