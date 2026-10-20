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

package main

import (
	"runtime/debug"
	"strings"
)

// Set by goreleaser from the tag. Other builds fall back to Go's VCS stamp.
var version = ""

// The major sets every resource's schema version, so the fallback must never report major 0.
const devVersion = "1.0.0-dev"

func providerVersion() string {
	buildVersion := ""
	if info, ok := debug.ReadBuildInfo(); ok {
		buildVersion = info.Main.Version
	}
	return resolveVersion(version, buildVersion)
}

func resolveVersion(ldflagsVersion, buildVersion string) string {
	v := ldflagsVersion
	if v == "" {
		v = buildVersion
	}
	v = strings.TrimPrefix(v, "v")
	if v == "" || v == "(devel)" || strings.HasPrefix(v, "0.") {
		return devVersion
	}
	return v
}
