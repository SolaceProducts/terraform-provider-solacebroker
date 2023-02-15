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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "oauth_profile_access_level_group_msg_vpn_access_level_exception",
		MarkdownDescription: "Message VPN access-level exceptions for members of this group.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\ngroupName|x|||\nmsgVpnName|x|||\noauthProfileName|x|||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been available since 2.24.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName}/msgVpnAccessLevelExceptions/{msgVpnName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "accessLevel",
				TerraformName:       "access_level",
				MarkdownDescription: "The message VPN access level. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"none\"`. The allowed values and their meaning are:\n\n<pre>\n\"none\" - User has no access to a Message VPN.\n\"read-only\" - User has read-only access to a Message VPN.\n\"read-write\" - User has read-write access to most Message VPN settings.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("none", "read-only", "read-write"),
				},
				Default: "none",
			},
			{
				SempName:            "groupName",
				TerraformName:       "group_name",
				MarkdownDescription: "The name of the group.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			{
				SempName:            "msgVpnName",
				TerraformName:       "msg_vpn_name",
				MarkdownDescription: "The name of the message VPN.",
				Identifying:         true,
				Required:            true,
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
				SempName:            "oauthProfileName",
				TerraformName:       "oauth_profile_name",
				MarkdownDescription: "The name of the OAuth profile.",
				Identifying:         true,
				Required:            true,
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
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
