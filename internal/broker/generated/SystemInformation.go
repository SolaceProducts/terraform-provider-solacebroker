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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "system_information",
		MarkdownDescription: "The System Information object provides metadata about the SEMP API.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\nplatform|||x|\nsemp_version|||x|\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.\n\nThis has been deprecated since SEMP API version 2.2. /systemInformation was replaced by /about/api.",
		ObjectType:          broker.DataSourceObject,
		PathTemplate:        "/systemInformation",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:      broker.String,
				SempName:      "id",
				TerraformName: "id",
				Type:          types.StringType,
				TerraformType: tftypes.String,
				Converter:     broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Default:       "",
			},
			{
				BaseType:            broker.String,
				SempName:            "platform",
				TerraformName:       "platform",
				MarkdownDescription: "The platform running the SEMP API. Deprecated since SEMP API version 2.2. /systemInformation was replaced by /about/api.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 10),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "sempVersion",
				TerraformName:       "semp_version",
				MarkdownDescription: "The version of the SEMP API. Deprecated since SEMP API version 2.2. /systemInformation was replaced by /about/api.",
				ReadOnly:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 10),
				},
			},
		},
	}
	broker.RegisterDataSource(info)
}
