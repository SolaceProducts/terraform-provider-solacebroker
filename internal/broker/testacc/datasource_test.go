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

	fwredatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + `
data "solacebroker_msg_vpn" "default" {
		msg_vpn_name = "default"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.solacebroker_msg_vpn.default", "authentication_basic_enabled", "true"),
				),
			},
		},
	})
}

func TestAllDataSourceSchemas(t *testing.T) {
	t.Parallel()

	for _, dataSource := range broker.DataSources {
		ctx := context.Background()
		schemaRequest := fwredatasource.SchemaRequest{}
		schemaResponse := &fwredatasource.SchemaResponse{}

		// Instantiate the resource.Resource and call its Schema method
		dataSource().Schema(ctx, schemaRequest, schemaResponse)

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
