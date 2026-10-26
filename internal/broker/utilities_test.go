package broker

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
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
		{"", "", "testbearertoken", "", "", "testbearertoken", ""},
		{"testuser", "testpassword", "", "", "", "testbearertoken", ""},
		{"testuser", "testpassword", "", "testuser", "testpassword", "testbearertoken", ""},
		{"", "", "testbearertoken", "testuser", "testpassword", "", ""},
		{"", "", "testbearertoken", "testuser", "testpassword", "testbearertoken", ""},
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

func TestClientURL(t *testing.T) {
	// Each server records that it was called, so the test sees which URL client() used
	var paramHit, envHit bool
	paramServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { paramHit = true }))
	defer paramServer.Close()
	envServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { envHit = true }))
	defer envServer.Close()

	matrix := []struct {
		ParamUrl string
		EnvUrl   string
		Expected string
	}{
		{paramServer.URL, "", ""},
		{"", envServer.URL, ""},
		{paramServer.URL, envServer.URL, ""},
		{"", "", "Broker URL must be provided in the provider block (url) or the SOLACEBROKER_URL environment variable"},
	}
	for testNr, test := range matrix {
		paramHit, envHit = false, false
		t.Setenv("SOLACEBROKER_URL", test.EnvUrl)
		url := types.StringNull()
		if test.ParamUrl != "" {
			url = types.StringValue(test.ParamUrl)
		}
		client, diag := client(&providerData{
			Username: types.StringValue("testuser"),
			Password: types.StringValue("testpassword"),
			Url:      url,
			Retries:  types.Int64Value(0),
		})
		summary := ""
		if diag != nil {
			summary = diag.Summary()
		}
		if summary != test.Expected {
			t.Errorf("Test %d: expected %q but got %q", testNr, test.Expected, summary)
		}
		if client == nil {
			continue
		}
		_, _ = client.RequestWithoutBody(context.Background(), http.MethodGet, "/about/api")
		// The provider block takes precedence over the environment variable
		wantParam := test.ParamUrl != ""
		if paramHit != wantParam || envHit == wantParam {
			t.Errorf("Test %d: expected the %s URL to be used, got param=%v env=%v", testNr, map[bool]string{true: "param", false: "env"}[wantParam], paramHit, envHit)
		}
	}
}

func TestProviderURLOptional(t *testing.T) {
	resp := &provider.SchemaResponse{}
	New("test")().Schema(context.Background(), provider.SchemaRequest{}, resp)
	url := resp.Schema.Attributes["url"]
	if !url.IsOptional() || url.IsRequired() {
		t.Errorf("url must be optional so SOLACEBROKER_URL can set it")
	}
}
