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
	"reflect"
	"terraform-provider-solacebroker/internal/semp"

	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var resourceToDataSourceTypes = map[reflect.Type]reflect.Type{
	reflect.TypeOf(rschema.StringAttribute{}):       reflect.TypeOf(dschema.StringAttribute{}),
	reflect.TypeOf(rschema.Int64Attribute{}):        reflect.TypeOf(dschema.Int64Attribute{}),
	reflect.TypeOf(rschema.BoolAttribute{}):         reflect.TypeOf(dschema.BoolAttribute{}),
	reflect.TypeOf(rschema.SingleNestedAttribute{}): reflect.TypeOf(dschema.SingleNestedAttribute{}),
}

type brokerEntity[T rschema.Schema | dschema.Schema] struct {
	schema T
	brokerEntityBase
}

type brokerEntityBase struct {
	pathTemplate          string
	postPathTemplate      string
	terraformName         string
	objectType            objectType
	identifyingAttributes []*AttributeInfo
	attributes            []*AttributeInfo
	converter             *ObjectConverter
	client                *semp.Client
}

func copyMatchingFields(prefix string, in reflect.Value, out reflect.Value) {
	tIn := in.Type()
	tOut := out.Type()
	for i := 0; i < in.NumField(); i++ {
		tIn := tIn.Field(i)
		tOut, tOutFound := tOut.FieldByName(tIn.Name)
		if tOutFound {
			fIn := in.Field(i)
			if fIn.IsZero() {
				continue
			}
			fOut := out.FieldByName(tIn.Name)
			if tIn.Type.AssignableTo(tOut.Type) {
				fOut.Set(fIn)
			} else {
				if tIn.Type.Kind() == reflect.Map {
					fOut.Set(reflect.MakeMap(fOut.Type()))
					iter := fIn.MapRange()
					for iter.Next() {
						k := iter.Key()
						vIn := iter.Value()
						iIn := vIn.Interface()
						tvIn := reflect.TypeOf(iIn)
						tvOut := resourceToDataSourceTypes[tvIn]
						vOut := reflect.New(tvOut).Elem()
						copyMatchingFields(fmt.Sprintf("%v.%v", prefix, k), reflect.ValueOf(iIn), vOut)
						fOut.SetMapIndex(k, vOut)

					}
				}
			}
		}
	}
}

func resourceEntityToDataSourceEntity(entity brokerEntity[rschema.Schema]) brokerEntity[dschema.Schema] {
	ds := dschema.Schema{}
	in := reflect.ValueOf(&entity.schema).Elem()
	out := reflect.ValueOf(&ds).Elem()
	copyMatchingFields(entity.terraformName, in, out)
	return brokerEntity[dschema.Schema]{
		schema:           ds,
		brokerEntityBase: entity.brokerEntityBase,
	}
}
