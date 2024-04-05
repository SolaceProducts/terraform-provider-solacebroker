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
package client

import (
	"strings"
	"terraform-provider-solacebroker/cmd/generator"
	"terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/semp"
	"time"
)

func CliClient(url string) *semp.Client {
	username := generator.StringWithDefaultFromEnv("username", true, "")
	password := generator.StringWithDefaultFromEnv("password", false, "")
	bearerToken := generator.StringWithDefaultFromEnv("bearer_token", false, "")
	retries, err := generator.Int64WithDefaultFromEnv("retries", false, 10)
	if err != nil {
		generator.ExitWithError("\nError: Unable to parse provider attribute. " + err.Error())
	}
	retryMinInterval, err := generator.DurationWithDefaultFromEnv("retry_min_interval", false, 3*time.Second)
	if err != nil {
		generator.ExitWithError("\nError: Unable to parse provider attribute. " + err.Error())
	}
	retryMaxInterval, err := generator.DurationWithDefaultFromEnv("retry_max_interval", false, 30*time.Second)
	if err != nil {
		generator.ExitWithError("\nError: Unable to parse provider attribute. " + err.Error())
	}
	requestTimeoutDuration, err := generator.DurationWithDefaultFromEnv("request_timeout_duration", false, time.Minute)
	if err != nil {
		generator.ExitWithError("\nError: Unable to parse provider attribute. " + err.Error())
	}
	requestMinInterval, err := generator.DurationWithDefaultFromEnv("request_min_interval", false, 100*time.Millisecond)
	if err != nil {
		generator.ExitWithError("\nError: Unable to parse provider attribute. " + err.Error())
	}
	insecure_skip_verify, err := generator.BooleanWithDefaultFromEnv("insecure_skip_verify", false, false)
	if err != nil {
		generator.ExitWithError("\nError: Unable to parse provider attribute. " + err.Error())
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
