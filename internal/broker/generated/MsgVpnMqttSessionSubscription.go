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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_mqtt_session_subscription",
		MarkdownDescription: "An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.\n\n\n\nThe minimum access scope/level required to perform this operation is \"vpn/read-only\".\n\nThis has been available since SEMP API version 2.1.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "mqttSessionClientId",
				TerraformName:       "mqtt_session_client_id",
				MarkdownDescription: "The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "mqttSessionVirtualRouter",
				TerraformName:       "mqtt_session_virtual_router",
				MarkdownDescription: "The virtual router of the MQTT Session.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The allowed values and their meaning are:\n\n<pre>\n\"primary\" - The MQTT Session belongs to the primary virtual router.\n\"backup\" - The MQTT Session belongs to the backup virtual router.\n\"auto\" - The MQTT Session is automatically assigned a virtual router at creation, depending on the broker's active-standby role.\n</pre>\n",
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
				BaseType:            broker.Int64,
				SempName:            "subscriptionQos",
				TerraformName:       "subscription_qos",
				MarkdownDescription: "The quality of service (QoS) for the subscription as either 0 (deliver at most once) or 1 (deliver at least once). QoS 2 is not supported, but QoS 2 messages attracted by QoS 0 or QoS 1 subscriptions are accepted and delivered accordingly.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 1),
				},
				Default: 0,
			},
			{
				BaseType:            broker.String,
				SempName:            "subscriptionTopic",
				TerraformName:       "subscription_topic",
				MarkdownDescription: "The MQTT subscription topic.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 250),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
