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
package broker

import (
	"os"
	"strings"
	"terraform-provider-solacebroker/cmd/command"
	"terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/semp"
	"time"
)

func CliClient(url string) *semp.Client {
	username := terraform.StringWithDefaultFromEnv("username", true, "")
	password := terraform.StringWithDefaultFromEnv("password", false, "")
	bearerToken := terraform.StringWithDefaultFromEnv("bearer_token", false, "")
	retries, err := terraform.Int64WithDefaultFromEnv("retries", false, 10)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	retryMinInterval, err := terraform.DurationWithDefaultFromEnv("retry_min_interval", false, 3*time.Second)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	retryMaxInterval, err := terraform.DurationWithDefaultFromEnv("retry_max_interval", false, 30*time.Second)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	requestTimeoutDuration, err := terraform.DurationWithDefaultFromEnv("request_timeout_duration", false, time.Minute)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	requestMinInterval, err := terraform.DurationWithDefaultFromEnv("request_min_interval", false, 100*time.Millisecond)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	insecure_skip_verify, err := terraform.BooleanWithDefaultFromEnv("insecure_skip_verify", false, false)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	client := semp.NewClient(
		getFullSempAPIURL(url),
		insecure_skip_verify,
		false, // this is a client for the generator
		semp.BasicAuth(username, password),
		semp.BearerToken(bearerToken),
		semp.Retries(uint(retries), retryMinInterval, retryMaxInterval),
		semp.RequestLimits(requestTimeoutDuration, requestMinInterval))
	return client
}

func getFullSempAPIURL(url string) string {
	url = strings.TrimSuffix(url, "/")
	baseBath := strings.TrimPrefix(broker.SempDetail.BasePath, "/")
	return url + "/" + baseBath
}
