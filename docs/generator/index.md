---
page_title: "solacebroker Generator"
subcategory: ""
description: |-
  Generates a Terraform configuration file for a specified PubSubPlus Broker object and all child objects known to the provider.
  Note that this is not a terraform command.
---

### Generator

The generate command on the provider binary generates a Terraform configuration file for the specified object and all
child objects known to the provider.
This is not a Terraform command. One can download the provider binary and can execute that binary with the `generate`
command to generate a Terraform configuration file from the current configuration of a PubSubPlus broker.

`<binary> generate <terraform resource address> <provider-specific identifier> <filename>`

- `<binary>` is the broker provider binary
- `<terraform resource address>` is the terraform resource address, for example `https://mybroker.example.org:1943/`.
- `<provider-specific identifier>` are the similar to the Terraform Import command. This is the resource name and possible values to find a specific resource.
- `<filename>` is the desirable name of the generated filename.
- There are also supported options, which mirror the configuration options for the provider object. These can be
  found [here](#supported-options)

### Usage

```shell
terraform-provider-solacebroker -h

Usage:
    terraform-provider-solacebroker [command]

Available Commands:
generate    Generates a Terraform configuration file for a specified PubSubPlus Broker object and all child objects known to the provider
help        Help about any command
version     Provides version information about the current binary
```

To `generate` the configuration, make sure all ENVIRONMENT VARIABLES, which mirrors the configuration options for the
provider object are set. The list of variables
are listed [here](#supported-options).

For example:
`SOLACEBROKER_USERNAME=admin SOLACEBROKER_PASSWORD=admin terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn.mq default my-messagevpn.tf`

This command would create a file `my-messagevpn.tf` that contains a resource definition for the default message VPN resource and
any child objects, assuming the appropriate broker credentials were set in environment variables.

Note: For objects with no child object, the file will only contain resource definition for that object.

For example:
`SOLACEBROKER_USERNAME=admin SOLACEBROKER_PASSWORD=admin terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn_queue.q default/test my-message-vpn-queue.tf`

This command would create a file `my-message-vpn-queue.tf` that contains the msg_vpn_queue resource , `test`  for the
message VPN, `default`, assuming a msg_vpn_queue resource called `test` exists for the message VPN, `default`.

### Troubleshooting

The following issues may arise while using the generator.

| Error           | SEMP call failed. unexpected status 401 (401 Unauthorized)                 |
|-----------------|----------------------------------------------------------------------------|
| Explanation     | Configurations to connect to the PubSubPlus Broker not accurate.           |
| Possible Action | Check and confirm, configuration details to PubSubPlus Broker are accurate |

| Error           | SOLACEBROKER_xxx is mandatory but not available                                    |
|-----------------|------------------------------------------------------------------------------------|
| Explanation     | A mandatory parameter which is required to connect to PubSubPlus broker is missing |
| Possible Action | Confirm if all [mandatory parameters](#supported-options) are correctly set.       |

| Error           | Error: Too many provider specific identifiers. Required identifiers: [{xxx}]                                                                                                                                        |
|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Explanation     | This indicates that identifiers specific to the provider are set in an ambiguous manner.                                                                                                                            |
| Possible Action | Ensure all identifiers are available and separated by `/` where needed. For example a msgVpnName will require `msgVpnName`, however a specific queueName under a specific msgVpnName will be `msgVpnName/queueName` |

| Error           | SEMP called failed. resource not found on path /xxx/xxx                                  |
|-----------------|------------------------------------------------------------------------------------------|
| Explanation     | This indicates the resource attributes attempted to be fetch could not be read           |
| Possible Action | Ensure identifiers values are consistent as set on the PubSubPlus broker configured with |

| Error           | Error: Broker resource not found by terraform name xxx                                                     |
|-----------------|------------------------------------------------------------------------------------------------------------|
| Explanation     | This indicates the resource by name _xxx_ is not recognized by the generator                               |
| Possible Action | Ensure the resource name used is available as a Terraform resource for the version of the provider in use. |

### Supported Options

These parameters could be set as ENVIRONMENT VARIABLES. When used as environment variables
each parameter should be preceded with _SOLACEBROKER__. For example for a PubSubPlus broker using username and password
_**admin/admin**_
will be:

`SOLACEBROKER_USERNAME=admin SOLACEBROKER_PASSWORD=admin`

- `bearer_token`, (String, Sensitive, Mandatory if `password` will not be provided)
- `insecure_skip_verify` (Boolean) Disable validation of server SSL certificates, accept/ignore self-signed.
- `password` (String, Sensitive, Mandatory is `bearer_token` will not be provided)
- `request_min_interval` (String)
- `request_timeout_duration` (String)
- `retries` (Number)
- `retry_max_interval` (String)
- `retry_min_interval` (String)
- `username` (String, Mandatory) The username for the broker request.
