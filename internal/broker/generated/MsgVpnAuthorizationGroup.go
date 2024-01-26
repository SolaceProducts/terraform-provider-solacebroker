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
		TerraformName:       "msg_vpn_authorization_group",
		MarkdownDescription: "To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.\n\n\nAttribute|Identifying|Write-Only\n:---|:---:|:---:\nauthorization_group_name|x|\nmsg_vpn_name|x|\norder_after_authorization_group_name||x\norder_before_authorization_group_name||x\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.0.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "aclProfileName",
				TerraformName:       "acl_profile_name",
				MarkdownDescription: "The ACL Profile of the Authorization Group. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"default\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
				Default: "default",
			},
			{
				BaseType:            broker.String,
				SempName:            "authorizationGroupName",
				TerraformName:       "authorization_group_name",
				MarkdownDescription: "The name of the Authorization Group. For LDAP groups, special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 256),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "clientProfileName",
				TerraformName:       "client_profile_name",
				MarkdownDescription: "The Client Profile of the Authorization Group. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"default\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^#?[A-Za-z0-9\\-_]+$"), ""),
				},
				Default: "default",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable the Authorization Group in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
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
				SempName:            "orderAfterAuthorizationGroupName",
				TerraformName:       "order_after_authorization_group_name",
				MarkdownDescription: "Lower the priority to be less than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
				Sensitive:           true,
				ConflictsWith:       []string{"order_before_authorization_group_name"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName("order_before_authorization_group_name"),
					),
					stringvalidator.LengthBetween(1, 256),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "orderBeforeAuthorizationGroupName",
				TerraformName:       "order_before_authorization_group_name",
				MarkdownDescription: "Raise the priority to be greater than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is not applicable.",
				Sensitive:           true,
				ConflictsWith:       []string{"order_after_authorization_group_name"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName("order_after_authorization_group_name"),
					),
					stringvalidator.LengthBetween(1, 256),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
