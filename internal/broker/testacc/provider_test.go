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

package acctest

import (
	"os"
	"terraform-provider-solacebroker/internal/broker"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-units"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"terraform-provider-solacebroker/internal/broker/generated"
)

var ProviderConfig string

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"solacebroker": providerserver.NewProtocol6WithError(broker.New("test")()),
	}
)

func init() {
	// start docker test broker
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "solace/solace-pubsub-standard:latest",
		ExposedPorts: []string{"8080/tcp"},
		Env: map[string]string{
			"username_admin_globalaccesslevel":  "admin",
			"username_admin_password":           "admin",
			"system_scaling_maxconnectioncount": "100",
		},
		Mounts: testcontainers.ContainerMounts{
			{
				Source: testcontainers.GenericVolumeMountSource{
					Name: "test-volume",
				},
				Target: "/var/lib/solace",
			},
		},
		HostConfigModifier: func(hostConfig *container.HostConfig) {
			hostConfig.ShmSize = 1000000000
			hostConfig.Ulimits = []*units.Ulimit{
				{
					Name: "nofile",
					Hard: 1048576,
					Soft: 1048576,
				},
			}
		},
		WaitingFor: wait.ForHTTP("/").WithPort("8080/tcp"),
	}
	solaceC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	endpoint, err := solaceC.Endpoint(ctx, "")
	if err != nil {
		panic(err)
	}
	ProviderConfig = `
provider "solacebroker" {
username = "admin"
password = "admin"
url      = "http://` + endpoint + `"
}
`
	const user = "admin"
	const password = "admin"

	if err = os.Setenv("SOLACEBROKER_URL", "http://"+endpoint); err != nil {
		panic(err)
	}

	if err = os.Setenv("SOLACEBROKER_USERNAME", user); err != nil {
		panic(err)
	}

	if err = os.Setenv("SOLACEBROKER_PASSWORD", password); err != nil {
		panic(err)
	}

	if generated.Platform == "Appliance" {
		if err = os.Setenv("SOLACEBROKER_SKIP_API_CHECK", "true"); err != nil {
			panic(err)
		}
	}
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}
