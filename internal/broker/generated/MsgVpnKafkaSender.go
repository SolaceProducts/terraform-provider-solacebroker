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

package generated

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_kafka_sender",
		MarkdownDescription: "A Kafka Sender sends messages to a Kafka Cluster.\n\n\nAttribute|Identifying|Write-Only|Opaque\n:---|:---:|:---:|:---:\nauthentication_basic_password||x|x\nauthentication_client_cert_content||x|x\nauthentication_client_cert_password||x|\nauthentication_oauth_client_secret||x|x\nauthentication_scram_password||x|x\nkafka_sender_name|x||\nmsg_vpn_name|x||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.36.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "authenticationBasicPassword",
				TerraformName:       "authentication_basic_password",
				MarkdownDescription: "The password for the Username. To be used when authentication_scheme is \"basic\". This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Sensitive:           true,
				Requires:            []string{"authentication_basic_username"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("authentication_basic_username"),
					),
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationBasicUsername",
				TerraformName:       "authentication_basic_username",
				MarkdownDescription: "The username the Kafka Sender uses to login to the remote Kafka broker. To be used when authentication_scheme is \"basic\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationClientCertContent",
				TerraformName:       "authentication_client_cert_content",
				MarkdownDescription: "The PEM formatted content for the client certificate used by the Kafka Sender to login to the remote Kafka broker. To be used when authentication_scheme is \"client-certificate\". Alternatively this will be used for other values of authentication_scheme when the Kafka broker has an `ssl.client.auth` setting of \"requested\" or \"required\" and KIP-684 (mTLS) is supported by the Kafka broker. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. The default value is `\"\"`.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32768),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationClientCertPassword",
				TerraformName:       "authentication_client_cert_password",
				MarkdownDescription: "The password for the client certificate. To be used when authentication_scheme is \"client-certificate\". Alternatively this will be used for other values of authentication_scheme when the Kafka broker has an `ssl.client.auth` setting of \"requested\" or \"required\" and KIP-684 (mTLS) is supported by the Kafka broker. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. The default value is `\"\"`.",
				Sensitive:           true,
				Requires:            []string{"authentication_client_cert_content"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("authentication_client_cert_content"),
					),
					stringvalidator.LengthBetween(0, 512),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthClientId",
				TerraformName:       "authentication_oauth_client_id",
				MarkdownDescription: "The OAuth client ID. To be used when authentication_scheme is \"oauth-client\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthClientScope",
				TerraformName:       "authentication_oauth_client_scope",
				MarkdownDescription: "The OAuth scope. To be used when authentication_scheme is \"oauth-client\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthClientSecret",
				TerraformName:       "authentication_oauth_client_secret",
				MarkdownDescription: "The OAuth client secret. To be used when authentication_scheme is \"oauth-client\". This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 512),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthClientTokenEndpoint",
				TerraformName:       "authentication_oauth_client_token_endpoint",
				MarkdownDescription: "The OAuth token endpoint URL that the Kafka Sender will use to request a token for login to the Kafka broker. Must begin with \"https\". To be used when authentication_scheme is \"oauth-client\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 2048),
					stringvalidator.RegexMatches(regexp.MustCompile("^([hH][tT][tT][pP][sS]://.+)?$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationScheme",
				TerraformName:       "authentication_scheme",
				MarkdownDescription: "The authentication scheme for the Kafka Sender. The bootstrap addresses must resolve to an appropriately configured and compatible listener port on the Kafka broker for the given scheme. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"none\"`. The allowed values and their meaning are:\n\n<pre>\n\"none\" - Anonymous Authentication. Used with Kafka broker PLAINTEXT listener ports.\n\"basic\" - Basic Authentication. Used with Kafka broker SASL_PLAINTEXT and SASL_SSL listener ports.\n\"scram\" - Salted Challenge Response Authentication. Used with Kafka broker SASL_PLAINTEXT and SASL_SSL listener ports.\n\"client-certificate\" - Client Certificate Authentication. Used with Kafka broker SSL listener ports.\n\"oauth-client\" - Oauth Authentication. Used with Kafka broker SASL_SSL listener ports.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("none", "basic", "scram", "client-certificate", "oauth-client"),
				},
				Default: "none",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationScramHash",
				TerraformName:       "authentication_scram_hash",
				MarkdownDescription: "The hash used for SCRAM authentication. To be used when authentication_scheme is \"scram\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"sha-512\"`. The allowed values and their meaning are:\n\n<pre>\n\"sha-256\" - SHA-2 256 bits.\n\"sha-512\" - SHA-2 512 bits.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("sha-256", "sha-512"),
				},
				Default: "sha-512",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationScramPassword",
				TerraformName:       "authentication_scram_password",
				MarkdownDescription: "The password for the Username. To be used when authentication_scheme is \"scram\". This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Sensitive:           true,
				Requires:            []string{"authentication_scram_username"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("authentication_scram_username"),
					),
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationScramUsername",
				TerraformName:       "authentication_scram_username",
				MarkdownDescription: "The username the Kafka Sender uses to login to the remote Kafka broker. To be used when authentication_scheme is \"scram\". Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "batchDelay",
				TerraformName:       "batch_delay",
				MarkdownDescription: "Delay (in ms) to wait to accumulate a batch of messages to send. Batching is done for all Senders on a per-partition basis.\n\nThis corresponds to the Kafka producer API `linger.ms` configuration setting.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `5`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(0, 900000),
				},
				Default: 5,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "batchMaxMsgCount",
				TerraformName:       "batch_max_msg_count",
				MarkdownDescription: "Maximum number of messages sent in a single batch. Batching is done for all Senders on a per-partition basis. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 1000000),
				},
				Default: 10000,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "batchMaxSize",
				TerraformName:       "batch_max_size",
				MarkdownDescription: "Maximum size of a message batch, in bytes (B). Batching is done for all Senders on a per-partition basis.\n\nThis corresponds to the Kafka producer API `batch.size` configuration setting, and should not exceed either the Kafka broker `message.max.bytes` configuration setting, or the per-Topic override of `max.message.bytes`.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 2147483647),
				},
				Default: 1e+06,
			},
			{
				BaseType:            broker.String,
				SempName:            "bootstrapAddressList",
				TerraformName:       "bootstrap_address_list",
				MarkdownDescription: "Comma separated list of addresses (and optional ports) of brokers in the Kafka Cluster from which the state of the entire Kafka Cluster can be learned. If a port is not provided with an address it will default to 9092.\n\nThis corresponds to the Kafka producer API `bootstrap.servers` configuration setting.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 1044),
					stringvalidator.RegexMatches(regexp.MustCompile("^(((((([0-9a-zA-Z\\-\\.])+)|\\[([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,7}:\\]|\\[([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}\\]|\\[([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}\\]|\\[([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}\\]|\\[[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})\\]|\\[:((:[0-9a-fA-F]{1,4}){1,7}|:)\\])((:[0-9]{1,5}){0,1})),)*(((([0-9a-zA-Z\\-\\.])+)|\\[([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,7}:\\]|\\[([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}\\]|\\[([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}\\]|\\[([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}\\]|\\[([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}\\]|\\[[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})\\]|\\[:((:[0-9a-fA-F]{1,4}){1,7}|:)\\])((:[0-9]{1,5}){0,1})))?$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable the Kafka Sender. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "idempotenceEnabled",
				TerraformName:       "idempotence_enabled",
				MarkdownDescription: "Enable or disable idempotence for the Kafka Sender. Idempotence guarantees in order at-least-once message delivery to the remote Kafka Topic, at the expense of performance. When idempotence is enabled the Queue Bindings of the Kafka Sender must have ack_mode of \"all\" to be operational.\n\nThis corresponds to the Kafka producer API `enable.idempotence` configuration setting.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "kafkaSenderName",
				TerraformName:       "kafka_sender_name",
				MarkdownDescription: "The name of the Kafka Sender.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^#*? ]([^*?]*[^*? ])?$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "msgVpnName",
				TerraformName:       "msg_vpn_name",
				MarkdownDescription: "The name of the Message VPN.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?]+$"), ""),
				},
			},
			{
				BaseType:            broker.Bool,
				SempName:            "transportCompressionEnabled",
				TerraformName:       "transport_compression_enabled",
				MarkdownDescription: "Enable or disable compression for the Kafka Sender. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.Int64,
				SempName:            "transportCompressionLevel",
				TerraformName:       "transport_compression_level",
				MarkdownDescription: "Compression level. The valid range is dependent on the compression type.\n\nThis corresponds to the Kafka producer API `compression.level` configuration setting.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `-1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(-1, 22),
				},
				Default: -1,
			},
			{
				BaseType:            broker.String,
				SempName:            "transportCompressionType",
				TerraformName:       "transport_compression_type",
				MarkdownDescription: "Compression type. Only relevant if compression is enabled.\n\nThis corresponds to the Kafka producer API `compression.type` configuration setting.\n\nModifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"gzip\"`. The allowed values and their meaning are:\n\n<pre>\n\"gzip\" - GZIP Compression.\n\"snappy\" - Snappy Compression.\n\"lz4\" - LZ4 Compression.\n\"zstd\" - Zstandard Compression.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("gzip", "snappy", "lz4", "zstd"),
				},
				Default: "gzip",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "transportTlsEnabled",
				TerraformName:       "transport_tls_enabled",
				MarkdownDescription: "Enable or disable encryption (TLS) for the Kafka Sender. The bootstrap addresses must resolve to PLAINTEXT or SASL_PLAINTEXT listener ports when disabled, and SSL or SASL_SSL listener ports when enabled. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
