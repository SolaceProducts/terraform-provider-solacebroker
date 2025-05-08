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
		TerraformName:       "msg_vpn_telemetry_profile",
		MarkdownDescription: "Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.\n\n\n\nThe minimum access scope/level required to perform this operation is \"vpn/read-only\".\n\nThis has been available since SEMP API version 2.31.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
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
				BaseType:            broker.Struct,
				SempName:            "queueEventBindCountThreshold",
				TerraformName:       "queue_event_bind_count_threshold",
				MarkdownDescription: "The thresholds for the Queue consumer flows event, relative to `queue_max_bind_count`.",
				Attributes: []*broker.AttributeInfo{
					{
						BaseType:            broker.Int64,
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `60`.",
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
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
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
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `80`.",
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
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
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
				MarkdownDescription: "The thresholds for the message spool usage event of the Queue, relative to `queue_max_msg_spool_usage`.",
				Attributes: []*broker.AttributeInfo{
					{
						BaseType:            broker.Int64,
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
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
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
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
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `2`.",
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
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
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
				MarkdownDescription: "The maximum number of consumer flows that can bind to the Queue.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
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
				SempName:            "queueMaxMsgSpoolUsage",
				TerraformName:       "queue_max_msg_spool_usage",
				MarkdownDescription: "The maximum message spool usage allowed by the Queue, in megabytes (MB).\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `800000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 6000000),
				},
				Default: 800000,
			},
			{
				BaseType:            broker.String,
				SempName:            "receiverAclConnectDefaultAction",
				TerraformName:       "receiver_acl_connect_default_action",
				MarkdownDescription: "The default action to take when a receiver client connects to the broker.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"disallow\"`. The allowed values and their meaning are:\n\n<pre>\n\"allow\" - Allow client connection unless an exception is found for it.\n\"disallow\" - Disallow client connection unless an exception is found for it.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("allow", "disallow"),
				},
				Default: "disallow",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "receiverEnabled",
				TerraformName:       "receiver_enabled",
				MarkdownDescription: "Enable or disable the ability for receiver clients to consume from the #telemetry queue.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Struct,
				SempName:            "receiverEventConnectionCountPerClientUsernameThreshold",
				TerraformName:       "receiver_event_connection_count_per_client_username_threshold",
				MarkdownDescription: "The thresholds for the receiver connection count event, relative to `receiver_max_connection_count_per_client_username`.",
				Attributes: []*broker.AttributeInfo{
					{
						BaseType:            broker.Int64,
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `60`.",
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
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							int64validator.Between(0, 200000),
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
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `80`.",
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
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". This attribute may not be returned in a GET. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Int64Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							int64validator.Between(0, 200000),
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
				SempName:            "receiverMaxConnectionCountPerClientUsername",
				TerraformName:       "receiver_max_connection_count_per_client_username",
				MarkdownDescription: "The maximum number of receiver connections per Client Username.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is the maximum value supported by the platform.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 200000),
				},
			},
			{
				BaseType:            broker.Int64,
				SempName:            "receiverTcpCongestionWindowSize",
				TerraformName:       "receiver_tcp_congestion_window_size",
				MarkdownDescription: "The TCP initial congestion window size for clients using the Client Profile, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact support before changing this value.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `2`.",
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
				SempName:            "receiverTcpKeepaliveCount",
				TerraformName:       "receiver_tcp_keepalive_count",
				MarkdownDescription: "The number of TCP keepalive retransmissions to a client using the Client Profile before declaring that it is not available.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `5`.",
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
				SempName:            "receiverTcpKeepaliveIdleTime",
				TerraformName:       "receiver_tcp_keepalive_idle_time",
				MarkdownDescription: "The amount of time a client connection using the Client Profile must remain idle before TCP begins sending keepalive probes, in seconds.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3`.",
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
				SempName:            "receiverTcpKeepaliveInterval",
				TerraformName:       "receiver_tcp_keepalive_interval",
				MarkdownDescription: "The amount of time between TCP keepalive retransmissions to a client using the Client Profile when no acknowledgment is received, in seconds.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
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
				SempName:            "receiverTcpMaxSegmentSize",
				TerraformName:       "receiver_tcp_max_segment_size",
				MarkdownDescription: "The TCP maximum segment size for clients using the Client Profile, in bytes. Changes are applied to all existing connections.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1460`.",
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
				SempName:            "receiverTcpMaxWindowSize",
				TerraformName:       "receiver_tcp_max_window_size",
				MarkdownDescription: "The TCP maximum window size for clients using the Client Profile, in kilobytes. Changes are applied to all existing connections. This setting is ignored on the software broker.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `256`.",
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
				SempName:            "telemetryProfileName",
				TerraformName:       "telemetry_profile_name",
				MarkdownDescription: "The name of the Telemetry Profile.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 21),
					stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9\\-_]+$"), ""),
				},
			},
			{
				BaseType:            broker.Bool,
				SempName:            "traceEnabled",
				TerraformName:       "trace_enabled",
				MarkdownDescription: "Enable or disable generation of all trace span data messages. When enabled, the state of configured trace filters control which messages get traced. When disabled, trace span data messages are never generated, regardless of the state of trace filters.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "traceSendSpanGenerationEnabled",
				TerraformName:       "trace_send_span_generation_enabled",
				MarkdownDescription: "Enable or disable generation of send spans. For the most complete view of broker message processing, this should be enabled. If the information provided by send spans are not needed, send spans can be disabled to reduce the performance impact of tracing.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since SEMP API version 2.36.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
