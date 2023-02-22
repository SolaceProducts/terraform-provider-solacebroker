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
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_authentication_oauth_provider",
		MarkdownDescription: "OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\naudience_claim_name|||x|\naudience_claim_source|||x|\naudience_claim_value|||x|\naudience_validation_enabled|||x|\nauthorization_group_claim_name|||x|\nauthorization_group_claim_source|||x|\nauthorization_group_enabled|||x|\ndisconnect_on_token_expiration_enabled|||x|\nenabled|||x|\njwks_refresh_interval|||x|\njwks_uri|||x|\nmsg_vpn_name|x||x|\noauth_provider_name|x||x|\ntoken_ignore_time_limits_enabled|||x|\ntoken_introspection_parameter_name|||x|\ntoken_introspection_password||x|x|x\ntoken_introspection_timeout|||x|\ntoken_introspection_uri|||x|\ntoken_introspection_username|||x|\nusername_claim_name|||x|\nusername_claim_source|||x|\nusername_validate_enabled|||x|\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been deprecated since 2.25. Replaced by authenticationoauth_profiles.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "audienceClaimName",
				TerraformName:       "audience_claim_name",
				MarkdownDescription: "The audience claim name, indicating which part of the object to use for determining the audience. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"aud\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "aud",
			},
			{
				SempName:            "audienceClaimSource",
				TerraformName:       "audience_claim_source",
				MarkdownDescription: "The audience claim source, indicating where to search for the audience value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"id-token\"`. The allowed values and their meaning are:\n\n<pre>\n\"access-token\" - The OAuth v2 access_token.\n\"id-token\" - The OpenID Connect id_token.\n\"introspection\" - The result of introspecting the OAuth v2 access_token.\n</pre>\n Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("access-token", "id-token", "introspection"),
				},
				Default: "id-token",
			},
			{
				SempName:            "audienceClaimValue",
				TerraformName:       "audience_claim_value",
				MarkdownDescription: "The required audience value for a token to be considered valid. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 64),
				},
				Default: "",
			},
			{
				SempName:            "audienceValidationEnabled",
				TerraformName:       "audience_validation_enabled",
				MarkdownDescription: "Enable or disable audience validation. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "authorizationGroupClaimName",
				TerraformName:       "authorization_group_claim_name",
				MarkdownDescription: "The authorization group claim name, indicating which part of the object to use for determining the authorization group. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"scope\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "scope",
			},
			{
				SempName:            "authorizationGroupClaimSource",
				TerraformName:       "authorization_group_claim_source",
				MarkdownDescription: "The authorization group claim source, indicating where to search for the authorization group name. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"id-token\"`. The allowed values and their meaning are:\n\n<pre>\n\"access-token\" - The OAuth v2 access_token.\n\"id-token\" - The OpenID Connect id_token.\n\"introspection\" - The result of introspecting the OAuth v2 access_token.\n</pre>\n Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("access-token", "id-token", "introspection"),
				},
				Default: "id-token",
			},
			{
				SempName:            "authorizationGroupEnabled",
				TerraformName:       "authorization_group_enabled",
				MarkdownDescription: "Enable or disable OAuth based authorization. When enabled, the configured authorization type for OAuth clients is overridden. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "disconnectOnTokenExpirationEnabled",
				TerraformName:       "disconnect_on_token_expiration_enabled",
				MarkdownDescription: "Enable or disable the disconnection of clients when their tokens expire. Changing this value does not affect existing clients, only new client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable OAuth Provider client authentication. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "jwksRefreshInterval",
				TerraformName:       "jwks_refresh_interval",
				MarkdownDescription: "The number of seconds between forced JWKS public key refreshing. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `86400`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(60, 31536000),
				},
				Default: 86400,
			},
			{
				SempName:            "jwksUri",
				TerraformName:       "jwks_uri",
				MarkdownDescription: "The URI where the OAuth provider publishes its JWKS public keys. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 2048),
				},
				Default: "",
			},
			{
				SempName:            "msgVpnName",
				TerraformName:       "msg_vpn_name",
				MarkdownDescription: "The name of the Message VPN. Deprecated since 2.25. Replaced by authenticationoauth_profiles.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?]+$"), ""),
				},
			},
			{
				SempName:            "oauthProviderName",
				TerraformName:       "oauth_provider_name",
				MarkdownDescription: "The name of the OAuth Provider. Deprecated since 2.25. Replaced by authenticationoauth_profiles.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 31),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^~]*$"), ""),
				},
			},
			{
				SempName:            "tokenIgnoreTimeLimitsEnabled",
				TerraformName:       "token_ignore_time_limits_enabled",
				MarkdownDescription: "Enable or disable whether to ignore time limits and accept tokens that are not yet valid or are no longer valid. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "tokenIntrospectionParameterName",
				TerraformName:       "token_introspection_parameter_name",
				MarkdownDescription: "The parameter name used to identify the token during access token introspection. A standards compliant OAuth introspection server expects \"token\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"token\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "token",
			},
			{
				SempName:            "tokenIntrospectionPassword",
				TerraformName:       "token_introspection_password",
				MarkdownDescription: "The password to use when logging into the token introspection URI. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Sensitive:           true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 64),
				},
				Default: "",
			},
			{
				SempName:            "tokenIntrospectionTimeout",
				TerraformName:       "token_introspection_timeout",
				MarkdownDescription: "The maximum time in seconds a token introspection is allowed to take. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 60),
				},
				Default: 1,
			},
			{
				SempName:            "tokenIntrospectionUri",
				TerraformName:       "token_introspection_uri",
				MarkdownDescription: "The token introspection URI of the OAuth authentication server. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 2048),
				},
				Default: "",
			},
			{
				SempName:            "tokenIntrospectionUsername",
				TerraformName:       "token_introspection_username",
				MarkdownDescription: "The username to use when logging into the token introspection URI. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "",
			},
			{
				SempName:            "usernameClaimName",
				TerraformName:       "username_claim_name",
				MarkdownDescription: "The username claim name, indicating which part of the object to use for determining the username. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"sub\"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 64),
				},
				Default: "sub",
			},
			{
				SempName:            "usernameClaimSource",
				TerraformName:       "username_claim_source",
				MarkdownDescription: "The username claim source, indicating where to search for the username value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"id-token\"`. The allowed values and their meaning are:\n\n<pre>\n\"access-token\" - The OAuth v2 access_token.\n\"id-token\" - The OpenID Connect id_token.\n\"introspection\" - The result of introspecting the OAuth v2 access_token.\n</pre>\n Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("access-token", "id-token", "introspection"),
				},
				Default: "id-token",
			},
			{
				SempName:            "usernameValidateEnabled",
				TerraformName:       "username_validate_enabled",
				MarkdownDescription: "Enable or disable whether the API provided username will be validated against the username calculated from the token(s); the connection attempt is rejected if they differ. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.",
				Deprecated:          true,
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
