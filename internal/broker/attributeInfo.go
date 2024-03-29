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

package broker

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type objectType int8

const (
	DataSourceObject objectType = iota
	SingletonObject
	ReplaceOnlyObject
	StandardObject
)

type attributeType int8

const (
	String attributeType = iota
	Int64
	Bool
	Struct
)

type AttributeInfo struct {
	BaseType            attributeType
	SempName            string
	TerraformName       string
	Description         string
	MarkdownDescription string
	Identifying         bool
	Required            bool
	Sensitive           bool
	ReadOnly            bool
	RequiresReplace     bool
	Deprecated          bool
	Requires            []string
	ConflictsWith       []string
	Type                attr.Type
	TerraformType       tftypes.Type
	Attributes          []*AttributeInfo
	Converter           Converter
	StringValidators    []validator.String
	Int64Validators     []validator.Int64
	BoolValidators      []validator.Bool
	Default             any
}
