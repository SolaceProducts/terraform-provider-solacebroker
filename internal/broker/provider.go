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
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const minRequiredBrokerSempApiVersion = "2.33" // Shipped with broker version 10.3

var _ provider.Provider = &BrokerProvider{}
var ProviderVersion string

var Cookiejar, _ = cookiejar.New(nil)

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
				MarkdownDescription: "The base URL of the event broker, for example `https://mybroker.example.org:1943/`. The trailing / can be omitted.",
				Required:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "The username to connect to the broker with.  Requires password and conflicts with bearer_token.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The password to connect to the broker with. Requires username and conflicts with bearer_token.",
				Optional:            true,
				Sensitive:           true,
			},
			"bearer_token": schema.StringAttribute{
				MarkdownDescription: "A bearer token that will be sent in the Authorization header of SEMP requests. Conflicts with username and password.",
				Optional:            true,
				Sensitive:           true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "The number of retries for a SEMP call. The default value is 10.",
				Optional:            true,
			},
			"retry_min_interval": schema.StringAttribute{
				MarkdownDescription: "A [duration](https://pkg.go.dev/maze.io/x/duration#ParseDuration) string indicating how long to wait after an initial failed request before the first retry.  Exponential backoff is used, up to the limit set by retry_max_interval. The default value is 3s.",
				Optional:            true,
			},
			"retry_max_interval": schema.StringAttribute{
				MarkdownDescription: "A [duration](https://pkg.go.dev/maze.io/x/duration#ParseDuration) string indicating the maximum retry interval. The default value is 30s.",
				Optional:            true,
			},
			"request_timeout_duration": schema.StringAttribute{
				MarkdownDescription: "A [duration](https://pkg.go.dev/maze.io/x/duration#ParseDuration) string indicating the maximum time to wait for a SEMP request.  The default value is 1m.",
				Optional:            true,
			},
			"request_min_interval": schema.StringAttribute{
				MarkdownDescription: "A [duration](https://pkg.go.dev/maze.io/x/duration#ParseDuration) string indicating the minimum interval between requests; this serves as a rate limit. This setting does not apply to retries. Set to 0 for no rate limit. The default value is 100ms (which equates to a rate limit of 10 calls per second).",
				Optional:            true,
			},
			"insecure_skip_verify": schema.BoolAttribute{
				MarkdownDescription: "Disable validation of server SSL certificates, accept/ignore self-signed. The default value is false.",
				Optional:            true,
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

	ctx = tflog.SetField(ctx, "solacebroker_url", strings.Trim(config.Url.String(), "\""))
	ctx = tflog.SetField(ctx, "solacebroker_provider_version", p.Version)
	tflog.Debug(ctx, "Configuring solacebroker provider client")

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
		addErrorToDiagnostics(&resp.Diagnostics, "SEMP call failed", err)
		return
	}
	brokerSempVersion, err := version.NewVersion(result["sempVersion"].(string))
	if err != nil {
		addErrorToDiagnostics(&resp.Diagnostics, "unable to parse SEMP version returned from \"/about/api\"", err)
		return
	}
	minSempVersion, _ := version.NewVersion(minRequiredBrokerSempApiVersion)
	if brokerSempVersion.LessThan(minSempVersion) {
		err := fmt.Errorf("BrokerSempVersion %s is less than required %s", brokerSempVersion, minSempVersion)
		addErrorToDiagnostics(&resp.Diagnostics, "Broker does not meet minimum SEMP API version", err)
		return
	}

	tflog.Info(ctx, "Solacebroker provider client config success")

	resp.ResourceData = &config
	resp.DataSourceData = &config
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
	InsecureSkipVerify     types.Bool   `tfsdk:"insecure_skip_verify"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &BrokerProvider{
			Version: version,
		}
	}
}
