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
	"fmt"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"math/big"
)

type Converter interface {
	ToTerraform(v any) (tftypes.Value, error)
	FromTerraform(v tftypes.Value) (any, error)
}

var _ Converter = &SimpleConverter[bool]{}
var _ Converter = &SimpleConverter[string]{}

type SimpleConverter[T any] struct {
	TerraformType tftypes.Type
}

func (c SimpleConverter[T]) ToTerraform(v any) (tftypes.Value, error) {
	err := tftypes.ValidateValue(c.TerraformType, v)
	if err != nil {
		return tftypes.Value{}, err
	}
	return tftypes.NewValue(c.TerraformType, v), nil
}

func (c SimpleConverter[T]) FromTerraform(v tftypes.Value) (any, error) {
	var result T
	err := v.As(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ Converter = &IntegerConverter{}

type IntegerConverter struct{}

func (IntegerConverter) ToTerraform(v any) (tftypes.Value, error) {
	err := tftypes.ValidateValue(tftypes.Number, v)
	if err != nil {
		return tftypes.Value{}, err
	}
	return tftypes.NewValue(tftypes.Number, v), nil
}

func (IntegerConverter) FromTerraform(v tftypes.Value) (any, error) {
	var f big.Float
	err := v.As(&f)
	if err != nil {
		return nil, err
	}
	i, _ := f.Int64()
	return i, nil
}

var _ Converter = &ObjectConverter{}

type ObjectConverter struct {
	terraformName string
	terraformType tftypes.Object
	attributes    []*AttributeInfo
}

func terraformTypes(attributes []*AttributeInfo) map[string]tftypes.Type {
	types := map[string]tftypes.Type{}
	for _, attr := range attributes {
		t := attr.TerraformType
		if attr.Attributes != nil {
			t = tftypes.Object{
				AttributeTypes: terraformTypes(attr.Attributes),
			}
		}
		types[attr.TerraformName] = t
	}
	return types
}

func NewObjectConverter(terraformName string, attributes []*AttributeInfo) *ObjectConverter {
	return &ObjectConverter{
		terraformName: terraformName,
		terraformType: tftypes.Object{AttributeTypes: terraformTypes(attributes)},
		attributes:    attributes,
	}
}

func (c *ObjectConverter) ToTerraform(v any) (tftypes.Value, error) {
	m, ok := v.(map[string]any)
	if !ok {
		return tftypes.Value{}, fmt.Errorf("unexpected type %T for attribute %v received in SEMP response; expected %T", v, c.terraformName, m)
	}
	tfData := map[string]tftypes.Value{}
	for _, sempAttribute := range c.attributes {
		v, ok := m[sempAttribute.SempName]
		if ok {
			tfv, err := sempAttribute.Converter.ToTerraform(v)
			if err != nil {
				return tftypes.Value{}, err
			}
			tfData[sempAttribute.TerraformName] = tfv
		} else {
			tfData[sempAttribute.TerraformName] = tftypes.NewValue(sempAttribute.TerraformType, nil)
		}
	}
	err := tftypes.ValidateValue(c.terraformType, tfData)
	if err != nil {
		return tftypes.Value{}, err
	}
	return tftypes.NewValue(c.terraformType, tfData), nil
}

func (c *ObjectConverter) FromTerraform(v tftypes.Value) (any, error) {
	sempData := map[string]any{}
	tfAttributes := map[string]tftypes.Value{}
	err := v.As(&tfAttributes)
	if err != nil {
		return nil, err
	}
	for _, attr := range c.attributes {
		v, ok := tfAttributes[attr.TerraformName]
		if ok && v.IsKnown() && !v.IsNull() {
			v, err := attr.Converter.FromTerraform(v)
			if err != nil {
				return nil, err
			}
			sempData[attr.SempName] = v
		}
	}
	return sempData, nil
}
