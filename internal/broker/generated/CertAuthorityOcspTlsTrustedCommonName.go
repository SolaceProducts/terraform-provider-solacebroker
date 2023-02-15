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
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "cert_authority_ocsp_tls_trusted_common_name",
		MarkdownDescription: "When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\ncertAuthorityName|x||x|\nocspTlsTrustedCommonName|x||x|\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been deprecated since 2.19. Replaced by clientCertAuthorities.",
		ObjectType:          broker.ReplaceOnlyObject,
		PathTemplate:        "/certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}",
		PostPathTemplate:    "/certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "certAuthorityName",
				TerraformName:       "cert_authority_name",
				MarkdownDescription: "The name of the Certificate Authority. Deprecated since 2.19. Replaced by clientCertAuthorities.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
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
				SempName:            "ocspTlsTrustedCommonName",
				TerraformName:       "ocsp_tls_trusted_common_name",
				MarkdownDescription: "The expected Trusted Common Name of the OCSP responder remote certificate. Deprecated since 2.19. Replaced by clientCertAuthorities.",
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
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
