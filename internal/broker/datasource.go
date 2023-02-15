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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"net/http"
	"os"
	"terraform-provider-solacebroker/internal/semp"
)

func newBrokerDataSourceGenerator(inputs EntityInputs) func() datasource.DataSource {
	return newBrokerDataSourceClosure(newBrokerEntity(inputs, false))
}

func newBrokerDataSourceClosure(templateEntity brokerEntity) func() datasource.DataSource {
	return func() datasource.DataSource {
		var ds = brokerDataSource(templateEntity)
		return &ds
	}
}

var (
	_ datasource.DataSourceWithConfigure        = &brokerDataSource{}
	_ datasource.DataSourceWithConfigValidators = &brokerDataSource{}
)

type brokerDataSource brokerEntity

func (ds *brokerDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return ds.schema, nil
}

func (ds *brokerDataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_" + ds.terraformName
}

func (ds *brokerDataSource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	config, ok := request.ProviderData.(*providerData)
	if !ok {
		response.Diagnostics = diag.Diagnostics{diag.NewErrorDiagnostic("Unexpected resource configuration", fmt.Sprintf("Unexpected type %T for provider data; expected %T.", request.ProviderData, config))}
		return
	}
	ds.providerData = config
}

func (ds *brokerDataSource) client() (*semp.Client, diag.Diagnostic) {

	// User must provide a user to the provider
	var username string
	if ds.providerData.Username.Unknown {
		// Cannot connect to client with an unknown value
		return nil, diag.NewErrorDiagnostic(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
	}

	if ds.providerData.Username.Null {
		username = os.Getenv("SOLACE_BROKER_USERNAME")
	} else {
		username = ds.providerData.Username.Value
	}

	if username == "" {
		return nil, diag.NewErrorDiagnostic(
			"Unable to find username",
			"Username cannot be an empty string",
		)
	}

	// User must provide a password to the provider
	var password string
	if ds.providerData.Password.Unknown {
		return nil, diag.NewErrorDiagnostic(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
	}

	if ds.providerData.Password.Null {
		password = os.Getenv("SOLACE_BROKER_PASSWORD")
	} else {
		password = ds.providerData.Password.Value
	}

	if password == "" {
		return nil, diag.NewErrorDiagnostic(
			"Unable to find password",
			"password cannot be an empty string",
		)
	}

	// User must specify an url
	var url string
	if ds.providerData.Url.Unknown {
		return nil, diag.NewErrorDiagnostic(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
	}

	if ds.providerData.Url.Null {
		url = os.Getenv("SOLACE_BROKER_URL")
	} else {
		url = ds.providerData.Url.Value
	}

	if url == "" {
		return nil, diag.NewErrorDiagnostic(
			"Unable to find url",
			"Url cannot be an empty string",
		)
	}

	return semp.NewClient(url, semp.BasicAuth(username, password)), nil
}

func (ds *brokerDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	client, d := ds.client()
	if d != nil {
		response.Diagnostics.Append(d)
		if response.Diagnostics.HasError() {
			return
		}
	}

	path, err := resolveSempPath(ds.pathTemplate, ds.identifyingAttributes, request.Config.Raw)
	if err != nil {
		response.Diagnostics = generateDiagnostics("Error generating SEMP path", err)
		return
	}
	sempData, err := client.RequestWithoutBody(http.MethodGet, path)
	if err != nil {
		response.Diagnostics = generateDiagnostics("SEMP call failed", err)
		return
	}

	responseData, err := ds.converter.ToTerraform(sempData)
	if err != nil {
		response.Diagnostics = generateDiagnostics("SEMP response conversion failed", err)
		return
	}

	response.State.Raw = responseData
}

func (ds *brokerDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return nil
}
