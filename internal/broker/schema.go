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

package broker

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"sort"
	"strings"
)

var DataSources []func() datasource.DataSource

var Entities []EntityInputs

func RegisterDataSource(inputs EntityInputs) {
	DataSources = append(DataSources, newBrokerDataSourceGenerator(inputs))
}

var Resources []func() resource.Resource

func RegisterResource(inputs EntityInputs) {
	Resources = append(Resources, newBrokerResourceGenerator(inputs))
	Entities = append(Entities, inputs)
}

var SempDetail SempVersionDetail

func RegisterSempVersionDetails(sempAPIBasePath string, sempVersion string, platform string) {
	SempDetail = SempVersionDetail{
		BasePath:    sempAPIBasePath,
		SempVersion: sempVersion,
		Platform:    platform,
	}
}

func addObjectConverters(attributes []*AttributeInfo) {
	for _, attr := range attributes {
		// if it is an object, we need to add a converter for it (simple attributes will already have converters)
		if attr.Attributes != nil {
			addObjectConverters(attr.Attributes)
			attr.Converter = NewObjectConverter(attr.TerraformName, attr.Attributes)
		}
	}
}

func modifiers[T any](requiresReplace bool, f func() T) []T {
	if requiresReplace {
		return []T{f()}
	}
	return nil
}

func terraformAttributeMap(attributes []*AttributeInfo, isResource bool, requiresReplace bool) map[string]schema.Attribute {
	tfAttributes := map[string]schema.Attribute{}
	for _, attr := range attributes {
		if attr.Sensitive && !isResource {
			// write-only attributes can't be retrieved so we don't expose them in the datasource
			continue
		}
		if !attr.Identifying && attr.ReadOnly && isResource {
			// read-only attributes should only be in the datasource
			continue
		}
		attrRequiresReplace := isResource && (requiresReplace || attr.RequiresReplace)
		if len(attr.Attributes) != 0 {
			childTypes := map[string]tftypes.Type{}
			for _, cAttr := range attr.Attributes {
				childTypes[cAttr.TerraformName] = cAttr.TerraformType
			}
			attr.TerraformType = tftypes.Object{
				AttributeTypes: childTypes,
			}
		}
		var deprecationMessage string
		if attr.Deprecated {
			deprecationMessage = "This attribute is deprecated."
		}
		switch attr.BaseType {
		case String:
			tfAttributes[attr.TerraformName] = schema.StringAttribute{
				Description:         attr.Description,
				MarkdownDescription: attr.MarkdownDescription,
				Required:            attr.Required && isResource || attr.Identifying,
				Optional:            !attr.Required && isResource,
				Computed:            !attr.Identifying && !isResource,
				Sensitive:           attr.Sensitive,
				DeprecationMessage:  deprecationMessage,
				Validators:          attr.StringValidators,
				PlanModifiers:       modifiers[planmodifier.String](attrRequiresReplace, stringplanmodifier.RequiresReplace),
			}
		case Int64:
			tfAttributes[attr.TerraformName] = schema.Int64Attribute{
				Description:         attr.Description,
				MarkdownDescription: attr.MarkdownDescription,
				Required:            attr.Required && isResource || attr.Identifying,
				Optional:            !attr.Required && isResource,
				Computed:            !attr.Identifying && !isResource,
				Sensitive:           attr.Sensitive,
				DeprecationMessage:  deprecationMessage,
				Validators:          attr.Int64Validators,
				PlanModifiers:       modifiers[planmodifier.Int64](attrRequiresReplace, int64planmodifier.RequiresReplace),
			}
		case Bool:
			tfAttributes[attr.TerraformName] = schema.BoolAttribute{
				Description:         attr.Description,
				MarkdownDescription: attr.MarkdownDescription,
				Required:            attr.Required && isResource || attr.Identifying,
				Optional:            !attr.Required && isResource,
				Computed:            !attr.Identifying && !isResource,
				Sensitive:           attr.Sensitive,
				DeprecationMessage:  deprecationMessage,
				Validators:          attr.BoolValidators,
				PlanModifiers:       modifiers[planmodifier.Bool](attrRequiresReplace, boolplanmodifier.RequiresReplace),
			}
		case Struct:
			tfAttributes[attr.TerraformName] = schema.SingleNestedAttribute{
				Attributes:          terraformAttributeMap(attr.Attributes, isResource, requiresReplace || attr.RequiresReplace),
				Description:         attr.Description,
				MarkdownDescription: attr.MarkdownDescription,
				Required:            attr.Required && isResource || attr.Identifying,
				Optional:            !attr.Required && isResource,
				Computed:            !attr.Identifying && !isResource,
				Sensitive:           attr.Sensitive,
				DeprecationMessage:  deprecationMessage,
				PlanModifiers:       modifiers[planmodifier.Object](attrRequiresReplace, objectplanmodifier.RequiresReplace),
			}
		}
	}
	return tfAttributes
}

type EntityInputs struct {
	TerraformName       string
	Description         string
	MarkdownDescription string
	DeprecationMessage  string
	ObjectType          objectType
	PathTemplate        string
	PostPathTemplate    string
	Version             int64
	Attributes          []*AttributeInfo
}

func newBrokerEntity(inputs EntityInputs, isResource bool) brokerEntity[schema.Schema] {
	addObjectConverters(inputs.Attributes)
	tfAttributes := terraformAttributeMap(inputs.Attributes, isResource, inputs.ObjectType == ReplaceOnlyObject)
	var identifyingAttributes []*AttributeInfo
	for _, attr := range inputs.Attributes {
		if attr.Identifying {
			identifyingAttributes = append(identifyingAttributes, attr)
		}
	}
	sort.Slice(identifyingAttributes, func(i, j int) bool {
		iAttr := identifyingAttributes[i]
		jAttr := identifyingAttributes[j]
		iIndex := strings.Index(inputs.PathTemplate, "{"+iAttr.SempName+"}")
		jIndex := strings.Index(inputs.PathTemplate, "{"+jAttr.SempName+"}")
		return iIndex < jIndex
	})
	s := schema.Schema{
		Attributes:          tfAttributes,
		Description:         inputs.Description,
		MarkdownDescription: inputs.MarkdownDescription,
		DeprecationMessage:  inputs.DeprecationMessage,
		Version:             inputs.Version,
	}
	return brokerEntity[schema.Schema]{
		schema: s,
		brokerEntityBase: brokerEntityBase{
			pathTemplate:          inputs.PathTemplate,
			postPathTemplate:      inputs.PostPathTemplate,
			terraformName:         inputs.TerraformName,
			objectType:            inputs.ObjectType,
			identifyingAttributes: identifyingAttributes,
			attributes:            inputs.Attributes,
			converter:             NewObjectConverter(inputs.TerraformName, inputs.Attributes),
		},
	}
}
