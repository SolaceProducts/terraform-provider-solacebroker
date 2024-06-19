// terraform-provider-solacebroker
//
// Copyright 2024 Solace Corporation. All rights reserved.
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
	"net/url"
	"os"
	"strconv"
	"strings"
	"terraform-provider-solacebroker/internal/semp"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func resolveSempPath(pathTemplate string, attributes []*AttributeInfo, v tftypes.Value) (string, error) {
	m := map[string]tftypes.Value{}
	err := v.As(&m)
	if err != nil {
		return "", err
	}
	identifiers := map[string]string{}
	for _, attr := range attributes {
		v, err := attr.Converter.FromTerraform(m[attr.TerraformName])
		if err != nil {
			return "", err
		}
		identifiers[attr.SempName] = fmt.Sprintf("%v", v)
	}
	// doing it this way identifies missed parameters (as opposed to doing strings.Replace or something like that)
	var path string
	split := strings.SplitN(pathTemplate, "{", 2)
	for len(split) == 2 {
		path += split[0]
		split = strings.SplitN(split[1], "}", 2)
		v, ok := identifiers[split[0]]
		if !ok {
			return "", fmt.Errorf("no value provided for SEMP path parameter %v", split[0])
		}
		path += url.PathEscape(v)
		split = strings.SplitN(split[1], "{", 2)
	}
	return path + split[0], nil
}

func stringWithDefaultFromEnv(value types.String, name string) (string, error) {
	if value.IsUnknown() {
		return "", fmt.Errorf("cannot use unknown value as %v", name)
	}

	var s string
	if value.IsNull() {
		// If env var is not found then the default return value will be empty string
		s = os.Getenv("SOLACEBROKER_" + strings.ToUpper(name))
	} else {
		s = value.ValueString()
	}

	return s, nil
}

func int64WithDefaultFromEnv(value types.Int64, name string, def int64) (int64, error) {
	if value.IsUnknown() {
		return 0, fmt.Errorf("cannot use unknown value as %v", name)
	}

	if !value.IsNull() {
		return value.ValueInt64(), nil
	}

	envName := "SOLACEBROKER_" + strings.ToUpper(name)
	s, ok := os.LookupEnv(envName)
	if !ok {
		return def, nil
	}
	return strconv.ParseInt(s, 10, 64)
}

func booleanWithDefaultFromEnv(value types.Bool, name string, def bool) (bool, error) {
	if value.IsUnknown() {
		return false, fmt.Errorf("cannot use unknown value as %v", name)
	}

	if !value.IsNull() {
		return value.ValueBool(), nil
	}

	envName := "SOLACEBROKER_" + strings.ToUpper(name)
	s, ok := os.LookupEnv(envName)
	if !ok {
		return def, nil
	}
	return strconv.ParseBool(s)
}

func durationWithDefaultFromEnv(value types.String, name string, def time.Duration) (time.Duration, error) {
	if value.IsUnknown() {
		return 0, fmt.Errorf("cannot use unknown value as %v", name)
	}

	var s string
	if value.IsNull() {
		s = os.Getenv("SOLACEBROKER_" + strings.ToUpper(name))
	} else {
		s = value.ValueString()
	}

	if s == "" {
		return def, nil
	}
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, fmt.Errorf("%v is not valid; %q cannot be parsed as a duration: %w", name, s, err)
	}

	return d, nil
}

func client(providerData *providerData) (*semp.Client, diag.Diagnostic) {
	// username, password and bearer token will be set to "" if not provided through env or config
	username, err := stringWithDefaultFromEnv(providerData.Username, "username")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	password, err := stringWithDefaultFromEnv(providerData.Password, "password")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	bearerToken, err := stringWithDefaultFromEnv(providerData.BearerToken, "bearer_token")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	// Check for params credentials conflicts
	// Logic:
	// If there is any 1 complete set of credentials in the provider block those are always used and are the priority.
	// If there is not any 1 complete set of credentials in the provider block then look for 1 complete set in the env vars.
	// If there are multiple complete sets in either the provider block or env vars this is an error.
	// If there are no complete sets in the env vars this is an error. (edited)
	if username == "" && password == "" && bearerToken == "" {
		return nil, diag.NewErrorDiagnostic("Bearer token or basic authentication credentials must be provided", semp.ErrProviderParametersError.Error())
	}
	if !providerData.BearerToken.IsNull() && providerData.Username.IsNull() && providerData.Password.IsNull() ||
		providerData.BearerToken.IsNull() && !providerData.Username.IsNull() && !providerData.Password.IsNull() {
		// these are valid combinations in the provider block, no need to check further
	} else {
		if (!providerData.BearerToken.IsNull() && (!providerData.Username.IsNull() || !providerData.Password.IsNull())) ||
			(bearerToken != "" && (username != "" || password != "")) {
			return nil, diag.NewErrorDiagnostic("Cannot use Bearer token with basic authentication credentials", semp.ErrProviderParametersError.Error())
		}
		if !providerData.Username.IsNull() && providerData.Password.IsNull() || providerData.Username.IsNull() && !providerData.Password.IsNull() ||
			username != "" && password == "" || username == "" && password != "" {
			return nil, diag.NewErrorDiagnostic("Both username and password must be provided for basic authentication and cannot mix params and env vars", semp.ErrProviderParametersError.Error())
		}
	}
	url, err := stringWithDefaultFromEnv(providerData.Url, "url")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	retries, err := int64WithDefaultFromEnv(providerData.Retries, "retries", semp.DefaultRetries)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	retryMinInterval, err := durationWithDefaultFromEnv(providerData.RetryMinInterval, "retry_min_interval", semp.DefaultRetryMinInterval)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	retryMaxInterval, err := durationWithDefaultFromEnv(providerData.RetryMaxInterval, "retry_max_interval", semp.DefaultRetryMaxInterval)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	requestTimeoutDuration, err := durationWithDefaultFromEnv(providerData.RequestTimeoutDuration, "request_timeout_duration", semp.DefaultRequestTimeout)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	requestMinInterval, err := durationWithDefaultFromEnv(providerData.RequestMinInterval, "request_min_interval", semp.DefaultRequestInterval)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	insecureSkipVerify, err := booleanWithDefaultFromEnv(providerData.InsecureSkipVerify, "insecure_skip_verify", false)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	url = getFullSempAPIURL(url)
	skipApiCheck, err = booleanWithDefaultFromEnv(providerData.SkipApiCheck, "skip_api_check", false) // This variable is used in resource
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to parse provider attribute", err.Error())
	}
	client := semp.NewClient(
		url,
		insecureSkipVerify,
		true, // this is a client for the provider
		semp.BasicAuth(username, password),
		semp.BearerToken(bearerToken),
		semp.Retries(retries, retryMinInterval, retryMaxInterval),
		semp.RequestLimits(requestTimeoutDuration, requestMinInterval))
	return client, nil
}

func getFullSempAPIURL(url string) string {
	url = strings.TrimSuffix(url, "/")
	baseBath := strings.TrimPrefix(SempDetail.BasePath, "/")
	return url + "/" + baseBath
}

func getProviderMajorVersion(semverVersion string) int64 {
	parts := strings.Split(semverVersion, ".")
	if len(parts) == 0 {
		return 0
	}
	majorVersion, _ := strconv.ParseInt(parts[0], 10, 64)
	return majorVersion
}
