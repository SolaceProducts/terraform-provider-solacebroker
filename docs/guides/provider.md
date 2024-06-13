---
page_title: "Solacebroker Provider Guide"
---

# Solace PubSub+ Software Event Broker (solacebroker) Provider

The `solacebroker` provider supports Terraform CLI operations including basic CRUD (create, read, update, and delete) and import.

The provider leverages the [SEMP (Solace Element Management Protocol)](https://docs.solace.com/Admin/SEMP/Using-SEMP.htm) REST API to configure the PubSub+ event broker. The API reference is available from the [Solace PubSub+ documentation](https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/appliance/config/index.html).

This provider supports configuring software event brokers and will fail if applied against an appliance. This check may be overridden by specifying the `skip_api_check = true` configuration argument.

## Mapping of SEMP API and Provider Names

Terraform uses the [snake case](https://en.wikipedia.org/wiki/Snake_case) naming scheme, while SEMP uses camel case. Resources and datasource are also prefixed with the provider local name, `solacebroker_`.  For example, `solacebroker_msg_vpn` is the message-vpn resource name and `max_subscription_count` is the attribute for the maximum subscription count, since `MsgVpn` is the SEMP API object name and `maxSubscriptionCount` is the name of the SEMP attribute.

## Broker SEMP API Access

The broker SEMP service, by default at port 8080 for HTTP and TLS port 1943 for HTTPS, must be accessible to the console running Terraform CLI.

The supported access credentials are basic authentication using username and password, and OAuth using a token. The two options are mutually exclusive and the provider will fail if both are configured.

-> The [user access levels](https://docs.solace.com/Admin/CLI-User-Access-Levels.htm) associated with the credentials used must be properly configured on the broker so that the desired actions are authorized.

## SEMP API Versioning and Provider Broker Compatibility

The SEMP API minor version reflects the supported set of objects, attributes, their properties and possible deprecations.

New versions of the PubSub+ event broker with new features typically require a newer SEMP API version that supports the new or updated objects, attributes, etc. The SEMP API version of a broker version can be determined from the [Solace PubSub+ documentation](https://docs.solace.com/Admin/SEMP/SEMP-API-Versions.htm#SEMP_v2_to_SolOS_Version_Mapping).

A given version of the provider is built to support a specific version of the SEMP API. For the SEMP API version of the provider and corresponding broker version, refer to the [Version Compatibility section](https://docs.solace.com/Admin/SEMP/Declarative-SEMP.htm#Version) of the Solace PubSub+ documentation.

* Broker versions at the same SEMP API version level as the provider can be fully configured.
* Broker versions at a lower SEMP API version level than the provider can be configured, except for objects or attributes that have been deprecated and removed in the provider's SEMP version. However, configuration will fail when attempting to configure objects or attributes that have been introduced in a later SEMP version than the broker supports.
* Broker versions at a higher SEMP API version level than the provider can be configured for objects or attributes that are included in the provider's SEMP version. Objects or attributes that have been introduced in a later SEMP version will be unknown to the provider. Objects or attributes that have been deprecated in the broker SEMP version may result in configuration failure.

## Object Relationships

Broker inter-object references must be correctly encoded in Terraform configuration to have the apply operation work. This requires an understanding of the PubSub+ event broker objects. For more information about each object consult the [SEMP API reference](https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/software-broker/config/index.htm) and especially look for "Identifying" attributes that give a hint to required pre-existing objects.
For example:

```terraform
resource "solacebroker_msg_vpn" "test" {
  # on the resource itself, specify the value
  msg_vpn_name        = "new"
  # ... other attributes
}

resource "solacebroker_msg_vpn_queue" "q" {
  # on dependent resources, specify as a reference so
  # that Terraform creates the referenced object first
  msg_vpn_name    = solacebroker_msg_vpn.test.msg_vpn_name
  # ... other attributes
}
```

## The Broker Object

The Broker object is the `solacebroker_broker` resource. This object contains global configuration for the PubSub+ event broker.

The Broker object differs from all other objects as it always exists for a given broker and can only be updated.

-> Important: only attributes that are specified will be set to their configured value. Unspecified attributes will not be set to their default-attribute value. This may result in `terraform plan` indicating a change to set attributes to default even after an `apply`, for example after removing an attribute from the configuration.

## Default Objects

There are objects that are preexisting defaults and cannot be created or destroyed, only updated. The default Message VPN and the default client profile are examples of this. Any direct attempt to remove these resources will fail. If the parent object is not a default object then deleting that parent will also remove its child default object.

> If for example a configuration includes creating a non-default Message VPN and modifying its default client-profile then when destroying the configuration the provider would report an error about removing the client-profile. But the client-profile object will be eventually deleted because the whole VPN will also be deleted, which includes the default client-profile.

## Broker-Defined Attributes

Some attributes don't have a default value. In this case their value will be determined by the broker. Typically, these defaults depend on the broker scaling settings. While Terraform plan and apply operations function the same way as with other attributes, import will set the Terraform state of the attribute to the broker value (instead of null), even if they were set at default. You can use subsequent plan and apply operations to fix this.

## Object Type Attributes

An object type attribute is a collection of attributes, for example `"event_ingress_msg_rate_threshold": { "clear_value": 2000000, "set_value": 5000000 }`. Note that due to Terraform provider framework limitations, there is no error reported when providing unknown nested attributes.

## Importing Resources

When [importing resources to Terraform](https://developer.hashicorp.com/terraform/language/import#syntax) an `id` is required. This `id` shall be constructed as a path from the highest parent object down to the resource.

For example, when importing a `solacebroker_msg_vpn_queue_subscription`, the parent relationship is `msg_vpn` > `msg_vpn_queue` > `msg_vpn_queue_subscription`. To construct the `id`, concatenate the identifications of parents and the particular resource identification, separated by `/` (slash). Also note that elements containing `/` must be URL-encoded.

For this example:
```
id = <msg-vpn-name>/<queue-name>/<url-encoded subscription-name>
# using my-vpn, my-queue, a/b/c
id = my-vpn/my-queue/a%2Fb%2Fc
```

## Notes

* Terraform `apply` is not atomic.  If interrupted by a user, failure, reboot, or switchover the configuration changes may be partly applied.  Terraform does not perform rollbacks.
* Terraform must be the authoritative source of configuration.  If there is any overlap between Terraform controlled configuration and either pre-existing configuration or modifications from other management interfaces the behaviour will be undefined.
* Apply operations may impact broker AD performance, especially large changes.  The `request_min_interval` attribute on the provider limits the request rate and can be adjusted to control the impact.
* Application of configuration may cause brief service interruptions to the resources affected.  These can include a queue missing a published message or clients being briefly disconnected.  These outages are no different from a current administrator manually making an equivalent change to a broker.