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
package terraform

import (
	"testing"
)

func TestCreateBrokerObjectRelationships(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Generate Broker Relationship"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBrokerObjectRelationships()
			if len(BrokerObjectRelationship) == 0 {
				t.Errorf("Broker relationship not built ")
			}
			_, exist := BrokerObjectRelationship["msg_vpn"]
			if !exist {
				t.Errorf("Broker relationship does not contain msgVPn relation")
			}
		})
	}
}

func TestGetNameForResource(t *testing.T) {
	type args struct {
		resourceTerraformName      string
		attributeResourceTerraform ResourceConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"GetNameForMsgVPN",
			args{
				resourceTerraformName: "solacebroker_msg_vpn qn",
				attributeResourceTerraform: ResourceConfig{
					ResourceAttributes: map[string]ResourceAttributeInfo{"msg_vpn_name": {
						"test",
						"no comment",
					},
						"ingress": {
							"0",
							"comment here",
						},
					},
				},
			},
			"_test",
		},
		{
			"GetNameForAclProfile",
			args{
				resourceTerraformName: "solacebroker_msg_vpn_acl_profile acl",
				attributeResourceTerraform: ResourceConfig{
					ResourceAttributes: map[string]ResourceAttributeInfo{"acl_profile_name": {
						"default",
						"no comment",
					},
						"random": {
							"0",
							"comment here",
						},
					},
				},
			},
			"_default",
		},
		{
			"GetNameForTopicName",
			args{
				resourceTerraformName: "solacebroker_msg_vpn_jndi_topic topic",
				attributeResourceTerraform: ResourceConfig{
					ResourceAttributes: map[string]ResourceAttributeInfo{"topic_name": {
						"random",
						"no comment",
					},
						"mock": {
							"parameter",
							"comment here",
						},
					},
				},
			},
			"_random",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNameForResource(tt.args.resourceTerraformName, tt.args.attributeResourceTerraform); got != tt.want {
				t.Errorf("GetNameForResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
