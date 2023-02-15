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

package main

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name solacebroker

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"log"
	"os"
	"terraform-provider-solacebroker/internal/broker"
	_ "terraform-provider-solacebroker/internal/broker/generated"
)

var (
// these will be set by the goreleaser configuration
// to appropriate values for the compiled binary

// commenting out for now, using version from a different file
//	version string = "dev"

// goreleaser can also pass the specific commit if you want
// commit  string = ""
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	registry, ok := os.LookupEnv("SOLACEBROKER_REGISTRY_OVERRIDE")
	if !ok {
		registry = "registry.terraform.io"
	}

	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of your provider.
		Address: registry + "/" + providerNamespace + "/" + providerType,
		Debug:   debug,
	}

	if debug {
		go debugRun(os.Getenv("SOLACEBROKER_DEBUG_RUN"), opts.Address)
	}

	err := providerserver.Serve(context.Background(), broker.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
