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
		TerraformName:       "msg_vpn_replay_log",
		MarkdownDescription: "When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.\n\n\n\nThe minimum access scope/level required to perform this operation is \"vpn/read-only\".\n\nThis has been available since SEMP API version 2.10.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/replayLogs/{replayLogName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.Bool,
				SempName:            "egressEnabled",
				TerraformName:       "egress_enabled",
				MarkdownDescription: "Enable or disable the transmission of messages from the Replay Log.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"global/mesh-manager or vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "ingressEnabled",
				TerraformName:       "ingress_enabled",
				MarkdownDescription: "Enable or disable the reception of messages to the Replay Log.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"global/mesh-manager or vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxSpoolUsage",
				TerraformName:       "max_spool_usage",
				MarkdownDescription: "The maximum spool usage allowed by the Replay Log, in megabytes (MB). If this limit is exceeded, old messages will be trimmed.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"global/mesh-manager or vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 6000000),
				},
				Default: 0,
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
				SempName:            "replayLogName",
				TerraformName:       "replay_log_name",
				MarkdownDescription: "The name of the Replay Log.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 185),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&;]+$"), ""),
				},
			},
			{
				BaseType:            broker.Bool,
				SempName:            "topicFilterEnabled",
				TerraformName:       "topic_filter_enabled",
				MarkdownDescription: "Enable or disable topic filtering for the Replay Log.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The minimum access scope/level required to change this attribute is \"global/mesh-manager or vpn/read-write\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since SEMP API version 2.27.",
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
