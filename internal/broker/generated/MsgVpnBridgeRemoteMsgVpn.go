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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_bridge_remote_msg_vpn",
		MarkdownDescription: "The Remote Message VPN is the Message VPN that the Bridge connects to.\n\n\n\nThe minimum access scope/level required to perform this operation is \"vpn/read-only\".\n\nThis has been available since SEMP API version 2.0.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "bridgeName",
				TerraformName:       "bridge_name",
				MarkdownDescription: "The name of the Bridge.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 150),
					stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9\"~`!\\\\@$%|\\^()_+={}:,.#\\-;\\[\\]]+$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "bridgeVirtualRouter",
				TerraformName:       "bridge_virtual_router",
				MarkdownDescription: "The virtual router of the Bridge.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The allowed values and their meaning are:\n\n<pre>\n\"primary\" - The Bridge is used for the primary virtual router.\n\"backup\" - The Bridge is used for the backup virtual router.\n\"auto\" - The Bridge is automatically assigned a virtual router at creation, depending on the broker's active-standby role.\n</pre>\n",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("primary", "backup", "auto"),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "clientUsername",
				TerraformName:       "client_username",
				MarkdownDescription: "The Client Username the Bridge uses to login to the remote Message VPN. This per remote Message VPN value overrides the value provided for the Bridge overall.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Requires:            []string{"password"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("password"),
					),
					stringvalidator.LengthBetween(0, 189),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^?*]*$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "compressedDataEnabled",
				TerraformName:       "compressed_data_enabled",
				MarkdownDescription: "Enable or disable data compression for the remote Message VPN connection.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "connectOrder",
				TerraformName:       "connect_order",
				MarkdownDescription: "The preference given to incoming connections from remote Message VPN hosts, from 1 (highest priority) to 4 (lowest priority).\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 4),
				},
				Default: 4,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "egressFlowWindowSize",
				TerraformName:       "egress_flow_window_size",
				MarkdownDescription: "The number of outstanding guaranteed messages that can be transmitted over the remote Message VPN connection before an acknowledgment is received.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `255`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
				Default: 255,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable the remote Message VPN.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
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
				SempName:            "password",
				TerraformName:       "password",
				MarkdownDescription: "The password for the Client Username.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Sensitive:           true,
				Requires:            []string{"client_username"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("client_username"),
					),
					stringvalidator.LengthBetween(0, 128),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "queueBinding",
				TerraformName:       "queue_binding",
				MarkdownDescription: "The queue binding of the Bridge in the remote Message VPN.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&;]*$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "remoteMsgVpnInterface",
				TerraformName:       "remote_msg_vpn_interface",
				MarkdownDescription: "The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remote_msg_vpn_location` must not be a virtual router name.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 15),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "remoteMsgVpnLocation",
				TerraformName:       "remote_msg_vpn_location",
				MarkdownDescription: "The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 259),
					stringvalidator.RegexMatches(regexp.MustCompile("^((((([0-9a-zA-Z_\\-\\.])+)|\\[([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,7}:\\]|\\[([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}\\]|\\[([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}\\]|\\[([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}\\]|\\[[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})\\]|\\[:((:[0-9a-fA-F]{1,4}){1,7}|:)\\])((:[0-9]{1,5}){0,1}))|(v:.+))$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "remoteMsgVpnName",
				TerraformName:       "remote_msg_vpn_name",
				MarkdownDescription: "The name of the remote Message VPN.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
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
				BaseType:            broker.Bool,
				SempName:            "tlsEnabled",
				TerraformName:       "tls_enabled",
				MarkdownDescription: "Enable or disable encryption (TLS) for the remote Message VPN connection.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "unidirectionalClientProfile",
				TerraformName:       "unidirectional_client_profile",
				MarkdownDescription: "The Client Profile for the unidirectional Bridge of the remote Message VPN. The Client Profile must exist in the local Message VPN, and it is used only for the TCP parameters. Note that the default client profile has a TCP maximum window size of 2 MB.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"#client-profile\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^#?[A-Za-z0-9\\-_]+$"), ""),
				},
				Default: "#client-profile",
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
