---
page_title: "Command-line Terraform Configuration Generator Guide"
---

# Command-Line Terraform Configuration Generator

Normally, provider binaries are not run standalone, they are started and their services are used by Terraform CLI.

The `solacebroker` provider, however, includes an additional feature where you can run its binary outside of Terraform CLI. In this case, you can use the "generate" command on the provider binary to generate a Terraform HCL configuration file for a specified object and all its child objects known to the provider.

You can [locate](https://terra-farm.github.io/main/installation.html) the provider binary in the `.terraform/providers` directory of an existing Terraform configuration directory that uses the `solacebroker` provider.

## Important notes

You should review the generated configuration for the following:

* The provider block values in the generated configuration (URL, username, etc.) are exposed via Terraform input variables. Some write-only and related attribute values may also be assigned from input variables. It is recommended to check the variables created by the generator: you will need to assign value to those variables when applying the configuration or Terraform will prompt for the variable value.
* Some optional write-only attributes that cannot be determined by the generator if they were configured, are omitted from the generated configuration. You may need to add them manually.
* Default resources may be present that you can omit.
* It may be required to add a "depends_on" meta-argument between generated objects. Refer to the "System provisioned objects" section.
* The generator uses a naming scheme for the resources. You can update this by manually replacing the generated names.

## Usage

`<binary> generate [flags] <terraform resource address> <provider-specific identifier> <filename>`

* `<binary>` is the broker provider binary.
* `[flags]` are the [supported parameters](https://registry.terraform.io/providers/SolaceProducts/solacebroker/latest/docs/guides/config-generator#supported-parameters), which mirror the [configuration options for the provider object](https://registry.terraform.io/providers/SolaceProducts/solacebroker/latest/docs#schema), for example `--url=https://localhost:1943`. Parameters can alternatively be set via environment variables, for this example through setting `SOLACEBROKER_URL`.
* `<terraform resource address>` is the address of the specified object instance in the generated configuration, in the form of `<resource_type>.<resource_name>` (for example `solacebroker_msg_vpn.myvpn`). 
* `<provider-specific identifier>` is the import identifier of the specified object instance as in the Terraform Import command. The import identifier is available from the documentation of each resource type.
* `<filename>` is the name of the generated file.
This generator supports obtaining the configuration of software event brokers and will fail if applied against an appliance. This check may be overridden by setting the SOLACEBROKER_SKIP_API_CHECK=true environment variable.

Example:
```bash
SOLACEBROKER_USERNAME=admin SOLACEBROKER_PASSWORD=admin terraform-provider-solacebroker generate --url=https://localhost:8080 solacebroker_msg_vpn_queue.q default/test my-message-vpn-queue.tf
```

This will generate the configuration for queue `test` in message VPN `default`, and the configuration of all children, for example all subscriptions that have been configured to this queue.

### Supported parameters

The following parameters can be set as flags or environment variables (flags take precedence if both defined):

| Parameter                      | Required | Flag                  | Env var                      | Default |
|------------------------------- |-----------|-----------------------|------------------------------|---------|
| url | Yes | --url | SOLACEBROKER_URL | None |
| username (Note1)          | Yes       | --username  | SOLACEBROKER_USERNAME       | None    |
| password (Note1)         | No        | --password            | SOLACEBROKER_PASSWORD       | None    |
| bearer-token (Note1)     | No        | --bearer-token        | SOLACEBROKER_BEARER_TOKEN   | None    |
| insecure-skip-verify | No     | --insecure-skip-verify | SOLACEBROKER_INSECURE_SKIP_VERIFY | false |
| request-min-interval | No    | --request-min-interval | SOLACEBROKER_REQUEST_MIN_INTERVAL | 100ms |
| request-timeout-duration | No | --request-timeout-duration | SOLACEBROKER_REQUEST_TIMEOUT_DURATION | 1m |
| retries           | No        | --retries             | SOLACEBROKER_RETRIES        | 10    |
| retry-min-interval | No     | --retry-min-interval   | SOLACEBROKER_RETRY_MIN_INTERVAL | 3s |
| retry-max-interval | No     | --retry-max-interval   | SOLACEBROKER_RETRY_MAX_INTERVAL | 30s |
| skip-api-check    | No        | --skip-api-check      | SOLACEBROKER_SKIP_API_CHECK | false    |

Note1: Only one authentication method can be used at a time: either bearer-token or username/password.

## Attribute generation

For each object, all attributes will be generated as attributes on the corresponding resource with the exception of:
* attributes that are at the default value (as per the broker version corresponding to the broker provider)
* write-only attributes that cannot be determined if they were configured (not coupled with another non write-only attribute)

Write-only attributes that are coupled with another non write-only attribute will be generated as variable references. Variables for coupled attributes that are not write-only will have a commented-out default value with the value of the attribute, which you can choose to uncomment. Having no default means that Terraform will prompt for the variable value.

## System provisioned objects

System provisioned broker objects are created as a side-effect of creating other objects. These other objects are referred to as "parent object". The generator is attempting to recognize system provisioned objects and omit them from the configuration or add a warning comment, as direct creation of such objects will fail.

If an object's attribute is referencing a possible system-provisioned object, there may be a conflict at apply-time if the referenced object has not yet been created. The generator will add a comment when recognizing such references and it may be necessary to add a "depends_on" meta-argument between the referencoing and the parent objects to ensure proper create sequence.

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
| Explanation     | This indicates the resource attributes attempted to be fetched could not be read.          |
| Possible Action | Ensure identifiers values are consistent as set on the PubSub+ broker configured with.   |

| Error           | Error: Broker resource not found by terraform name xxx                                                     |
|-----------------|------------------------------------------------------------------------------------------------------------|
| Explanation     | This indicates the resource by name _xxx_ is not recognized by the generator.                              |
| Possible Action | Ensure the resource name used is available as a Terraform resource for the version of the provider in use. |

| Error           | Error: Broker check failed                                                                                  |
|-----------------|-------------------------------------------------------------------------------------------------------------|
| Explanation     | This indicates that broker platform does not match provider supported platform                              |
| Possible Action | Ensure that an appliance provider is used against a software broker platform and not an appliance platform. |
