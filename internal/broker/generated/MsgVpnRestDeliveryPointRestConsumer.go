// terraform-provider-solacebroker
//
// Copyright 2025 Solace Corporation. All rights reserved.
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
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
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
		TerraformName:       "msg_vpn_rest_delivery_point_rest_consumer",
		MarkdownDescription: "REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.0.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}",
		Version:             0, // Placeholder: value will be replaced in the provider code
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "authenticationAwsAccessKeyId",
				TerraformName:       "authentication_aws_access_key_id",
				MarkdownDescription: "The AWS access key id. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.26.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationAwsRegion",
				TerraformName:       "authentication_aws_region",
				MarkdownDescription: "The AWS region id. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.26.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 20),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationAwsSecretAccessKey",
				TerraformName:       "authentication_aws_secret_access_key",
				MarkdownDescription: "The AWS secret access key. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.26.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 64),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationAwsService",
				TerraformName:       "authentication_aws_service",
				MarkdownDescription: "The AWS service id. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.26.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 50),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationClientCertContent",
				TerraformName:       "authentication_client_cert_content",
				MarkdownDescription: "The PEM formatted content for the client certificate that the REST Consumer will present to the REST host. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. The default value is `\"\"`. Available since SEMP API version 2.9.",
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
				MarkdownDescription: "The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. The default value is `\"\"`. Available since SEMP API version 2.9.",
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
				SempName:            "authenticationHttpBasicPassword",
				TerraformName:       "authentication_http_basic_password",
				MarkdownDescription: "The password for the username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Sensitive:           true,
				Requires:            []string{"authentication_http_basic_username"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("authentication_http_basic_username"),
					),
					stringvalidator.LengthBetween(0, 128),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationHttpBasicUsername",
				TerraformName:       "authentication_http_basic_username",
				MarkdownDescription: "The username that the REST Consumer will use to login to the REST host. Normally a username is only configured when basic authentication is selected for the REST Consumer. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Requires:            []string{"authentication_http_basic_password"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("authentication_http_basic_password"),
					),
					stringvalidator.LengthBetween(0, 189),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationHttpHeaderName",
				TerraformName:       "authentication_http_header_name",
				MarkdownDescription: "The authentication header name. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.15.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 50),
					stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9!#$%&'*+\\-.\\^_`|~]*$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationHttpHeaderValue",
				TerraformName:       "authentication_http_header_value",
				MarkdownDescription: "The authentication header value. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.15.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 2100),
					stringvalidator.RegexMatches(regexp.MustCompile("^[ -~\\t]*$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthClientId",
				TerraformName:       "authentication_oauth_client_id",
				MarkdownDescription: "The OAuth client ID. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.19.",
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
				SempName:            "authenticationOauthClientProxyName",
				TerraformName:       "authentication_oauth_client_proxy_name",
				MarkdownDescription: "The name of the proxy to use. Leave empty for no proxy. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.42.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthClientScope",
				TerraformName:       "authentication_oauth_client_scope",
				MarkdownDescription: "The OAuth scope. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.19.",
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
				MarkdownDescription: "The OAuth client secret. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.19.",
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
				MarkdownDescription: "The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.19.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 2048),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "authenticationOauthClientTokenExpiryDefault",
				TerraformName:       "authentication_oauth_client_token_expiry_default",
				MarkdownDescription: "The default expiry time for a token, in seconds. Only used when the token endpoint does not return an expiry time. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `900`. Available since SEMP API version 2.30.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(60, 86400),
				},
				Default: 900,
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthJwtProxyName",
				TerraformName:       "authentication_oauth_jwt_proxy_name",
				MarkdownDescription: "The name of the proxy to use. Leave empty for no proxy. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.42.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthJwtSecretKey",
				TerraformName:       "authentication_oauth_jwt_secret_key",
				MarkdownDescription: "The OAuth secret key used to sign the token request JWT. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.21.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 4096),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationOauthJwtTokenEndpoint",
				TerraformName:       "authentication_oauth_jwt_token_endpoint",
				MarkdownDescription: "The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.21.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 2048),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "authenticationOauthJwtTokenExpiryDefault",
				TerraformName:       "authentication_oauth_jwt_token_expiry_default",
				MarkdownDescription: "The default expiry time for a token, in seconds. Only used when the token endpoint does not return an expiry time. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `900`. Available since SEMP API version 2.30.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(60, 86400),
				},
				Default: 900,
			},
			{
				BaseType:            broker.String,
				SempName:            "authenticationScheme",
				TerraformName:       "authentication_scheme",
				MarkdownDescription: "The authentication scheme used by the REST Consumer to login to the REST host. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"none\"`. The allowed values and their meaning are:\n\n<pre>\n\"none\" - Login with no authentication. This may be useful for anonymous connections or when a REST Consumer does not require authentication.\n\"http-basic\" - Login with a username and optional password according to HTTP Basic authentication as per RFC 2616.\n\"client-certificate\" - Login with a client TLS certificate as per RFC 5246. Client certificate authentication is only available on TLS connections.\n\"http-header\" - Login with a specified HTTP header.\n\"oauth-client\" - Login with OAuth 2.0 client credentials.\n\"oauth-jwt\" - Login with OAuth (RFC 7523 JWT Profile).\n\"transparent\" - Login using the Authorization header from the message properties, if present. Transparent authentication passes along existing Authorization header metadata instead of discarding it. Note that if the message is coming from a REST producer, the REST service must be configured to forward the Authorization header.\n\"aws\" - Login using AWS Signature Version 4 authentication (AWS4-HMAC-SHA256).\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("none", "http-basic", "client-certificate", "http-header", "oauth-client", "oauth-jwt", "transparent", "aws"),
				},
				Default: "none",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable the REST Consumer. When disabled, no connections are initiated or messages delivered to this particular REST Consumer. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "httpMethod",
				TerraformName:       "http_method",
				MarkdownDescription: "The HTTP method to use (POST or PUT). This is used only when operating in the REST service \"messaging\" mode and is ignored in \"gateway\" mode. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"post\"`. The allowed values and their meaning are:\n\n<pre>\n\"post\" - Use the POST HTTP method.\n\"put\" - Use the PUT HTTP method.\n</pre>\n Available since SEMP API version 2.17.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("post", "put"),
				},
				Default: "post",
			},
			{
				BaseType:            broker.String,
				SempName:            "localInterface",
				TerraformName:       "local_interface",
				MarkdownDescription: "The interface that will be used for all outgoing connections associated with the REST Consumer. When unspecified, an interface is automatically chosen. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 15),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "maxPostWaitTime",
				TerraformName:       "max_post_wait_time",
				MarkdownDescription: "The maximum amount of time (in seconds) to wait for an HTTP POST response from the REST Consumer. Once this time is exceeded, the TCP connection is reset. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `30`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 300),
				},
				Default: 30,
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
				BaseType:            broker.Int64,
				SempName:            "outgoingConnectionCount",
				TerraformName:       "outgoing_connection_count",
				MarkdownDescription: "The number of concurrent TCP connections open to the REST Consumer. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 50),
				},
				Default: 3,
			},
			{
				BaseType:            broker.String,
				SempName:            "proxyName",
				TerraformName:       "proxy_name",
				MarkdownDescription: "The name of the proxy to use. Leave empty for no proxy. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since SEMP API version 2.36.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "remoteHost",
				TerraformName:       "remote_host",
				MarkdownDescription: "The IP address or DNS name to which the broker is to connect to deliver messages for the REST Consumer. A host value must be configured for the REST Consumer to be operationally up. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 253),
					stringvalidator.RegexMatches(regexp.MustCompile("^([0-9a-zA-Z_\\-\\.]*|([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|\\[([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}\\]|([0-9a-fA-F]{1,4}:){1,7}:|\\[([0-9a-fA-F]{1,4}:){1,7}:\\]|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|\\[([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}\\]|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|\\[([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}\\]|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|\\[([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}\\]|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|\\[([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}\\]|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|\\[([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}\\]|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|\\[[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})\\]|:((:[0-9a-fA-F]{1,4}){1,7}|:)|\\[:((:[0-9a-fA-F]{1,4}){1,7}|:)\\])$"), ""),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "remotePort",
				TerraformName:       "remote_port",
				MarkdownDescription: "The port associated with the host of the REST Consumer. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `8080`.",
				Requires:            []string{"tls_enabled"},
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("tls_enabled"),
					),
					int64validator.Between(1, 65535),
				},
				Default: 8080,
			},
			{
				BaseType:            broker.String,
				SempName:            "restConsumerName",
				TerraformName:       "rest_consumer_name",
				MarkdownDescription: "The name of the REST Consumer.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "restDeliveryPointName",
				TerraformName:       "rest_delivery_point_name",
				MarkdownDescription: "The name of the REST Delivery Point.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			{
				BaseType:            broker.Int64,
				SempName:            "retryDelay",
				TerraformName:       "retry_delay",
				MarkdownDescription: "The number of seconds that must pass before retrying the remote REST Consumer connection. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 300),
				},
				Default: 3,
			},
			{
				BaseType:            broker.String,
				SempName:            "tlsCipherSuiteList",
				TerraformName:       "tls_cipher_suite_list",
				MarkdownDescription: "The colon-separated list of cipher suites the REST Consumer uses in its encrypted connection. The value `\"default\"` implies all supported suites ordered from most secure to least secure. The list of default cipher suites is available in the `tlsCipherSuiteMsgBackboneDefaultList` attribute of the broker object in the Monitoring API. The REST Consumer should choose the first suite from this list that it supports. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"default\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 1559),
				},
				Default: "default",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "tlsEnabled",
				TerraformName:       "tls_enabled",
				MarkdownDescription: "Enable or disable encryption (TLS) for the REST Consumer. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Requires:            []string{"remote_port"},
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				BoolValidators: []validator.Bool{
					boolvalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("remote_port"),
					),
				},
				Default: false,
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
