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
		TerraformName:       "cert_authority",
		MarkdownDescription: "Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\ncertAuthorityName|x||x|\ncertContent|||x|\ncrlDayList|||x|\ncrlTimeList|||x|\ncrlUrl|||x|\nocspNonResponderCertEnabled|||x|\nocspOverrideUrl|||x|\nocspTimeout|||x|\nrevocationCheckEnabled|||x|\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/certAuthorities/{certAuthorityName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "certAuthorityName",
				TerraformName:       "cert_authority_name",
				MarkdownDescription: "The name of the Certificate Authority. Deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			{
				SempName:            "certContent",
				TerraformName:       "cert_content",
				MarkdownDescription: "The PEM formatted content for the trusted root certificate of a Certificate Authority. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 32768),
				},
				Default: "",
			},
			{
				SempName:            "crlDayList",
				TerraformName:       "crl_day_list",
				MarkdownDescription: "The scheduled CRL refresh day(s), specified as \"daily\" or a comma-separated list of days. Days must be specified as \"Sun\", \"Mon\", \"Tue\", \"Wed\", \"Thu\", \"Fri\", or \"Sat\", with no spaces, and in sorted order from Sunday to Saturday. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"daily\"`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
				Deprecated:          true,
				Requires:            []string{"crl_time_list"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("crl_time_list"),
					),
					stringvalidator.LengthBetween(0, 100),
				},
				Default: "daily",
			},
			{
				SempName:            "crlTimeList",
				TerraformName:       "crl_time_list",
				MarkdownDescription: "The scheduled CRL refresh time(s), specified as \"hourly\" or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"3:00\"`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
				Deprecated:          true,
				Requires:            []string{"crl_day_list"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("crl_day_list"),
					),
					stringvalidator.LengthBetween(0, 100),
				},
				Default: "3:00",
			},
			{
				SempName:            "crlUrl",
				TerraformName:       "crl_url",
				MarkdownDescription: "The URL for the CRL source. This is a required attribute for CRL to be operational and the URL must be complete with http:// included. IPv6 addresses must be enclosed in square-brackets. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as revocationCheckEnabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 2048),
					stringvalidator.RegexMatches(regexp.MustCompile("^(.+://.+)?$"), ""),
				},
				Default: "",
			},
			{
				SempName:            "ocspNonResponderCertEnabled",
				TerraformName:       "ocsp_non_responder_cert_enabled",
				MarkdownDescription: "Enable or disable allowing a non-responder certificate to sign an OCSP response. Typically used with an OCSP override URL in cases where a single certificate is used to sign client certificates and OCSP responses. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
				Deprecated:          true,
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "ocspOverrideUrl",
				TerraformName:       "ocsp_override_url",
				MarkdownDescription: "The OCSP responder URL to use for overriding the one supplied in the client certificate. The URL must be complete with http:// included. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 2048),
					stringvalidator.RegexMatches(regexp.MustCompile("^((.|..|...|....|[^h]....|.[^t]...|..[^t]..|...[^p].|....[^s]|.......*)://.+)?$"), ""),
				},
				Default: "",
			},
			{
				SempName:            "ocspTimeout",
				TerraformName:       "ocsp_timeout",
				MarkdownDescription: "The timeout in seconds to receive a response from the OCSP responder after sending a request or making the initial connection attempt. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `5`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
				Deprecated:          true,
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 86400),
				},
				Default: 5,
			},
			{
				SempName:            "revocationCheckEnabled",
				TerraformName:       "revocation_check_enabled",
				MarkdownDescription: "Enable or disable Certificate Authority revocation checking. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities.",
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
