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
		TerraformName:       "msg_vpn_topic_endpoint_template",
		MarkdownDescription: "A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.14.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "accessType",
				TerraformName:       "access_type",
				MarkdownDescription: "The access type for delivering messages to consumer flows. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"exclusive\"`. The allowed values and their meaning are:\n\n<pre>\n\"exclusive\" - Exclusive delivery of messages to the first bound consumer flow.\n\"non-exclusive\" - Non-exclusive delivery of messages to bound consumer flows in a round-robin fashion.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("exclusive", "non-exclusive"),
				},
				Default: "exclusive",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "consumerAckPropagationEnabled",
				TerraformName:       "consumer_ack_propagation_enabled",
				MarkdownDescription: "Enable or disable the propagation of consumer acknowledgments (ACKs) received on the active replication Message VPN to the standby replication Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.String,
				SempName:            "deadMsgQueue",
				TerraformName:       "dead_msg_queue",
				MarkdownDescription: "The name of the Dead Message Queue (DMQ). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"#DEAD_MSG_QUEUE\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&;]+$"), ""),
				},
				Default: "#DEAD_MSG_QUEUE",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "deliveryDelay",
				TerraformName:       "delivery_delay",
				MarkdownDescription: "The delay, in seconds, to apply to messages arriving on the Topic Endpoint before the messages are eligible for delivery. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`. Available since SEMP API version 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				Default: 0,
			},
			{
				BaseType:            broker.Struct,
				SempName:            "eventBindCountThreshold",
				TerraformName:       "event_bind_count_threshold",
				MarkdownDescription: "The thresholds for the Topic Endpoint consumer flows event, relative to `max_bind_count`.",
				Attributes: []*broker.AttributeInfo{
					{
						BaseType:            broker.Int64,
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET. The default value is: `60`.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							int64validator.Between(0, 100),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 60,
					},
					{
						BaseType:            broker.Int64,
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							int64validator.Between(0, 10000),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						BaseType:            broker.Int64,
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET. The default value is: `80`.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							int64validator.Between(0, 100),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 80,
					},
					{
						BaseType:            broker.Int64,
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							int64validator.Between(0, 10000),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				BaseType:            broker.Struct,
				SempName:            "eventMsgSpoolUsageThreshold",
				TerraformName:       "event_msg_spool_usage_threshold",
				MarkdownDescription: "The thresholds for the message spool usage event of the Topic Endpoint, relative to `max_spool_usage`.",
				Attributes: []*broker.AttributeInfo{
					{
						BaseType:            broker.Int64,
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET. The default value is: `18`.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							int64validator.Between(0, 100),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 18,
					},
					{
						BaseType:            broker.Int64,
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							int64validator.Between(0, 6000000),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						BaseType:            broker.Int64,
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET. The default value is: `25`.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							int64validator.Between(0, 100),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 25,
					},
					{
						BaseType:            broker.Int64,
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							int64validator.Between(0, 6000000),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				BaseType:            broker.Struct,
				SempName:            "eventRejectLowPriorityMsgLimitThreshold",
				TerraformName:       "event_reject_low_priority_msg_limit_threshold",
				MarkdownDescription: "The thresholds for the maximum allowed number of any priority messages queued in the Topic Endpoint event, relative to `reject_low_priority_msg_limit`.",
				Attributes: []*broker.AttributeInfo{
					{
						BaseType:            broker.Int64,
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET. The default value is: `60`.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							int64validator.Between(0, 100),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 60,
					},
					{
						BaseType:            broker.Int64,
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							int64validator.Between(0, 4294967295),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						BaseType:            broker.Int64,
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET. The default value is: `80`.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							int64validator.Between(0, 100),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 80,
					},
					{
						BaseType:            broker.Int64,
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							int64validator.Between(0, 4294967295),
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxBindCount",
				TerraformName:       "max_bind_count",
				MarkdownDescription: "The maximum number of consumer flows that can bind. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 10000),
				},
				Default: 1,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxDeliveredUnackedMsgsPerFlow",
				TerraformName:       "max_delivered_unacked_msgs_per_flow",
				MarkdownDescription: "The maximum number of messages delivered but not acknowledged per flow. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 1000000),
				},
				Default: 10000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxMsgSize",
				TerraformName:       "max_msg_size",
				MarkdownDescription: "The maximum message size allowed, in bytes (B). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 30000000),
				},
				Default: 1e+07,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxMsgSpoolUsage",
				TerraformName:       "max_msg_spool_usage",
				MarkdownDescription: "The maximum message spool usage allowed, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `5000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 6000000),
				},
				Default: 5000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxRedeliveryCount",
				TerraformName:       "max_redelivery_count",
				MarkdownDescription: "The maximum number of message redelivery attempts that will occur prior to the message being discarded or moved to the DMQ. A value of 0 means to retry forever. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
				Default: 0,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxTtl",
				TerraformName:       "max_ttl",
				MarkdownDescription: "The maximum time in seconds a message can stay in the Topic Endpoint when `respect_ttl_enabled` is `\"true\"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `max_ttl` configured for the Topic Endpoint, is exceeded. A value of 0 disables expiry. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				Default: 0,
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
				SempName:            "permission",
				TerraformName:       "permission",
				MarkdownDescription: "The permission level for all consumers, excluding the owner. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"no-access\"`. The allowed values and their meaning are:\n\n<pre>\n\"no-access\" - Disallows all access.\n\"read-only\" - Read-only access to the messages.\n\"consume\" - Consume (read and remove) messages.\n\"modify-topic\" - Consume messages or modify the topic/selector.\n\"delete\" - Consume messages, modify the topic/selector or delete the Client created endpoint altogether.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("no-access", "read-only", "consume", "modify-topic", "delete"),
				},
				Default: "no-access",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "redeliveryDelayEnabled",
				TerraformName:       "redelivery_delay_enabled",
				MarkdownDescription: "Enable or disable a message redelivery delay. When false, messages are redelivered as-soon-as-possible.  When true, messages are redelivered according to the initial, max and multiplier.  This should only be enabled when redelivery is enabled. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since SEMP API version 2.33.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "redeliveryDelayInitialInterval",
				TerraformName:       "redelivery_delay_initial_interval",
				MarkdownDescription: "The delay to be used between the first 2 redelivery attempts.  This value is in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`. Available since SEMP API version 2.33.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 3600000),
				},
				Default: 1000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "redeliveryDelayMaxInterval",
				TerraformName:       "redelivery_delay_max_interval",
				MarkdownDescription: "The maximum delay to be used between any 2 redelivery attempts.  This value is in milliseconds.  Due to technical limitations, some redelivery attempt delays may slightly exceed this value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `64000`. Available since SEMP API version 2.33.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 10800000),
				},
				Default: 64000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "redeliveryDelayMultiplier",
				TerraformName:       "redelivery_delay_multiplier",
				MarkdownDescription: "The amount each delay interval is multiplied by after each failed delivery attempt.  This number is in a fixed-point decimal format in which you must divide by 100 to get the floating point value. For example, a value of 125 would cause the delay to be multiplied by 1.25. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `200`. Available since SEMP API version 2.33.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(100, 500),
				},
				Default: 200,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "redeliveryEnabled",
				TerraformName:       "redelivery_enabled",
				MarkdownDescription: "Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by max_redelivery_count. When disabled, the message will never be delivered from the topic-endpoint more than once. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since SEMP API version 2.18.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "rejectLowPriorityMsgEnabled",
				TerraformName:       "reject_low_priority_msg_enabled",
				MarkdownDescription: "Enable or disable the checking of low priority messages against the `reject_low_priority_msg_limit`. This may only be enabled if `reject_msg_to_sender_on_discard_behavior` does not have a value of `\"never\"`. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "rejectLowPriorityMsgLimit",
				TerraformName:       "reject_low_priority_msg_limit",
				MarkdownDescription: "The number of messages that are permitted before low priority messages are rejected. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				Default: 0,
			},
			{
				BaseType:            broker.String,
				SempName:            "rejectMsgToSenderOnDiscardBehavior",
				TerraformName:       "reject_msg_to_sender_on_discard_behavior",
				MarkdownDescription: "Determines when to return negative acknowledgments (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"never\"`. The allowed values and their meaning are:\n\n<pre>\n\"never\" - Silently discard messages.\n\"when-topic-endpoint-enabled\" - NACK each message discard back to the client, except messages that are discarded because an endpoint is administratively disabled.\n\"always\" - NACK each message discard back to the client, including messages that are discarded because an endpoint is administratively disabled.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("never", "when-topic-endpoint-enabled", "always"),
				},
				Default: "never",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "respectMsgPriorityEnabled",
				TerraformName:       "respect_msg_priority_enabled",
				MarkdownDescription: "Enable or disable the respecting of message priority. When enabled, messages are delivered in priority order, from 9 (highest) to 0 (lowest). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "respectTtlEnabled",
				TerraformName:       "respect_ttl_enabled",
				MarkdownDescription: "Enable or disable the respecting of the time-to-live (TTL) for messages. When enabled, expired messages are discarded or moved to the DMQ. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "topicEndpointNameFilter",
				TerraformName:       "topic_endpoint_name_filter",
				MarkdownDescription: "A pattern used to determine which Topic Endpoints use settings from this Template. Two different wildcards can be used in the pattern: * and &gt;. Similar to topic filters or subscription patterns, a &gt; matches anything (but only when used at the end), and a * matches zero or more characters but never a slash (/). A &gt; is only a wildcard when  used at the end, after a /. A * is only allowed at the end, after a slash (/). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "topicEndpointTemplateName",
				TerraformName:       "topic_endpoint_template_name",
				MarkdownDescription: "The name of the Topic Endpoint Template.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^#?*][^?*]*$"), ""),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
