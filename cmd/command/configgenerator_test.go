package terraform

import (
	"context"
	"fmt"
	"reflect"
	"terraform-provider-solacebroker/internal/semp"
	"testing"
)

func TestCreateBrokerObjectRelationships(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBrokerObjectRelationships()
			println("Broker relationship object is " + fmt.Sprint(BrokerObjectRelationship))
		})
	}
}

func TestGetNameForResource(t *testing.T) {
	type args struct {
		resourceTerraformName      string
		attributeResourceTerraform string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNameForResource(tt.args.resourceTerraformName, tt.args.attributeResourceTerraform); got != tt.want {
				t.Errorf("GetNameForResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTerraformObject(t *testing.T) {
	type args struct {
		ctx                                        context.Context
		client                                     semp.Client
		resourceName                               string
		brokerObjectTerraformName                  string
		providerSpecificIdentifier                 string
		parentBrokerResourceAttributesRelationship map[string]string
		parentResult                               map[string]any
	}
	tests := []struct {
		name string
		args args
		want GeneratorTerraformOutput
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTerraformObject(tt.args.ctx, tt.args.client, tt.args.resourceName, tt.args.brokerObjectTerraformName, tt.args.providerSpecificIdentifier, tt.args.parentBrokerResourceAttributesRelationship, tt.args.parentResult); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTerraformObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_getDataSourceNameIfDatasource(t *testing.T) {
//	type args struct {
//		parent string
//		child  string
//	}
//	tests := []struct {
//		name  string
//		args  args
//		want  []string
//		want1 bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := getDataSourceNameIfDatasource(tt.args.parent, tt.args.child)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getDataSourceNameIfDatasource() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("getDataSourceNameIfDatasource() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
