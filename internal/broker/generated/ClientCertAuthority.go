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
		TerraformName:       "client_cert_authority",
		MarkdownDescription: "Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.\n\n\n\nThe minimum access scope/level required to perform this operation is \"global/read-only\".\n\nThis has been available since SEMP API version 2.19.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/clientCertAuthorities/{certAuthorityName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "certAuthorityName",
				TerraformName:       "cert_authority_name",
				MarkdownDescription: "The name of the Certificate Authority.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9_\\-.]+$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "certContent",
				TerraformName:       "cert_content",
				MarkdownDescription: "The PEM formatted content for the trusted root certificate of a client Certificate Authority.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32768),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "crlDayList",
				TerraformName:       "crl_day_list",
				MarkdownDescription: "The scheduled CRL refresh day(s), specified as \"daily\" or a comma-separated list of days. Days must be specified as \"Sun\", \"Mon\", \"Tue\", \"Wed\", \"Thu\", \"Fri\", or \"Sat\", with no spaces, and in sorted order from Sunday to Saturday. The empty-string (\"\") can also be specified, indicating no schedule is configured (\"crl_time_list\" must also be configured to the empty-string).\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"daily\"`.",
				Requires:            []string{"crl_time_list"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("crl_time_list"),
					),
					stringvalidator.LengthBetween(0, 100),
				},
				Default: "daily",
			},
			{
				BaseType:            broker.String,
				SempName:            "crlTimeList",
				TerraformName:       "crl_time_list",
				MarkdownDescription: "The scheduled CRL refresh time(s), specified as \"hourly\" or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times (up to 4) must be in sorted order from 0:00 to 23:59. The empty-string (\"\") can also be specified, indicating no schedule is configured (\"crl_day_list\" must also be configured to the empty-string).\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"3:00\"`.",
				Requires:            []string{"crl_day_list"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("crl_day_list"),
					),
					stringvalidator.LengthBetween(0, 100),
				},
				Default: "3:00",
			},
			{
				BaseType:            broker.String,
				SempName:            "crlUrl",
				TerraformName:       "crl_url",
				MarkdownDescription: "The URL for the CRL source. This is a required attribute for CRL to be operational and the URL must be complete with http:// included. IPv6 addresses must be enclosed in square-brackets.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as revocation_check_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 2048),
					stringvalidator.RegexMatches(regexp.MustCompile("^(.+://.+)?$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "ocspNonResponderCertEnabled",
				TerraformName:       "ocsp_non_responder_cert_enabled",
				MarkdownDescription: "Enable or disable allowing a non-responder certificate to sign an OCSP response. Typically used with an OCSP override URL in cases where a single certificate is used to sign client certificates and OCSP responses.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "ocspOverrideUrl",
				TerraformName:       "ocsp_override_url",
				MarkdownDescription: "The OCSP responder URL to use for overriding the one supplied in the client certificate. The URL must be complete with http:// included.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 2048),
					stringvalidator.RegexMatches(regexp.MustCompile("^((.|..|...|....|[^h]....|.[^t]...|..[^t]..|...[^p].|....[^s]|.......*)://.+)?$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "ocspTimeout",
				TerraformName:       "ocsp_timeout",
				MarkdownDescription: "The timeout in seconds to receive a response from the OCSP responder after sending a request or making the initial connection attempt.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `5`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: 5,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "revocationCheckEnabled",
				TerraformName:       "revocation_check_enabled",
				MarkdownDescription: "Enable or disable Certificate Authority revocation checking.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/admin\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
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
