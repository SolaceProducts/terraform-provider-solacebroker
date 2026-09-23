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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_telemetry_profile_trace_filter_subscription",
		MarkdownDescription: "Trace filter subscriptions control which messages will be attracted by the tracing filter.\n\n\n\nThe minimum access scope/level required to perform this operation is \"vpn/read-only\".\n\nThis has been available since SEMP API version 2.31.",
		ObjectType:          broker.ReplaceOnlyObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions/{subscription},{subscriptionSyntax}",
		PostPathTemplate:    "/msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions",
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
				BaseType:            broker.String,
				SempName:            "subscription",
				TerraformName:       "subscription",
				MarkdownDescription: "Messages matching this subscription will follow this filter's configuration.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
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
			{
				BaseType:            broker.String,
				SempName:            "subscriptionSyntax",
				TerraformName:       "subscription_syntax",
				MarkdownDescription: "The syntax of the trace filter subscription.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\". The allowed values and their meaning are:\n\n<pre>\n\"smf\" - Subscription uses SMF syntax.\n\"mqtt\" - Subscription uses MQTT syntax.\n</pre>\n",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("smf", "mqtt"),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "telemetryProfileName",
				TerraformName:       "telemetry_profile_name",
				MarkdownDescription: "The name of the Telemetry Profile.\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
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
				BaseType:            broker.String,
				SempName:            "traceFilterName",
				TerraformName:       "trace_filter_name",
				MarkdownDescription: "A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \"appNameDebug\".\n\nThe minimum access scope/level required to retrieve this attribute is \"vpn/read-only\".",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 127),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^#*? ]([^*?]*[^*? ])?$"), ""),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
