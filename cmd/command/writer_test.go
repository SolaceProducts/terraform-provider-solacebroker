package terraform

import "testing"

func TestGenerateTerraformFile(t *testing.T) {
	type args struct {
		terraformObjectInfo *ObjectInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateTerraformFile(tt.args.terraformObjectInfo); (err != nil) != tt.wantErr {
				t.Errorf("GenerateTerraformFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
