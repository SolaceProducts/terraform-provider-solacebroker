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
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_client_profile",
		MarkdownDescription: "Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\nallowCutThroughForwardingEnabled|||x|\napiQueueManagementCopyFromOnCreateName|||x|\napiTopicEndpointManagementCopyFromOnCreateName|||x|\nclientProfileName|x|||\nmsgVpnName|x|||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since 2.0.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "allowBridgeConnectionsEnabled",
				TerraformName:       "allow_bridge_connections_enabled",
				MarkdownDescription: "Enable or disable allowing Bridge clients using the Client Profile to connect. Changing this setting does not affect existing Bridge client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "allowCutThroughForwardingEnabled",
				TerraformName:       "allow_cut_through_forwarding_enabled",
				MarkdownDescription: "Enable or disable allowing clients using the Client Profile to bind to endpoints with the cut-through forwarding delivery mode. Changing this value does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.22. This attribute has been deprecated. Please visit the Solace Product Lifecycle Policy web page for details on deprecated features.",
				Deprecated:          true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "allowGuaranteedEndpointCreateDurability",
				TerraformName:       "allow_guaranteed_endpoint_create_durability",
				MarkdownDescription: "The types of Queues and Topic Endpoints that clients using the client-profile can create. Changing this value does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"all\"`. The allowed values and their meaning are:\n\n<pre>\n\"all\" - Client can create any type of endpoint.\n\"durable\" - Client can create only durable endpoints.\n\"non-durable\" - Client can create only non-durable endpoints.\n</pre>\n Available since 2.14.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("all", "durable", "non-durable"),
				},
				Default: "all",
			},
			{
				SempName:            "allowGuaranteedEndpointCreateEnabled",
				TerraformName:       "allow_guaranteed_endpoint_create_enabled",
				MarkdownDescription: "Enable or disable allowing clients using the Client Profile to create topic endponts or queues. Changing this value does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "allowGuaranteedMsgReceiveEnabled",
				TerraformName:       "allow_guaranteed_msg_receive_enabled",
				MarkdownDescription: "Enable or disable allowing clients using the Client Profile to receive guaranteed messages. Changing this setting does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "allowGuaranteedMsgSendEnabled",
				TerraformName:       "allow_guaranteed_msg_send_enabled",
				MarkdownDescription: "Enable or disable allowing clients using the Client Profile to send guaranteed messages. Changing this setting does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "allowSharedSubscriptionsEnabled",
				TerraformName:       "allow_shared_subscriptions_enabled",
				MarkdownDescription: "Enable or disable allowing shared subscriptions. Changing this setting does not affect existing subscriptions. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.11.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "allowTransactedSessionsEnabled",
				TerraformName:       "allow_transacted_sessions_enabled",
				MarkdownDescription: "Enable or disable allowing clients using the Client Profile to establish transacted sessions. Changing this setting does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "apiQueueManagementCopyFromOnCreateName",
				TerraformName:       "api_queue_management_copy_from_on_create_name",
				MarkdownDescription: "The name of a queue to copy settings from when a new queue is created by a client using the Client Profile. The referenced queue must exist in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.14. This attribute has been replaced with `apiQueueManagementCopyFromOnCreateTemplateName`.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 200),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&;]*$"), ""),
				},
				Default: "",
			},
			{
				SempName:            "apiQueueManagementCopyFromOnCreateTemplateName",
				TerraformName:       "api_queue_management_copy_from_on_create_template_name",
				MarkdownDescription: "The name of a queue template to copy settings from when a new queue is created by a client using the Client Profile. If the referenced queue template does not exist, queue creation will fail when it tries to resolve this template. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since 2.14.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^#?*]?[^?*]*$"), ""),
				},
				Default: "",
			},
			{
				SempName:            "apiTopicEndpointManagementCopyFromOnCreateName",
				TerraformName:       "api_topic_endpoint_management_copy_from_on_create_name",
				MarkdownDescription: "The name of a topic endpoint to copy settings from when a new topic endpoint is created by a client using the Client Profile. The referenced topic endpoint must exist in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.14. This attribute has been replaced with `apiTopicEndpointManagementCopyFromOnCreateTemplateName`.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "",
			},
			{
				SempName:            "apiTopicEndpointManagementCopyFromOnCreateTemplateName",
				TerraformName:       "api_topic_endpoint_management_copy_from_on_create_template_name",
				MarkdownDescription: "The name of a topic endpoint template to copy settings from when a new topic endpoint is created by a client using the Client Profile. If the referenced topic endpoint template does not exist, topic endpoint creation will fail when it tries to resolve this template. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since 2.14.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^#?*]?[^?*]*$"), ""),
				},
				Default: "",
			},
			{
				SempName:            "clientProfileName",
				TerraformName:       "client_profile_name",
				MarkdownDescription: "The name of the Client Profile.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^#?[A-Za-z0-9\\-_]+$"), ""),
				},
			},
			{
				SempName:            "compressionEnabled",
				TerraformName:       "compression_enabled",
				MarkdownDescription: "Enable or disable allowing clients using the Client Profile to use compression. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since 2.10.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "elidingDelay",
				TerraformName:       "eliding_delay",
				MarkdownDescription: "The amount of time to delay the delivery of messages to clients using the Client Profile after the initial message has been delivered (the eliding delay interval), in milliseconds. A value of 0 means there is no delay in delivering messages to clients. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 60000),
				},
				Default: 0,
			},
			{
				SempName:            "elidingEnabled",
				TerraformName:       "eliding_enabled",
				MarkdownDescription: "Enable or disable message eliding for clients using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "elidingMaxTopicCount",
				TerraformName:       "eliding_max_topic_count",
				MarkdownDescription: "The maximum number of topics tracked for message eliding per client connection using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `256`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 32000),
				},
				Default: 256,
			},
			{
				SempName:            "eventClientProvisionedEndpointSpoolUsageThreshold",
				TerraformName:       "event_client_provisioned_endpoint_spool_usage_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.",
						Requires:            []string{"set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
						Default: 18,
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.",
						Requires:            []string{"clear_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
						},
						Default: 25,
					},
				},
			},
			{
				SempName:            "eventConnectionCountPerClientUsernameThreshold",
				TerraformName:       "event_connection_count_per_client_username_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 200000),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 200000),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventEgressFlowCountThreshold",
				TerraformName:       "event_egress_flow_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventEndpointCountPerClientUsernameThreshold",
				TerraformName:       "event_endpoint_count_per_client_username_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventIngressFlowCountThreshold",
				TerraformName:       "event_ingress_flow_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventServiceSmfConnectionCountPerClientUsernameThreshold",
				TerraformName:       "event_service_smf_connection_count_per_client_username_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 200000),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 200000),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventServiceWebConnectionCountPerClientUsernameThreshold",
				TerraformName:       "event_service_web_connection_count_per_client_username_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 200000),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 200000),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventSubscriptionCountThreshold",
				TerraformName:       "event_subscription_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventTransactedSessionCountThreshold",
				TerraformName:       "event_transacted_session_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventTransactionCountThreshold",
				TerraformName:       "event_transaction_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "maxConnectionCountPerClientUsername",
				TerraformName:       "max_connection_count_per_client_username",
				MarkdownDescription: "The maximum number of client connections per Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is the maximum value supported by the platform.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 200000),
				},
			},
			{
				SempName:            "maxEgressFlowCount",
				TerraformName:       "max_egress_flow_count",
				MarkdownDescription: "The maximum number of transmit flows that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 200000),
				},
				Default: 1000,
			},
			{
				SempName:            "maxEndpointCountPerClientUsername",
				TerraformName:       "max_endpoint_count_per_client_username",
				MarkdownDescription: "The maximum number of queues and topic endpoints that can be created by clients with the same Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 200000),
				},
				Default: 1000,
			},
			{
				SempName:            "maxIngressFlowCount",
				TerraformName:       "max_ingress_flow_count",
				MarkdownDescription: "The maximum number of receive flows that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 1000000),
				},
				Default: 1000,
			},
			{
				SempName:            "maxMsgsPerTransaction",
				TerraformName:       "max_msgs_per_transaction",
				MarkdownDescription: "The maximum number of publisher and consumer messages combined that is allowed within a transaction for each client associated with this client-profile. Exceeding this limit will result in a transaction prepare or commit failure. Changing this value during operation will not affect existing sessions. It is only validated at transaction creation time. Large transactions consume more resources and are more likely to require retrieving messages from the ADB or from disk to process the transaction prepare or commit requests. The transaction processing rate may diminish if a large number of messages must be retrieved from the ADB or from disk. Care should be taken to not use excessively large transactions needlessly to avoid exceeding resource limits and to avoid reducing the overall broker performance. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `256`. Available since 2.20.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 20000),
				},
				Default: 256,
			},
			{
				SempName:            "maxSubscriptionCount",
				TerraformName:       "max_subscription_count",
				MarkdownDescription: "The maximum number of subscriptions per client using the Client Profile. This limit is not enforced when a client adds a subscription to an endpoint, except for MQTT QoS 1 subscriptions. In addition, this limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default varies by platform.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 4294967295),
				},
			},
			{
				SempName:            "maxTransactedSessionCount",
				TerraformName:       "max_transacted_session_count",
				MarkdownDescription: "The maximum number of transacted sessions that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 30000),
				},
				Default: 10,
			},
			{
				SempName:            "maxTransactionCount",
				TerraformName:       "max_transaction_count",
				MarkdownDescription: "The maximum number of transactions that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default varies by platform.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 50000),
				},
			},
			{
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
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?]+$"), ""),
				},
			},
			{
				SempName:            "queueControl1MaxDepth",
				TerraformName:       "queue_control1_max_depth",
				MarkdownDescription: "The maximum depth of the \"Control 1\" (C-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				SempName:            "queueControl1MinMsgBurst",
				TerraformName:       "queue_control1_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Control 1\" (C-1) priority queue, regardless of the `queueControl1MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				SempName:            "queueDirect1MaxDepth",
				TerraformName:       "queue_direct1_max_depth",
				MarkdownDescription: "The maximum depth of the \"Direct 1\" (D-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				SempName:            "queueDirect1MinMsgBurst",
				TerraformName:       "queue_direct1_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Direct 1\" (D-1) priority queue, regardless of the `queueDirect1MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				SempName:            "queueDirect2MaxDepth",
				TerraformName:       "queue_direct2_max_depth",
				MarkdownDescription: "The maximum depth of the \"Direct 2\" (D-2) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				SempName:            "queueDirect2MinMsgBurst",
				TerraformName:       "queue_direct2_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Direct 2\" (D-2) priority queue, regardless of the `queueDirect2MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				SempName:            "queueDirect3MaxDepth",
				TerraformName:       "queue_direct3_max_depth",
				MarkdownDescription: "The maximum depth of the \"Direct 3\" (D-3) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				SempName:            "queueDirect3MinMsgBurst",
				TerraformName:       "queue_direct3_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Direct 3\" (D-3) priority queue, regardless of the `queueDirect3MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 262144),
				},
				Default: 4,
			},
			{
				SempName:            "queueGuaranteed1MaxDepth",
				TerraformName:       "queue_guaranteed1_max_depth",
				MarkdownDescription: "The maximum depth of the \"Guaranteed 1\" (G-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 262144),
				},
				Default: 20000,
			},
			{
				SempName:            "queueGuaranteed1MinMsgBurst",
				TerraformName:       "queue_guaranteed1_min_msg_burst",
				MarkdownDescription: "The number of messages that are always allowed entry into the \"Guaranteed 1\" (G-3) priority queue, regardless of the `queueGuaranteed1MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `255`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 262144),
				},
				Default: 255,
			},
			{
				SempName:            "rejectMsgToSenderOnNoSubscriptionMatchEnabled",
				TerraformName:       "reject_msg_to_sender_on_no_subscription_match_enabled",
				MarkdownDescription: "Enable or disable the sending of a negative acknowledgement (NACK) to a client using the Client Profile when discarding a guaranteed message due to no matching subscription found. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.2.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "replicationAllowClientConnectWhenStandbyEnabled",
				TerraformName:       "replication_allow_client_connect_when_standby_enabled",
				MarkdownDescription: "Enable or disable allowing clients using the Client Profile to connect to the Message VPN when its replication state is standby. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceMinKeepaliveTimeout",
				TerraformName:       "service_min_keepalive_timeout",
				MarkdownDescription: "The minimum client keepalive timeout which will be enforced for client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `30`. Available since 2.19.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(3, 3600),
				},
				Default: 30,
			},
			{
				SempName:            "serviceSmfMaxConnectionCountPerClientUsername",
				TerraformName:       "service_smf_max_connection_count_per_client_username",
				MarkdownDescription: "The maximum number of SMF client connections per Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is the maximum value supported by the platform.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 200000),
				},
			},
			{
				SempName:            "serviceSmfMinKeepaliveEnabled",
				TerraformName:       "service_smf_min_keepalive_enabled",
				MarkdownDescription: "Enable or disable the enforcement of a minimum keepalive timeout for SMF clients. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.19.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceWebInactiveTimeout",
				TerraformName:       "service_web_inactive_timeout",
				MarkdownDescription: "The timeout for inactive Web Transport client sessions using the Client Profile, in seconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `30`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 4294967295),
				},
				Default: 30,
			},
			{
				SempName:            "serviceWebMaxConnectionCountPerClientUsername",
				TerraformName:       "service_web_max_connection_count_per_client_username",
				MarkdownDescription: "The maximum number of Web Transport client connections per Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is the maximum value supported by the platform.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 200000),
				},
			},
			{
				SempName:            "serviceWebMaxPayload",
				TerraformName:       "service_web_max_payload",
				MarkdownDescription: "The maximum Web Transport payload size before fragmentation occurs for clients using the Client Profile, in bytes. The size of the header is not included. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(300, 10000000),
				},
				Default: 1e+06,
			},
			{
				SempName:            "tcpCongestionWindowSize",
				TerraformName:       "tcp_congestion_window_size",
				MarkdownDescription: "The TCP initial congestion window size for clients using the Client Profile, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact support before changing this value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `2`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 7826),
				},
				Default: 2,
			},
			{
				SempName:            "tcpKeepaliveCount",
				TerraformName:       "tcp_keepalive_count",
				MarkdownDescription: "The number of TCP keepalive retransmissions to a client using the Client Profile before declaring that it is not available. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `5`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 5),
				},
				Default: 5,
			},
			{
				SempName:            "tcpKeepaliveIdleTime",
				TerraformName:       "tcp_keepalive_idle_time",
				MarkdownDescription: "The amount of time a client connection using the Client Profile must remain idle before TCP begins sending keepalive probes, in seconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(3, 120),
				},
				Default: 3,
			},
			{
				SempName:            "tcpKeepaliveInterval",
				TerraformName:       "tcp_keepalive_interval",
				MarkdownDescription: "The amount of time between TCP keepalive retransmissions to a client using the Client Profile when no acknowledgement is received, in seconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 30),
				},
				Default: 1,
			},
			{
				SempName:            "tcpMaxSegmentSize",
				TerraformName:       "tcp_max_segment_size",
				MarkdownDescription: "The TCP maximum segment size for clients using the Client Profile, in bytes. Changes are applied to all existing connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1460`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(256, 1460),
				},
				Default: 1460,
			},
			{
				SempName:            "tcpMaxWindowSize",
				TerraformName:       "tcp_max_window_size",
				MarkdownDescription: "The TCP maximum window size for clients using the Client Profile, in kilobytes. Changes are applied to all existing connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `256`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(16, 65536),
				},
				Default: 256,
			},
			{
				SempName:            "tlsAllowDowngradeToPlainTextEnabled",
				TerraformName:       "tls_allow_downgrade_to_plain_text_enabled",
				MarkdownDescription: "Enable or disable allowing a client using the Client Profile to downgrade an encrypted connection to plain text. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since 2.8.",
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
