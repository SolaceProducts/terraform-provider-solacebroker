// terraform-provider-solacebroker
//
// Copyright 2026 Solace Corporation. All rights reserved.
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
		TerraformName:       "msg_vpn_rest_delivery_point",
		MarkdownDescription: "A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.\n\n\n\nThe minimum access scope/level required to perform this operation is \"vpn/read-only\".\n\nThis has been available since SEMP API version 2.0.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "clientProfileName",
				TerraformName:       "client_profile_name",
				MarkdownDescription: "The Client Profile of the REST Delivery Point. It must exist in the local Message VPN. Its TCP parameters are used for all REST Consumers in this RDP. Its queue properties are used by the RDP client. The Client Profile is used inside the auto-generated Client Username for this RDP.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"default\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^#?[A-Za-z0-9\\-_]+$"), ""),
				},
				Default: "default",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable the REST Delivery Point. When disabled, no connections are initiated or messages delivered to any of the contained REST Consumers.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "msgVpnName",
				TerraformName:       "msg_vpn_name",
				MarkdownDescription: "The name of the Message VPN.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?]+$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "restDeliveryPointName",
				TerraformName:       "rest_delivery_point_name",
				MarkdownDescription: "The name of the REST Delivery Point.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "service",
				TerraformName:       "service",
				MarkdownDescription: "The name of the service that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.19.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 50),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "vendor",
				TerraformName:       "vendor",
				MarkdownDescription: "The name of the vendor that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.19.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 50),
				},
				Default: "",
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
