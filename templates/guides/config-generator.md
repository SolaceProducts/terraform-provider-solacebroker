---
page_title: "Command-line Terraform Configuration Generator Guide"
---

# Command-Line Terraform Configuration Generator

Normally, provider binaries are not run standalone, they are started and their services are used by Terraform CLI.

The `solacebroker` provider, however, includes an additional feature where you can run its binary outside of Terraform CLI. In this case, you can use the "generate" command on the provider binary to generate a Terraform HLC configuration file for a specified object and all child objects known to the provider.

You can [locate](https://terra-farm.github.io/main/installation.html) the provider binary in the `.terraform/providers` directory of an existing Terraform configuration directory that uses the `solacebroker` provider.

You can run the provider binary directly with the `generate` command to generate a Terraform configuration file from the current configuration of a PubSub+ broker.

`<binary> generate <broker URL> <provider-specific identifier> <filename>`

- `<binary>` is the broker provider binary.
- `<broker URL>` is the broker address, for example `https://mybroker.example.org:1943/`.
- `<provider-specific identifier>` is similar to the Terraform Import command. This is the resource name and possible values to find a specific resource.
- `<filename>` is the desirable name of the generated filename.
- There are also supported options, which mirror the configuration options for the provider object. These can be found [here](#supported-options).

This generator supports configuring software event brokers and will fail if applied against an appliance. This check may be overridden by setting the `SOLACEBROKER_SKIP_API_CHECK=true` environment variable.

## Important notes

You should review the generated configuration for the following:

* You may need to update provider configuration values (URL, username, etc.)
* Write-only attributes, such as passwords, are omitted from the config as they cannot be read from the broker configuration. You need to add them manually.
* Default resources may be present that you can omit.
* The generator uses a naming scheme for the resources. You can update this by manually replacing the generated names.

## Usage

```shell
terraform-provider-solacebroker -h

Usage:
    terraform-provider-solacebroker [command]

Available Commands:
generate    Generates a Terraform configuration file for a specified PubSub+ Broker object and all child objects known to the provider
help        Help about any command
version     Provides version information about the current binary
```

To `generate` the configuration, make sure all ENVIRONMENT VARIABLES, which mirror the configuration options for the
provider object are set. You can find the list of variables [here](#supported-options).

For example:
`SOLACEBROKER_USERNAME=admin SOLACEBROKER_PASSWORD=admin terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn.mq default my-messagevpn.tf`

This command would create a file `my-messagevpn.tf` that contains a resource definition for the default Message VPN resource and
any child objects, assuming the appropriate broker credentials were set in environment variables.

Note: For objects with no child object, the file will only contain a resource definition for that object.

For example:
`SOLACEBROKER_USERNAME=admin SOLACEBROKER_PASSWORD=admin terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn_queue.q default/test my-message-vpn-queue.tf`

This command would create a file `my-message-vpn-queue.tf` that contains the msg_vpn_queue resource , `test`  for the
Message VPN, `default`, assuming a msg_vpn_queue resource called `test` exists for the Message VPN, `default`.

### Supported Options

The following parameters can be set as ENVIRONMENT VARIABLES. When used as an environment variable,
each parameter must be preceded with _SOLACEBROKER__. An example for a PubSub+ broker using username and password
_**admin/password**_
would be:

`SOLACEBROKER_USERNAME=admin SOLACEBROKER_PASSWORD=password`

- `SOLACEBROKER_BEARER_TOKEN` (String, Sensitive, Mandatory if `password` not provided)
- `SOLACEBROKER_INSECURE_SKIP_VERIFY` (Boolean) Disable validation of server SSL certificates, accept/ignore self-signed.
- `SOLACEBROKER_PASSWORD` (String, Sensitive, Mandatory if `bearer_token` not provided)
- `SOLACEBROKER_REQUEST_MIN_INTERVAL` (String)
- `SOLACEBROKER_REQUEST_TIMEOUT_DURATION` (String)
- `SOLACEBROKER_RETRIES` (Number)
- `SOLACEBROKER_RETRY_MAX_INTERVAL` (String)
- `SOLACEBROKER_RETRY_MIN_INTERVAL` (String)
- `SOLACEBROKER_SKIP_API_CHECK` (String) Disable validation of the broker SEMP API for supported platform and minimum version.
- `SOLACEBROKER_USERNAME` (String, Mandatory) The username for the broker request.

## Troubleshooting

The following issues may arise while using the generator.

| Error           | SEMP call failed. unexpected status 401 (401 Unauthorized)                 |
|-----------------|----------------------------------------------------------------------------|
| Explanation     | Configurations to connect to the PubSub+ broker not accurate.              |
| Possible Action | Check and confirm, configuration details to PubSub+ broker are accurate.   |

| Error           | SOLACEBROKER_xxx is mandatory but not available                                    |
|-----------------|------------------------------------------------------------------------------------|
| Explanation     | A mandatory parameter which is required to connect to PubSub+ broker is missing.   |
| Possible Action | Confirm if all [mandatory parameters](#supported-options) are correctly set.       |

| Error           | Error: Too many provider specific identifiers. Required identifiers: [{xxx}] |
|-----------------|------------------------------------------------------------------------------|
| Explanation     | This indicates that identifiers specific to the provider are set in an ambiguous manner. |
| Possible Action | Ensure all identifiers are available and separated by `/` where needed. For example a msgVpnName will require `msgVpnName`, however a specific queueName under a specific msgVpnName will be `msgVpnName/queueName`. |

| Error           | SEMP called failed. resource not found on path /xxx/xxx                                  |
|-----------------|------------------------------------------------------------------------------------------|
| Explanation     | This indicates the resource attributes attempted to be fetch could not be read.          |
| Possible Action | Ensure identifiers values are consistent as set on the PubSub+ broker configured with.   |

| Error           | Error: Broker resource not found by terraform name xxx                                                     |
|-----------------|------------------------------------------------------------------------------------------------------------|
| Explanation     | This indicates the resource by name _xxx_ is not recognized by the generator.                              |
| Possible Action | Ensure the resource name used is available as a Terraform resource for the version of the provider in use. |

| Error           | Error: Broker check failed                                                                                  |
|-----------------|-------------------------------------------------------------------------------------------------------------|
| Explanation     | This indicates that broker platform does not match provider supported platform                              |
| Possible Action | Ensure that an appliance provider is used against a software broker platform and not an appliance platform. |
