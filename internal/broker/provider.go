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
	"context"
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net/http"
)

const minRequiredBrokerSempApiVersion = "2.33" // Shipped with broker version 10.3

var _ provider.Provider = &BrokerProvider{}

type BrokerProvider struct {
	Version string
}

func (p *BrokerProvider) Metadata(_ context.Context, _ provider.MetadataRequest, response *provider.MetadataResponse) {
	response.Version = p.Version
	response.TypeName = "solacebroker"
}

func (p *BrokerProvider) Schema(_ context.Context, _ provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				MarkdownDescription: "The base URL of the broker, for example `https://mybroker.example.org:1943/`.",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "The username for the broker request.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"bearer_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"retries": schema.Int64Attribute{
				Optional: true,
			},
			"retry_min_interval": schema.StringAttribute{
				Optional: true,
			},
			"retry_max_interval": schema.StringAttribute{
				Optional: true,
			},
			"request_timeout_duration": schema.StringAttribute{
				Optional: true,
			},
			"request_min_interval": schema.StringAttribute{
				Optional: true,
			},
		},
		MarkdownDescription: "",
	}
}

func (p *BrokerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, d := client(&config)
	if d != nil {
		resp.Diagnostics.Append(d)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	path := "/about/api"
	result, err := client.RequestWithoutBody(ctx, http.MethodGet, path)
	if err != nil {
		resp.Diagnostics = generateDiagnostics("SEMP call failed", err)
		return
	}
	brokerSempVersion, err := version.NewVersion(result["sempVersion"].(string))
	if err != nil {
		resp.Diagnostics = generateDiagnostics("Unable to parse SEMP version returned from \"/about/api\"", err)
		return
	}
	minSempVersion, _ := version.NewVersion(minRequiredBrokerSempApiVersion)
	if brokerSempVersion.LessThan(minSempVersion) {
		err := fmt.Errorf("BrokerSempVersion %s is less than required %s", brokerSempVersion, minSempVersion)
		resp.Diagnostics = generateDiagnostics("Broker does not meet minimum SEMP API version", err)
		return
	}

	resp.ResourceData = &config
	resp.DataSourceData = &config
}

func (p *BrokerProvider) SempVersionDetails() SempVersionDetail {
	return SempDetail
}

func (p *BrokerProvider) Resources(_ context.Context) []func() resource.Resource {
	return Resources
}

func (p *BrokerProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return DataSources
}

type providerData struct {
	Url                    types.String `tfsdk:"url"`
	Username               types.String `tfsdk:"username"`
	Password               types.String `tfsdk:"password"`
	BearerToken            types.String `tfsdk:"bearer_token"`
	Retries                types.Int64  `tfsdk:"retries"`
	RetryMinInterval       types.String `tfsdk:"retry_min_interval"`
	RetryMaxInterval       types.String `tfsdk:"retry_max_interval"`
	RequestTimeoutDuration types.String `tfsdk:"request_timeout_duration"`
	RequestMinInterval     types.String `tfsdk:"request_min_interval"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &BrokerProvider{
			Version: version,
		}
	}
}
