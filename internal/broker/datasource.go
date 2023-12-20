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
	"errors"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"terraform-provider-solacebroker/internal/semp"
)

func newBrokerDataSource(inputs EntityInputs) brokerEntity[schema.Schema] {
	return resourceEntityToDataSourceEntity(newBrokerEntity(inputs, false))
}

func newBrokerDataSourceGenerator(inputs EntityInputs) func() datasource.DataSource {
	return newBrokerDataSourceClosure(newBrokerDataSource(inputs))
}

func newBrokerDataSourceClosure(templateEntity brokerEntity[schema.Schema]) func() datasource.DataSource {
	return func() datasource.DataSource {
		var ds = brokerDataSource(templateEntity)
		return &ds
	}
}

var (
	_ datasource.DataSourceWithConfigure        = &brokerDataSource{}
	_ datasource.DataSourceWithConfigValidators = &brokerDataSource{}
)

type brokerDataSource brokerEntity[schema.Schema]

func (ds *brokerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = ds.schema
}

func (ds *brokerDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_" + ds.terraformName
}

func (ds *brokerDataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	client, ok := request.ProviderData.(*semp.Client)
	if !ok {
		response.Diagnostics.AddError(
			"Unexpected datasource configuration",
			fmt.Sprintf("Unexpected type %T for provider data; expected %T.", request.ProviderData, client),
		)
		return
	}
	ds.client = client
}

func (ds *brokerDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	client := ds.client
	if err := checkBrokerRequirements(ctx, client); err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Broker check failed", err)
		return
	}
	sempPath, err := resolveSempPath(ds.pathTemplate, ds.identifyingAttributes, request.Config.Raw)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "Error generating SEMP path", err)
		return
	}
	sempData, err := client.RequestWithoutBody(ctx, http.MethodGet, sempPath)
	if err != nil {
		if errors.Is(err, semp.ErrResourceNotFound) {
			addErrorToDiagnostics(&response.Diagnostics, fmt.Sprintf("Detected missing data source %v", sempPath), errors.Unwrap(err))
		} else {
			addErrorToDiagnostics(&response.Diagnostics, "SEMP call failed", err)
		}
		return
	}
	responseData, err := ds.converter.ToTerraform(sempData)
	if err != nil {
		addErrorToDiagnostics(&response.Diagnostics, "SEMP response conversion failed", err)
		return
	}

	response.State.Raw = responseData
}

func (ds *brokerDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return nil
}
