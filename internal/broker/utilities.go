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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"net/url"
	"os"
	"strconv"
	"strings"
	"terraform-provider-solacebroker/internal/semp"
	"time"
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
		// Cannot connect to client with an unknown value
		return "", fmt.Errorf("Cannot use unknown value as %v", name)
	}

	var s string
	if value.IsNull() {
		s = os.Getenv("SOLACEBROKER_" + strings.ToUpper(name))
	} else {
		s = value.ValueString()
	}

	return s, nil
}

func int64WithDefaultFromEnv(value types.Int64, name string, def int64) (int64, error) {
	if value.IsUnknown() {
		// Cannot connect to client with an unknown value
		return 0, fmt.Errorf("Cannot use unknown value as %v", name)
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

func durationWithDefaultFromEnv(value types.String, name string, def time.Duration) (time.Duration, error) {
	if value.IsUnknown() {
		// Cannot connect to client with an unknown value
		return 0, fmt.Errorf("Cannot use unknown value as %v", name)
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

	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, fmt.Errorf("%v is not valid; %q cannot be parsed as a duration: %w", name, s, err)
	}

	return d, nil
}

func client(providerData *providerData) (*semp.Client, diag.Diagnostic) {
	username, err := stringWithDefaultFromEnv(providerData.Username, "username")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	password, err := stringWithDefaultFromEnv(providerData.Password, "password")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	bearerToken, err := stringWithDefaultFromEnv(providerData.BearerToken, "bearer_token")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	url, err := stringWithDefaultFromEnv(providerData.Url, "url")
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	retries, err := int64WithDefaultFromEnv(providerData.Retries, "retries", 10)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	retryMinInterval, err := durationWithDefaultFromEnv(providerData.RetryMinInterval, "retry_min_interval", 3*time.Second)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	retryMaxInterval, err := durationWithDefaultFromEnv(providerData.RetryMaxInterval, "retry_max_interval", 30*time.Second)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	requestTimeoutDuration, err := durationWithDefaultFromEnv(providerData.RequestTimeoutDuration, "request_timeout_duration", time.Minute)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}
	requestMinInterval, err := durationWithDefaultFromEnv(providerData.RequestMinInterval, "request_min_interval", 100*time.Millisecond)
	if err != nil {
		return nil, diag.NewErrorDiagnostic("Unable to create client", err.Error())
	}

	client := semp.NewClient(
		url,
		semp.BasicAuth(username, password),
		semp.BearerToken(bearerToken),
		semp.Retries(uint(retries), retryMinInterval, retryMaxInterval),
		semp.RequestLimits(requestTimeoutDuration, requestMinInterval))
	return client, nil
}
