// terraform-provider-solacebroker
//
// Copyright 2024 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package generator

import (
	"reflect"
	"testing"
)

func TestStringWithDefaultFromEnv(t *testing.T) {
	type args struct {
		name        string
		isMandatory bool
		fallback    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"GetDefaultValue",
			args{
				name:        "REGISTRY",
				isMandatory: false,
				fallback:    "Test",
			},
			"Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringWithDefaultFromEnv(tt.args.name, tt.args.isMandatory, tt.args.fallback); got != tt.want {
				t.Errorf("StringWithDefaultFromEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAttributeInfo(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want ResourceAttributeInfo
	}{
		{
			"AttributeTest",
			args{value: "msg_vpn"},
			ResourceAttributeInfo{
				AttributeValue: "msg_vpn",
				Comment:        "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAttributeInfo(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAttributeInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeValidForTerraformIdentifier(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"SanitizeTextStartingWithNumber",
			args{name: "1testing"},
			"-testing",
		},
		{
			"SanitizeTextContainingSpecialCharacters",
			args{name: "*testing*"},
			"-testing-",
		},
		{
			"SanitizeTextContainingSpecialCharactersTwo",
			args{name: "#testing/"},
			"-testing-",
		},
		{
			"SanitizeTextContainingSpecialCharactersThree",
			args{name: "$testing\""},
			"-testing-",
		},
		{
			"SanitizeTextContainingSpecialCharactersFour",
			args{name: "%testing^"},
			"-testing-",
		},
		{
			"SanitizeTextContainingSpecialCharactersFive",
			args{name: "%testing^"},
			"-testing-",
		},
		{
			"SanitizeTextEmpty",
			args{name: ""},
			"",
		},
		{
			"SanitizeTextContainingEmpty",
			args{name: " "},
			"-",
		},
		{
			"SanitizeTextOnlySpecialCharacter",
			args{name: "#"},
			"-",
		},
		{
			"SanitizeTextOnlySpecialCharacterTwo",
			args{name: "\\"},
			"-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeValidForTerraformIdentifier(tt.args.name); got != tt.want {
				t.Errorf("makeValidForTerraformIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSanitizeHclValue(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"SanitizeValueStartingWithNumber",
			args{name: "1testing"},
			"1testing",
		},
		{
			"SanitizeValueContainingSpecialCharacters",
			args{name: "*testing*"},
			"*testing*",
		},
		{
			"SanitizeJSONValue",
			args{name: "{\"as\":\"ds\"}"},
			"{\\\"as\\\":\\\"ds\\\"}",
		},
		{
			"SanitizeOnlySpecialCharacterValue",
			args{name: "\\"},
			"\\\\",
		},
		{
			"SanitizeOnlySpecialCharacterValueTwo",
			args{name: "*"},
			"*",
		},
		{
			"SanitizeOnlySpecialCharacterValueThree",
			args{name: "\""},
			"\\\"",
		},
		{
			"SanitizeSubstituitionExpression",
			args{name: "time/${now()}"},
			"time/$${now()}",
		},
		{
			"SanitizeSubstituitionExpression",
			args{name: "time/%{now()}"},
			"time/%%{now()}",
		},
		{
			"SanitizeSubstituitionExpression",
			args{name: "time/$/test"},
			"time/$/test",
		},
		{
			"SanitizeSubstituitionExpression",
			args{name: "time/%/test"},
			"time/%/test",
		},
		{
			"SanitizeSubstituitionExpression",
			args{name: "${"},
			"$${",
		},
		{
			"SanitizeSubstituitionExpression",
			args{name: "${utcDate(\"/\")}/${utcTime(\"/\")}/${BASE32(randomBytes(15))}"},
			"$${utcDate(\\\"/\\\")}/$${utcTime(\\\"/\\\")}/$${BASE32(randomBytes(15))}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeHclStringValue(tt.args.name); got != tt.want {
				t.Errorf("SanitizeHclStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
