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
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "about_user",
		MarkdownDescription: "Session and access level information about the user accessing the SEMP API.\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.\n\nThis has been available since 2.2.",
		ObjectType:          broker.DataSourceObject,
		PathTemplate:        "/about/user",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "globalAccessLevel",
				TerraformName:       "global_access_level",
				MarkdownDescription: "The global access level of the User. The allowed values and their meaning are:\n\n<pre>\n\"admin\" - Full administrative access.\n\"none\" - No access.\n\"read-only\" - Read only access.\n\"read-write\" - Read and write access.\n</pre>\n",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("admin", "none", "read-only", "read-write"),
				},
			},
			{
				SempName:            "globalDmrBridgeAccessEnabled",
				TerraformName:       "global_dmr_bridge_access_enabled",
				MarkdownDescription: "Indicates whether global DMR Bridge access is enabled for the User. This is only for Solace internal use. This attribute may not be returned in a GET. Available since (hidden in public API).",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
			},
			{
				SempName:            "sessionActive",
				TerraformName:       "session_active",
				MarkdownDescription: "Indicates whether a session is active for this request. Available since 2.24.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
			},
			{
				SempName:            "sessionCreateTime",
				TerraformName:       "session_create_time",
				MarkdownDescription: "The timestamp of when the session was created. This attribute may not be returned in a GET. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			{
				SempName:            "sessionCurrentTime",
				TerraformName:       "session_current_time",
				MarkdownDescription: "The current server timestamp. This is provided as a reference point for the other timestamps provided. This attribute may not be returned in a GET. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			{
				SempName:            "sessionHardExpiryTime",
				TerraformName:       "session_hard_expiry_time",
				MarkdownDescription: "The hard expiry time for the session. After this time the session will be invalid, regardless of activity. This attribute may not be returned in a GET. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			{
				SempName:            "sessionId",
				TerraformName:       "session_id",
				MarkdownDescription: "An identifier for the session to differentiate this session from other sessions for the same user. This value is not guaranteed to be unique between active sessions for different users. This attribute may not be returned in a GET. Available since 2.21.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 56),
				},
			},
			{
				SempName:            "sessionIdleExpiryTime",
				TerraformName:       "session_idle_expiry_time",
				MarkdownDescription: "The session idle expiry time. After this time the session will be invalid if there has been no activity. This attribute may not be returned in a GET. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			{
				SempName:            "username",
				TerraformName:       "username",
				MarkdownDescription: "The username of the User. Available since 2.21.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 32),
				},
			},
		},
	}
	broker.RegisterDataSource(info)
}
