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
		TerraformName:       "dmr_cluster_link",
		MarkdownDescription: "A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.\n\n\nAttribute|Identifying|Write-Only|Opaque\n:---|:---:|:---:|:---:\nauthentication_basic_password||x|x\ndmr_cluster_name|x||\nremote_node_name|x||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.11.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "authenticationBasicPassword",
				TerraformName:       "authentication_basic_password",
				MarkdownDescription: "The password used to authenticate with the remote node when using basic internal authentication. If this per-Link password is not configured, the Cluster's password is used instead. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationScheme",
				TerraformName:       "authentication_scheme",
				MarkdownDescription: "The authentication scheme to be used by the Link which initiates connections to the remote node. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"basic\"`. The allowed values and their meaning are:\n\n<pre>\n\"basic\" - Basic Authentication Scheme (via username and password).\n\"client-certificate\" - Client Certificate Authentication Scheme (via certificate file or content).\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("basic", "client-certificate"),
				},
				Default: "basic",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueControl1MaxDepth",
				TerraformName:       "client_profile_queue_control1_max_depth",
				MarkdownDescription: "The maximum depth of the \"Control 1\" (C-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueControl1MinMsgBurst",
				TerraformName:       "client_profile_queue_control1_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Control 1\" (C-1) priority queue, regardless of the `client_profile_queue_control1_max_depth` value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueDirect1MaxDepth",
				TerraformName:       "client_profile_queue_direct1_max_depth",
				MarkdownDescription: "The maximum depth of the \"Direct 1\" (D-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueDirect1MinMsgBurst",
				TerraformName:       "client_profile_queue_direct1_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Direct 1\" (D-1) priority queue, regardless of the `client_profile_queue_direct1_max_depth` value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueDirect2MaxDepth",
				TerraformName:       "client_profile_queue_direct2_max_depth",
				MarkdownDescription: "The maximum depth of the \"Direct 2\" (D-2) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueDirect2MinMsgBurst",
				TerraformName:       "client_profile_queue_direct2_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Direct 2\" (D-2) priority queue, regardless of the `client_profile_queue_direct2_max_depth` value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueDirect3MaxDepth",
				TerraformName:       "client_profile_queue_direct3_max_depth",
				MarkdownDescription: "The maximum depth of the \"Direct 3\" (D-3) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueDirect3MinMsgBurst",
				TerraformName:       "client_profile_queue_direct3_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Direct 3\" (D-3) priority queue, regardless of the `client_profile_queue_direct3_max_depth` value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueGuaranteed1MaxDepth",
				TerraformName:       "client_profile_queue_guaranteed1_max_depth",
				MarkdownDescription: "The maximum depth of the \"Guaranteed 1\" (G-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileQueueGuaranteed1MinMsgBurst",
				TerraformName:       "client_profile_queue_guaranteed1_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Guaranteed 1\" (G-1) priority queue, regardless of the `client_profile_queue_guaranteed1_max_depth` value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `255`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 262144),
				},
				Default: 255,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileTcpCongestionWindowSize",
				TerraformName:       "client_profile_tcp_congestion_window_size",
				MarkdownDescription: "The TCP initial congestion window size, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact support before changing this value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `2`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(2, 7826),
				},
				Default: 2,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileTcpKeepaliveCount",
				TerraformName:       "client_profile_tcp_keepalive_count",
				MarkdownDescription: "The number of TCP keepalive retransmissions to be carried out before declaring that the remote end is not available. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `5`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(2, 5),
				},
				Default: 5,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileTcpKeepaliveIdleTime",
				TerraformName:       "client_profile_tcp_keepalive_idle_time",
				MarkdownDescription: "The amount of time a connection must remain idle before TCP begins sending keepalive probes, in seconds. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `3`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(3, 120),
				},
				Default: 3,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileTcpKeepaliveInterval",
				TerraformName:       "client_profile_tcp_keepalive_interval",
				MarkdownDescription: "The amount of time between TCP keepalive retransmissions when no acknowledgment is received, in seconds. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
				Default: 1,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileTcpMaxSegmentSize",
				TerraformName:       "client_profile_tcp_max_segment_size",
				MarkdownDescription: "The TCP maximum segment size, in bytes. Changes are applied to all existing connections. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `1460`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(256, 1460),
				},
				Default: 1460,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "clientProfileTcpMaxWindowSize",
				TerraformName:       "client_profile_tcp_max_window_size",
				MarkdownDescription: "The TCP maximum window size, in kilobytes. Changes are applied to all existing connections. This setting is ignored on the software broker. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `256`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(16, 65536),
				},
				Default: 256,
			},
			{
				BaseType:            broker.String,
				SempName:            "dmrClusterName",
				TerraformName:       "dmr_cluster_name",
				MarkdownDescription: "The name of the Cluster.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			{
				BaseType:            broker.Int64,
				SempName:            "egressFlowWindowSize",
				TerraformName:       "egress_flow_window_size",
				MarkdownDescription: "The number of outstanding guaranteed messages that can be sent over the Link before acknowledgment is received by the sender. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `255`.",
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
				MarkdownDescription: "Enable or disable the Link. When disabled, subscription sets of this and the remote node are not kept up-to-date, and messages are not exchanged with the remote node. Published guaranteed messages will be queued up for future delivery based on current subscription sets. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "initiator",
				TerraformName:       "initiator",
				MarkdownDescription: "The initiator of the Link's TCP connections. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"lexical\"`. The allowed values and their meaning are:\n\n<pre>\n\"lexical\" - The \"higher\" node-name initiates.\n\"local\" - The local node initiates.\n\"remote\" - The remote node initiates.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("lexical", "local", "remote"),
				},
				Default: "lexical",
			},
			{
				BaseType:            broker.String,
				SempName:            "queueDeadMsgQueue",
				TerraformName:       "queue_dead_msg_queue",
				MarkdownDescription: "The name of the Dead Message Queue (DMQ) used by the Queue for discarded messages. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"#DEAD_MSG_QUEUE\"`.",
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
				SempName:            "queueEventSpoolUsageThreshold",
				TerraformName:       "queue_event_spool_usage_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						BaseType:            broker.Int64,
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
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
						Default: 1,
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
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
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
						Default: 2,
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
				SempName:            "queueMaxDeliveredUnackedMsgsPerFlow",
				TerraformName:       "queue_max_delivered_unacked_msgs_per_flow",
				MarkdownDescription: "The maximum number of messages delivered but not acknowledged per flow for the Queue. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `1000000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 1000000),
				},
				Default: 1e+06,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "queueMaxMsgSpoolUsage",
				TerraformName:       "queue_max_msg_spool_usage",
				MarkdownDescription: "The maximum message spool usage by the Queue (quota), in megabytes (MB). Changes to this attribute are synchronized to HA mates via config-sync. The default value is `800000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 6000000),
				},
				Default: 800000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "queueMaxRedeliveryCount",
				TerraformName:       "queue_max_redelivery_count",
				MarkdownDescription: "The maximum number of times the Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `0`.",
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
				MarkdownDescription: "The maximum time in seconds a message can stay in the Queue when `queue_respect_ttl_enabled` is `true`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `queue_max_ttl` configured for the Queue, is exceeded. A value of 0 disables expiry. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `0`.",
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
				MarkdownDescription: "Determines when to return negative acknowledgments (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"always\"`. The allowed values and their meaning are:\n\n<pre>\n\"never\" - Silently discard messages.\n\"when-queue-enabled\" - NACK each message discard back to the client, except messages that are discarded because an endpoint is administratively disabled.\n\"always\" - NACK each message discard back to the client, including messages that are discarded because an endpoint is administratively disabled.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("never", "when-queue-enabled", "always"),
				},
				Default: "always",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "queueRespectTtlEnabled",
				TerraformName:       "queue_respect_ttl_enabled",
				MarkdownDescription: "Enable or disable the respecting of the time-to-live (TTL) for messages in the Queue. When enabled, expired messages are discarded or moved to the DMQ. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "remoteNodeName",
				TerraformName:       "remote_node_name",
				MarkdownDescription: "The name of the node at the remote end of the Link.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&/]+$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "span",
				TerraformName:       "span",
				MarkdownDescription: "The span of the Link, either internal or external. Internal Links connect nodes within the same Cluster. External Links connect nodes within different Clusters. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"external\"`. The allowed values and their meaning are:\n\n<pre>\n\"internal\" - Link to same cluster.\n\"external\" - Link to other cluster.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("internal", "external"),
				},
				Default: "external",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "transportCompressedEnabled",
				TerraformName:       "transport_compressed_enabled",
				MarkdownDescription: "Enable or disable compression on the Link. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "transportTlsEnabled",
				TerraformName:       "transport_tls_enabled",
				MarkdownDescription: "Enable or disable encryption (TLS) on the Link. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
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
