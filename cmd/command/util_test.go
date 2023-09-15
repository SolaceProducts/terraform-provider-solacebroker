package terraform

import (
	"reflect"
	"terraform-provider-solacebroker/internal/broker"
	"testing"
	"time"
)

func TestBooleanWithDefaultFromEnv(t *testing.T) {
	type args struct {
		name        string
		isMandatory bool
		fallback    bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BooleanWithDefaultFromEnv(tt.args.name, tt.args.isMandatory, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("BooleanWithDefaultFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BooleanWithDefaultFromEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDurationWithDefaultFromEnv(t *testing.T) {
	type args struct {
		name        string
		isMandatory bool
		fallback    time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DurationWithDefaultFromEnv(tt.args.name, tt.args.isMandatory, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("DurationWithDefaultFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DurationWithDefaultFromEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		n int
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
			if got := GenerateRandomString(tt.args.n); got != tt.want {
				t.Errorf("GenerateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateTerraformString(t *testing.T) {
	type args struct {
		attributes                     []*broker.AttributeInfo
		values                         []map[string]interface{}
		parentBrokerResourceAttributes map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateTerraformString(tt.args.attributes, tt.args.values, tt.args.parentBrokerResourceAttributes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateTerraformString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTerraformString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetParentResourceAttributes(t *testing.T) {
	type args struct {
		brokerParentResource map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetParentResourceAttributes(tt.args.brokerParentResource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParentResourceAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64WithDefaultFromEnv(t *testing.T) {
	type args struct {
		name        string
		isMandatory bool
		fallback    int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64WithDefaultFromEnv(tt.args.name, tt.args.isMandatory, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64WithDefaultFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64WithDefaultFromEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogCLIError(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"TestWritingErrorLog",
			args{
				"Sample write",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogCLIError(tt.args.err)
		})
	}
}

func TestLogCLIInfo(t *testing.T) {
	type args struct {
		info string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"TestWritingInfoLog",
			args{
				"Sample write",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogCLIInfo(tt.args.info)
		})
	}
}

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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringWithDefaultFromEnv(tt.args.name, tt.args.isMandatory, tt.args.fallback); got != tt.want {
				t.Errorf("StringWithDefaultFromEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randStr(t *testing.T) {
	type args struct {
		n int
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
			if got := randStr(tt.args.n); got != tt.want {
				t.Errorf("randStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
