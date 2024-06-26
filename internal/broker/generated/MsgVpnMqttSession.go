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
		TerraformName:       "msg_vpn_mqtt_session",
		MarkdownDescription: "An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.1.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable the MQTT Session. When disabled, the client is disconnected, new messages matching QoS 0 subscriptions are discarded, and new messages matching QoS 1 subscriptions are stored for future delivery. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "mqttSessionClientId",
				TerraformName:       "mqtt_session_client_id",
				MarkdownDescription: "The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.",
				Identifying:         true,
				Required:            true,
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
				MarkdownDescription: "The virtual router of the MQTT Session. The allowed values and their meaning are:\n\n<pre>\n\"primary\" - The MQTT Session belongs to the primary virtual router.\n\"backup\" - The MQTT Session belongs to the backup virtual router.\n\"auto\" - The MQTT Session is automatically assigned a virtual router at creation, depending on the broker's active-standby role.\n</pre>\n",
				Identifying:         true,
				Required:            true,
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
				SempName:            "owner",
				TerraformName:       "owner",
				MarkdownDescription: "The owner of the MQTT Session. For externally-created sessions this defaults to the Client Username of the connecting client. For management-created sessions this defaults to empty. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 189),
				},
				Default: "",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "queueConsumerAckPropagationEnabled",
				TerraformName:       "queue_consumer_ack_propagation_enabled",
				MarkdownDescription: "Enable or disable the propagation of consumer acknowledgments (ACKs) received on the active replication Message VPN to the standby replication Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since SEMP API version 2.14.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.String,
				SempName:            "queueDeadMsgQueue",
				TerraformName:       "queue_dead_msg_queue",
				MarkdownDescription: "The name of the Dead Message Queue (DMQ) used by the MQTT Session Queue. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"#DEAD_MSG_QUEUE\"`. Available since SEMP API version 2.14.",
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
				BaseType:            broker.Struct,
				SempName:            "queueEventBindCountThreshold",
				TerraformName:       "queue_event_bind_count_threshold",
				MarkdownDescription: "Thresholds for the high number of the MQTT Session Queue Consumers Event, relative to `queue_max_bind_count`. Available since SEMP API version 2.14.",
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
				BaseType:            broker.Struct,
				SempName:            "queueEventMsgSpoolUsageThreshold",
				TerraformName:       "queue_event_msg_spool_usage_threshold",
				MarkdownDescription: "The threshold for the Message Spool usage event of the MQTT Session Queue, relative to `queue_max_msg_spool_usage`. Available since SEMP API version 2.14.",
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
				BaseType:            broker.Struct,
				SempName:            "queueEventRejectLowPriorityMsgLimitThreshold",
				TerraformName:       "queue_event_reject_low_priority_msg_limit_threshold",
				MarkdownDescription: "The threshold for the maximum allowed number of any priority messages queued in the MQTT Session Queue, relative to `queue_reject_low_priority_msg_limit`. Available since SEMP API version 2.14.",
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
				SempName:            "queueMaxBindCount",
				TerraformName:       "queue_max_bind_count",
				MarkdownDescription: "The maximum number of consumer flows that can bind to the MQTT Session Queue. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`. Available since SEMP API version 2.14.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 10000),
				},
				Default: 1000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "queueMaxDeliveredUnackedMsgsPerFlow",
				TerraformName:       "queue_max_delivered_unacked_msgs_per_flow",
				MarkdownDescription: "The maximum number of messages delivered but not acknowledged per flow for the MQTT Session Queue. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000`. Available since SEMP API version 2.14.",
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
				SempName:            "queueMaxMsgSize",
				TerraformName:       "queue_max_msg_size",
				MarkdownDescription: "The maximum message size allowed in the MQTT Session Queue, in bytes (B). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000000`. Available since SEMP API version 2.14.",
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
				SempName:            "queueMaxMsgSpoolUsage",
				TerraformName:       "queue_max_msg_spool_usage",
				MarkdownDescription: "The maximum message spool usage allowed by the MQTT Session Queue, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `5000`. Available since SEMP API version 2.14.",
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
				SempName:            "queueMaxRedeliveryCount",
				TerraformName:       "queue_max_redelivery_count",
				MarkdownDescription: "The maximum number of times the MQTT Session Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`. Available since SEMP API version 2.14.",
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
				SempName:            "queueMaxTtl",
				TerraformName:       "queue_max_ttl",
				MarkdownDescription: "The maximum time in seconds a message can stay in the MQTT Session Queue when `queue_respect_ttl_enabled` is `\"true\"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `queue_max_ttl` configured for the MQTT Session Queue, is exceeded. A value of 0 disables expiry. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`. Available since SEMP API version 2.14.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default: 0,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "queueRejectLowPriorityMsgEnabled",
				TerraformName:       "queue_reject_low_priority_msg_enabled",
				MarkdownDescription: "Enable or disable the checking of low priority messages against the `queue_reject_low_priority_msg_limit`. This may only be enabled if `queue_reject_msg_to_sender_on_discard_behavior` does not have a value of `\"never\"`. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since SEMP API version 2.14.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "queueRejectLowPriorityMsgLimit",
				TerraformName:       "queue_reject_low_priority_msg_limit",
				MarkdownDescription: "The number of messages of any priority in the MQTT Session Queue above which low priority messages are not admitted but higher priority messages are allowed. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`. Available since SEMP API version 2.14.",
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
				SempName:            "queueRejectMsgToSenderOnDiscardBehavior",
				TerraformName:       "queue_reject_msg_to_sender_on_discard_behavior",
				MarkdownDescription: "Determines when to return negative acknowledgments (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as queue_reject_low_priority_msg_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"when-queue-enabled\"`. The allowed values and their meaning are:\n\n<pre>\n\"never\" - Silently discard messages.\n\"when-queue-enabled\" - NACK each message discard back to the client, except messages that are discarded because an endpoint is administratively disabled.\n\"always\" - NACK each message discard back to the client, including messages that are discarded because an endpoint is administratively disabled.\n</pre>\n Available since SEMP API version 2.14.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("never", "when-queue-enabled", "always"),
				},
				Default: "when-queue-enabled",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "queueRespectTtlEnabled",
				TerraformName:       "queue_respect_ttl_enabled",
				MarkdownDescription: "Enable or disable the respecting of the time-to-live (TTL) for messages in the MQTT Session Queue. When enabled, expired messages are discarded or moved to the DMQ. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since SEMP API version 2.14.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
