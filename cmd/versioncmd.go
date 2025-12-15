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
package cmd

import (
	"fmt"
	"terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/broker/generated"

	"github.com/spf13/cobra"
)

var (
	BrokerPlatformName = map[string]string{
		"VMR":       "Software Event Broker",
		"Appliance": "Appliance",
	}
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Provides version information about the current binary",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Terraform Provider for Solace %s platform, version: %s, based on Semp version %s\n", BrokerPlatformName[generated.Platform], broker.ProviderVersion, generated.SempVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
