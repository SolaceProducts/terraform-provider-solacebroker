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
		TerraformName:       "msg_vpn_kafka_receiver_topic_binding",
		MarkdownDescription: "A Topic Binding receives messages from a remote Kafka Topic.\n\n\nAttribute|Identifying\n:---|:---:\nkafka_receiver_name|x\nmsg_vpn_name|x\ntopic_name|x\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.36.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName}/topicBindings/{topicName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable this topic binding of the Kafka Receiver. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "initialOffset",
				TerraformName:       "initial_offset",
				MarkdownDescription: "The initial offset to consume from the Kafka Topic if no member of the group has consumed and committed any offset already, or if the last committed offset has been deleted. Offsets are unique per partition.\n\nThis corresponds to the Kafka consumer API `auto.offset.reset` configuration setting.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"end\"`. The allowed values and their meaning are:\n\n<pre>\n\"beginning\" - Start with the earliest offset available.\n\"end\" - Start with new offsets only.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("beginning", "end"),
				},
				Default: "end",
			},
			{
				BaseType:            broker.String,
				SempName:            "kafkaReceiverName",
				TerraformName:       "kafka_receiver_name",
				MarkdownDescription: "The name of the Kafka Receiver.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^#*? ]([^*?]*[^*? ])?$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "localKey",
				TerraformName:       "local_key",
				MarkdownDescription: "The Substitution Expression used to generate the key for each message received from Kafka. This expression can include fields extracted from the metadata of each individual Kafka message as it is received from the Kafka Topic.\n\nIf empty, no key is included for each message as it is published into Solace.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 1024),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "localTopic",
				TerraformName:       "local_topic",
				MarkdownDescription: "The Substitution Expression used to generate the Solace Topic for each message received from Kafka. This expression can include data extracted from the metadata of each individual Kafka message as it is received from the Kafka Topic.\n\nIf empty, the Topic Binding will not be operational.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 1024),
				},
				Default: "",
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
				SempName:            "topicName",
				TerraformName:       "topic_name",
				MarkdownDescription: "The name of the Topic.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^\\^.*|[a-zA-Z0-9\\._\\-]+$"), ""),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
