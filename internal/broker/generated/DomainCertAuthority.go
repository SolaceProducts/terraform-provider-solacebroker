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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "domain_cert_authority",
		MarkdownDescription: "Certificate Authorities trusted for domain verification.\n\n\nAttribute|Identifying\n:---|:---:\ncert_authority_name|x\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.19.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/domainCertAuthorities/{certAuthorityName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "certAuthorityName",
				TerraformName:       "cert_authority_name",
				MarkdownDescription: "The name of the Certificate Authority.",
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
				MarkdownDescription: "The PEM formatted content for the trusted root certificate of a domain Certificate Authority. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32768),
				},
				Default: "",
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
