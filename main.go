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

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"terraform-provider-solacebroker/cmd"
	"terraform-provider-solacebroker/internal/broker"
	_ "terraform-provider-solacebroker/internal/broker/generated"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
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
	if broker.SempDetail.Platform != expectedPlatform {
		fmt.Println(fmt.Sprintf("Provider error: wrong platform SEMP API spec \"%s\" used, expected \"%s\"", broker.SempDetail.Platform, expectedPlatform))
		os.Exit(1)
	}
	broker.ProviderVersion = version
	if len(os.Args) > 1 && (os.Args[1] == "generate" || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "version") {
		err := cmd.Execute()
		if err != nil && err.Error() != "" {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		var debug bool
		flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
		flag.Parse()
		registry, ok := os.LookupEnv("SOLACEBROKER_REGISTRY_OVERRIDE")
		if !ok {
			registry = "registry.terraform.io"
		}
		opts := providerserver.ServeOpts{
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
}
