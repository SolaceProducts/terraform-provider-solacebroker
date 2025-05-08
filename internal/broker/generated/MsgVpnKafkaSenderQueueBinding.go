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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_kafka_sender_queue_binding",
		MarkdownDescription: "A Queue Binding sends messages from a local Solace Queue to a remote Kafka topic.\n\n\n\nThe minimum access scope/level required to perform this operation is \"vpn/read-only\".\n\nThis has been available since SEMP API version 2.36.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}/queueBindings/{queueName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "ackMode",
				TerraformName:       "ack_mode",
				MarkdownDescription: "The number of acks required from the remote Kafka broker. When \"none\" messages are delivered at-most-once. When \"one\" or \"all\" messages are delivered at-least-once but may be reordered. This must be configured as \"all\" for an idempotent Kafka Sender, otherwise the Queue Binding will be operationally down.\n\nThis corresponds to the Kafka producer API `acks` configuration setting.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"all\"`. The allowed values and their meaning are:\n\n<pre>\n\"none\" - No Acks.\n\"one\" - Leader Ack Only.\n\"all\" - All Replica Acks.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("none", "one", "all"),
				},
				Default: "all",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable this queue binding of the Kafka Sender.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "kafkaSenderName",
				TerraformName:       "kafka_sender_name",
				MarkdownDescription: "The name of the Kafka Sender.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
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
				SempName:            "partitionConsistentHash",
				TerraformName:       "partition_consistent_hash",
				MarkdownDescription: "The hash algorithm to use for consistent partition selection.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"crc\"`. The allowed values and their meaning are:\n\n<pre>\n\"crc\" - CRC Hash.\n\"murmur2\" - Murmer2 Hash.\n\"fnv1a\" - Fowler-Noll-Vo 1a Hash.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("crc", "murmur2", "fnv1a"),
				},
				Default: "crc",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "partitionExplicitNumber",
				TerraformName:       "partition_explicit_number",
				MarkdownDescription: "The partition number to use for explicit partition selection.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				Default: 0,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "partitionRandomFallbackEnabled",
				TerraformName:       "partition_random_fallback_enabled",
				MarkdownDescription: "Enable or disable fallback to the random partition selection scheme when the consistent partition scheme is being used but no partition key is available for the message. When enabled a random partition will be selected for each unkeyed messages, otherwise some partition will be selected for groups of unkeyed messages.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.String,
				SempName:            "partitionScheme",
				TerraformName:       "partition_scheme",
				MarkdownDescription: "The partitioning scheme used to select a partition of the topic on the Kafka cluster to send messages to.\n\nThis corresponds to the Kafka producer API `partitioner.class` configuration setting.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"consistent\"`. The allowed values and their meaning are:\n\n<pre>\n\"consistent\" - Select a consistent partition for each key value. A hash of the key will be used to select the partition number.\n\"explicit\" - Select an explicit partition independent of key value.\n\"random\" - Select a random partition independent of key value.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("consistent", "explicit", "random"),
				},
				Default: "consistent",
			},
			{
				BaseType:            broker.String,
				SempName:            "queueName",
				TerraformName:       "queue_name",
				MarkdownDescription: "The name of the Queue.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&;]+$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "remoteKey",
				TerraformName:       "remote_key",
				MarkdownDescription: "The Substitution Expression used to generate the key for each message sent to Kafka. This expression can include fields extracted from the metadata of each individual Solace message as it is taken from the Solace Queue.\n\nIf empty, no key is included for each message as it is published into Kafka.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
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
				SempName:            "remoteTopic",
				TerraformName:       "remote_topic",
				MarkdownDescription: "The Kafka Topic on the Kafka Cluster to send each message taken from the Solace Queue to.\n\nIf empty, the Queue Binding will not be operational.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 249),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\._\\-]*$"), ""),
				},
				Default: "",
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
