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
		TerraformName:       "oauth_profile_client_required_claim",
		MarkdownDescription: "Additional claims to be verified in the ID token.\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.24.",
		ObjectType:          broker.ReplaceOnlyObject,
		PathTemplate:        "/oauthProfiles/{oauthProfileName}/clientRequiredClaims/{clientRequiredClaimName}",
		PostPathTemplate:    "/oauthProfiles/{oauthProfileName}/clientRequiredClaims",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "clientRequiredClaimName",
				TerraformName:       "client_required_claim_name",
				MarkdownDescription: "The name of the ID token claim to verify.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "clientRequiredClaimValue",
				TerraformName:       "client_required_claim_value",
				MarkdownDescription: "The required claim value, which must be a string containing a valid JSON value.",
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
				},
			},
			{
				BaseType:            broker.String,
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
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9_]+$"), ""),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
