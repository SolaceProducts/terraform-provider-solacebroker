// terraform-provider-solacebroker
//
// Copyright 2025 Solace Corporation. All rights reserved.
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

package generated

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "virtual_hostname",
		MarkdownDescription: "A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.\n\nClients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.\n\nThis mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.\n\n\n\nThe minimum access scope/level required to perform this operation is \"global/read-only\".\n\nThis has been available since SEMP API version 2.17.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/virtualHostnames/{virtualHostname}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable Virtual Hostname to Message VPN mapping.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/read-write\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "msgVpnName",
				TerraformName:       "msg_vpn_name",
				MarkdownDescription: "The message VPN to which this virtual hostname is mapped.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/read-write\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?]*$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "virtualHostname",
				TerraformName:       "virtual_hostname",
				MarkdownDescription: "The virtual hostname.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 253),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
