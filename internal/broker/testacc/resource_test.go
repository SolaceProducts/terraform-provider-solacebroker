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

package acctest

import (
	"context"
	"testing"

	"terraform-provider-solacebroker/internal/broker"

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: ProviderConfig + `
resource "solacebroker_msg_vpn" "test" {
		msg_vpn_name = "test"
		enabled      = true
		max_msg_spool_usage = 5
		max_egress_flow_count = 997
		max_endpoint_count = 998
		max_ingress_flow_count = 999
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("solacebroker_msg_vpn.test", "msg_vpn_name", "test"),
					resource.TestCheckResourceAttr("solacebroker_msg_vpn.test", "max_msg_spool_usage", "5"),
				),
			},
			// Update and Read testing
			{
				Config: ProviderConfig + `
resource "solacebroker_msg_vpn" "test" {
		msg_vpn_name = "test"
		enabled      = true
		max_msg_spool_usage = 10
		max_egress_flow_count = 997
		max_endpoint_count = 998
		max_ingress_flow_count = 999
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("solacebroker_msg_vpn.test", "msg_vpn_name", "test"),
					resource.TestCheckResourceAttr("solacebroker_msg_vpn.test", "max_msg_spool_usage", "10"),
				),
			},
			// ImportState testing
			{
				ResourceName:                         "solacebroker_msg_vpn.test",
				ImportState:                          true,
				ImportStateId:                        "test",
				ImportStateVerifyIdentifierAttribute: "msg_vpn_name",
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore: []string{
					// These attributes need to be ignored from the test as they have broker-defaults and cannot be imported so that state will be null
					"max_connection_count",
					"max_kafka_broker_connection_count",
					"max_subscription_count",
					"max_transacted_session_count",
					"max_transaction_count",
					"service_amqp_max_connection_count",
					"service_mqtt_max_connection_count",
					"service_rest_incoming_max_connection_count",
					"service_rest_outgoing_max_connection_count",
					"service_smf_max_connection_count",
					"service_web_max_connection_count",
					"authentication_basic_profile_name",
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAllResourceSchemas(t *testing.T) {
	t.Parallel()

	for _, resource := range broker.Resources {
		ctx := context.Background()
		schemaRequest := fwresource.SchemaRequest{}
		schemaResponse := &fwresource.SchemaResponse{}

		// Instantiate the resource.Resource and call its Schema method
		resource().Schema(ctx, schemaRequest, schemaResponse)

		if schemaResponse.Diagnostics.HasError() {
			t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
		}

		// Validate the schema
		diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

		if diagnostics.HasError() {
			t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
		}
	}
}
