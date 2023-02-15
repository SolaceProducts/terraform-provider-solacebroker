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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithMetadata = &BrokerProvider{}

type BrokerProvider struct {
	Version string
}

func (p *BrokerProvider) Metadata(_ context.Context, _ provider.MetadataRequest, response *provider.MetadataResponse) {
	response.Version = p.Version
	response.TypeName = "solacebroker"
}

func (p *BrokerProvider) GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"url": {
				Type:     types.StringType,
				Optional: true,
			},
			"username": {
				Type:     types.StringType,
				Optional: true,
			},
			"password": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
			"bearer_token": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
			"retries": {
				Type:     types.Int64Type,
				Optional: true,
			},
			"retry_wait": {
				Type:     types.Int64Type,
				Optional: true,
			},
			"retry_wait_max": {
				Type:     types.Int64Type,
				Optional: true,
			},
		},
	}, nil
}

func (p *BrokerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.ResourceData = &config
	resp.DataSourceData = &config
}

func (p *BrokerProvider) Resources(context.Context) []func() resource.Resource {
	return Resources
}

func (p *BrokerProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return DataSources
}

type providerData struct {
	Url          types.String `tfsdk:"url"`
	Username     types.String `tfsdk:"username"`
	Password     types.String `tfsdk:"password"`
	BearerToken  types.String `tfsdk:"bearer_token"`
	Retries      types.Int64  `tfsdk:"retries"`
	RetryWait    types.Int64  `tfsdk:"retry_wait"`
	RetryWaitMax types.Int64  `tfsdk:"retry_wait_max"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &BrokerProvider{
			Version: version,
		}
	}
}
