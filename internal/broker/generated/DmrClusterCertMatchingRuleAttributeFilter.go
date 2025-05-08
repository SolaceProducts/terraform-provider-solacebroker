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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "dmr_cluster_cert_matching_rule_attribute_filter",
		MarkdownDescription: "A Cert Matching Rule Attribute Filter compares a link attribute to a string.\n\n\n\nThe minimum access scope/level required to perform this operation is \"global/read-only\".\n\nThis has been available since SEMP API version 2.28.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/attributeFilters/{filterName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "attributeName",
				TerraformName:       "attribute_name",
				MarkdownDescription: "Link Attribute to be tested.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/mesh-manager\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^([A-Za-z][A-Za-z0-9\\-]*)?$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "attributeValue",
				TerraformName:       "attribute_value",
				MarkdownDescription: "Expected attribute value.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\". The minimum access scope/level required to change this attribute is \"global/mesh-manager\". Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "dmrClusterName",
				TerraformName:       "dmr_cluster_name",
				MarkdownDescription: "The name of the Cluster.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\".",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "filterName",
				TerraformName:       "filter_name",
				MarkdownDescription: "The name of the filter.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\".",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^#*? ]([^*?]*[^*? ])?$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "ruleName",
				TerraformName:       "rule_name",
				MarkdownDescription: "The name of the rule.\n\nThe minimum access scope/level required to retrieve this attribute is \"global/read-only\".",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
