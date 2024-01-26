// terraform-provider-solacebroker
//
// Copyright 2024 Solace Corporation. All rights reserved.
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
		TerraformName:       "msg_vpn_acl_profile",
		MarkdownDescription: "An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.\n\n\nAttribute|Identifying\n:---|:---:\nacl_profile_name|x\nmsg_vpn_name|x\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.0.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "aclProfileName",
				TerraformName:       "acl_profile_name",
				MarkdownDescription: "The name of the ACL Profile.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "clientConnectDefaultAction",
				TerraformName:       "client_connect_default_action",
				MarkdownDescription: "The default action to take when a client using the ACL Profile connects to the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"disallow\"`. The allowed values and their meaning are:\n\n<pre>\n\"allow\" - Allow client connection unless an exception is found for it.\n\"disallow\" - Disallow client connection unless an exception is found for it.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("allow", "disallow"),
				},
				Default: "disallow",
			},
			{
				BaseType:            broker.String,
				SempName:            "msgVpnName",
				TerraformName:       "msg_vpn_name",
				MarkdownDescription: "The name of the Message VPN.",
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
				SempName:            "publishTopicDefaultAction",
				TerraformName:       "publish_topic_default_action",
				MarkdownDescription: "The default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"disallow\"`. The allowed values and their meaning are:\n\n<pre>\n\"allow\" - Allow topic unless an exception is found for it.\n\"disallow\" - Disallow topic unless an exception is found for it.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("allow", "disallow"),
				},
				Default: "disallow",
			},
			{
				BaseType:            broker.String,
				SempName:            "subscribeShareNameDefaultAction",
				TerraformName:       "subscribe_share_name_default_action",
				MarkdownDescription: "The default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"allow\"`. The allowed values and their meaning are:\n\n<pre>\n\"allow\" - Allow topic unless an exception is found for it.\n\"disallow\" - Disallow topic unless an exception is found for it.\n</pre>\n Available since SEMP API version 2.14.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("allow", "disallow"),
				},
				Default: "allow",
			},
			{
				BaseType:            broker.String,
				SempName:            "subscribeTopicDefaultAction",
				TerraformName:       "subscribe_topic_default_action",
				MarkdownDescription: "The default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"disallow\"`. The allowed values and their meaning are:\n\n<pre>\n\"allow\" - Allow topic unless an exception is found for it.\n\"disallow\" - Disallow topic unless an exception is found for it.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("allow", "disallow"),
				},
				Default: "disallow",
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
