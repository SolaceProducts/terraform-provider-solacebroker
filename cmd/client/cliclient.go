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
)

func CliClient(cliParams generator.CliParams) *semp.Client {
	client := semp.NewClient(
		getFullSempAPIURL(*cliParams.Url),
		*cliParams.Insecure_skip_verify,
		false, // this is a client for the generator
		semp.BasicAuth(*cliParams.Username, *cliParams.Password),
		semp.BearerToken(*cliParams.Bearer_token),
		semp.Retries(uint(*cliParams.Retries), *cliParams.Retry_min_interval, *cliParams.Retry_max_interval),
		semp.RequestLimits(*cliParams.Request_timeout_duration, *cliParams.Request_min_interval))
	return client
}

func getFullSempAPIURL(url string) string {
	url = strings.TrimSuffix(url, "/")
	baseBath := strings.TrimPrefix(broker.SempDetail.BasePath, "/")
	return url + "/" + baseBath
}
