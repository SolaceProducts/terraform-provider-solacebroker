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

package generated

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "dmr_cluster_link_tls_trusted_common_name",
		MarkdownDescription: "The Trusted Common Names for the Link are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\ndmrClusterName|x||x|\nremoteNodeName|x||x|\ntlsTrustedCommonName|x||x|\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.",
		ObjectType:          broker.ReplaceOnlyObject,
		PathTemplate:        "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName}",
		PostPathTemplate:    "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "dmrClusterName",
				TerraformName:       "dmr_cluster_name",
				MarkdownDescription: "The name of the Cluster. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			{
				SempName:            "remoteNodeName",
				TerraformName:       "remote_node_name",
				MarkdownDescription: "The name of the node at the remote end of the Link. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&/]+$"), ""),
				},
			},
			{
				SempName:            "tlsTrustedCommonName",
				TerraformName:       "tls_trusted_common_name",
				MarkdownDescription: "The expected trusted common name of the remote certificate. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Deprecated:          true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 64),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
