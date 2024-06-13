package broker

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestClient(t *testing.T) {
	matrix := []struct {
		ParamUsername    string
		ParamPassword    string
		ParamBearertoken string
		EnvUsername      string
		EnvPassword      string
		EnvBearertoken   string
		Expected         string
	}{
		{"testuser", "testpassword", "testbearertoken", "", "", "", "Cannot use Bearer token with basic authentication credentials"},
		{"", "testpassword", "testbearertoken", "", "", "", "Cannot use Bearer token with basic authentication credentials"},
		{"testuser", "", "testbearertoken", "", "", "", "Cannot use Bearer token with basic authentication credentials"},
		{"", "", "", "testuser", "testpassword", "testbearertoken", "Cannot use Bearer token with basic authentication credentials"},
		{"", "", "", "testuser", "", "testbearertoken", "Cannot use Bearer token with basic authentication credentials"},
		{"", "", "", "", "testpassword", "testbearertoken", "Cannot use Bearer token with basic authentication credentials"},
		{"", "", "", "", "", "", "Bearer token or basic authentication credentials must be provided"},
		{"testuser", "testpassword", "", "", "", "", ""},
		{"", "", "testbearertoken", "", "", "", ""},
		{"", "", "", "", "", "testbearertoken", ""},
		{"", "", "", "testuser", "testpassword", "", ""},
		{"testuser", "", "", "", "", "", "Both username and password must be provided for basic authentication and cannot mix params and env vars"},
		{"", "testpassword", "", "", "", "", "Both username and password must be provided for basic authentication and cannot mix params and env vars"},
		{"testuser", "", "", "", "testpassword", "", "Both username and password must be provided for basic authentication and cannot mix params and env vars"},
		{"", "testpassword", "", "testuser", "", "", "Both username and password must be provided for basic authentication and cannot mix params and env vars"},
	}

	// Iterate over the test matrix
	for testNr, test := range matrix {
		// Set the environment variables
		os.Setenv("SOLACEBROKER_USERNAME", test.EnvUsername)
		os.Setenv("SOLACEBROKER_PASSWORD", test.EnvPassword)
		os.Setenv("SOLACEBROKER_BEARER_TOKEN", test.EnvBearertoken)

		// Check if the environment variables are set correctly
		env := os.Getenv("SOLACEBROKER_USERNAME")
		fmt.Println(env)
		// Create a providerData struct from the test matrix
		var username, password, bearertoken types.String
		if test.ParamUsername != "" {
			username = types.StringValue(test.ParamUsername)
		} else {
			username = types.StringNull()
		}
		if test.ParamPassword != "" {
			password = types.StringValue(test.ParamPassword)
		} else {
			password = types.StringNull()
		}
		if test.ParamBearertoken != "" {
			bearertoken = types.StringValue(test.ParamBearertoken)
		} else {
			bearertoken = types.StringNull()
		}
		providerData := &providerData{
			Username:    username,
			Password:    password,
			BearerToken: bearertoken,
			Url:         types.StringValue("https://example.com"),
		}
		_, diag := client(providerData)
		// Check if the actual value is equal to the expected value
		if diag != nil {
			summary := diag.Summary()
			if test.Expected != summary {
				t.Errorf("Test %d: expected %v but got %v", testNr, test.Expected, summary)
			}
		} else {
			if test.Expected != "" {
				t.Errorf("Test %d: expected %v but got nil diag", testNr, test.Expected)
			}
		}
	}
}
