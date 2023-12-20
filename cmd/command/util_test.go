// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
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
package terraform

import (
	"reflect"
	"testing"
)

func TestResolveSempPath(t *testing.T) {
	type args struct {
		pathTemplate string
		v            string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"MSGVPNParsing",
			args{
				pathTemplate: "/msgVpns/{msgVpnName}",
				v:            "Test",
			},
			"/msgVpns/Test",
			false,
		},
		{
			"MSGVPNParsing",
			args{
				pathTemplate: "/msgVpns/{msgVpnName}/{anotherMock}",
				v:            "Test/Mock",
			},
			"/msgVpns/Test/Mock",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveSempPath(tt.args.pathTemplate, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveSempPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResolveSempPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveSempPathWithParent(t *testing.T) {
	type args struct {
		pathTemplate string
		parentValues map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"MSGVPNParsing",
			args{
				pathTemplate: "/msgVpns/{msgVpnName}",
				parentValues: map[string]any{"msgVpnName": "Test"},
			},
			"/msgVpns/Test",
			false,
		},
		{
			"Parsing where all values not available",
			args{
				pathTemplate: "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}",
				parentValues: map[string]any{"msgVpnName": "Test"},
			},
			"/msgVpns/Test/bridges",
			false,
		},
		{
			"Parsing with commas",
			args{
				pathTemplate: "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}",
				parentValues: map[string]any{"msgVpnName": "Test", "bridgeName": "TestBridge", "bridgeVirtualRouter": "TestingBridgeRouter"},
			},
			"/msgVpns/Test/bridges/TestBridge,TestingBridgeRouter",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveSempPathWithParent(tt.args.pathTemplate, tt.args.parentValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveSempPathWithParent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResolveSempPathWithParent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_addCommentToAttributeInfo(t *testing.T) {
	type args struct {
		info    ResourceAttributeInfo
		comment string
	}
	tests := []struct {
		name string
		args args
		want ResourceAttributeInfo
	}{
		{
			"TestCommentAdd",
			args{
				info: ResourceAttributeInfo{
					"test",
					"",
				},
			},
			ResourceAttributeInfo{
				AttributeValue: "test",
				Comment:        "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addCommentToAttributeInfo(tt.args.info, tt.args.comment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addCommentToAttributeInfo() = %v, want %v", got, tt.want)
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

func TestSanitizeHclIdentifierName(t *testing.T) {
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
			"gn_1testing",
		},
		{
			"SanitizeTextContainingSpecialCharacters",
			args{name: "*testing*"},
			"_testing_",
		},
		{
			"SanitizeTextContainingSpecialCharactersTwo",
			args{name: "#testing/"},
			"_testing_",
		},
		{
			"SanitizeTextContainingSpecialCharactersThree",
			args{name: "$testing\""},
			"_testing_",
		},
		{
			"SanitizeTextContainingSpecialCharactersFour",
			args{name: "%testing^"},
			"_testing_",
		},
		{
			"SanitizeTextContainingSpecialCharactersFive",
			args{name: "%testing^"},
			"_testing_",
		},
		{
			"SanitizeTextEmpty",
			args{name: ""},
			"gn_",
		},
		{
			"SanitizeTextContainingEmpty",
			args{name: " "},
			"gn_",
		},
		{
			"SanitizeTextOnlySpecialCharacter",
			args{name: "#"},
			"gn__",
		},
		{
			"SanitizeTextOnlySpecialCharacterTwo",
			args{name: "\\"},
			"gn__",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeHclIdentifierName(tt.args.name); got != tt.want {
				t.Errorf("SanitizeHclIdentifierName() = %v, want %v", got, tt.want)
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
