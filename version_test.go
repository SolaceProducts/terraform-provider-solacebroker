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

import "testing"

func TestResolveVersion(t *testing.T) {
	tests := []struct {
		name, ldflags, build, want string
	}{
		{"goreleaser", "1.4.0", "v1.3.2-0.20261019120000-abcdef123456", "1.4.0"},
		{"pseudo-version", "", "v1.3.2-0.20261019120000-abcdef123456+dirty", "1.3.2-0.20261019120000-abcdef123456+dirty"},
		{"devel", "", "(devel)", devVersion},
		{"no tags", "", "v0.0.0-20261019120000-abcdef123456", devVersion},
		{"empty", "", "", devVersion},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveVersion(tt.ldflags, tt.build); got != tt.want {
				t.Errorf("resolveVersion(%q, %q) = %q, want %q", tt.ldflags, tt.build, got, tt.want)
			}
		})
	}
}
