---
page_title: "Solacebroker Provider Guide"
---

# Solace PubSub+ Software Event Broker (solacebroker) Provider

The `solacebroker` provider (Provider) supports Terraform CLI operations including basic CRUD (create, read, update, and delete) and import.

The Provider is leveraging the [SEMP (Solace Element Management Protocol)](https://docs.solace.com/Admin/SEMP/Using-SEMP.htm) REST API to configure the PubSub+ event broker. The API reference is available from the [Solace Documentation](https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/software-broker/config/index.html).

## Broker SEMP API access

The broker SEMP service, by default at port 8080 for HTTP and TLS port 1943 for HTTPS, must be accessible to the console running Terraform CLI.

The supported access credentials are basic authentication using username and password, and OAuth using a token. The two options are mutually exclusive and the provider will fail if both configured.

-> The [user access levels](https://docs.solace.com/Admin/CLI-User-Access-Levels.htm) associated with the credentials used must be properly configured on the broker so desired actions are authorized.

## SEMP API versioning and Provider broker compatibility

The SEMP API minor version reflects the supported set of objects, attributes, their properties and possible deprecations.

New versions of the PubSub+ event broker with new features typically require a newer SEMP API version that supports the new or updated objects, attributes, etc. The SEMP API version of a broker version can be determined from the [Solace documentation](https://docs.solace.com/Admin/SEMP/SEMP-API-Versions.htm#SEMP_v2_to_SolOS_Version_Mapping).

A given version of the Provider is built to support a specific version of the SEMP API. For the SEMP API version of the provider see the release notes in the GitHub repo.

* Broker versions at the same SEMP API version level as the Provider can be fully configured.
* Broker versions at a lower SEMP API version level than the Provider can be configured, with the exception of objects or attributes that have been deprecated and removed in the Provider's SEMP version. However, configuration will fail when attempting to configure objects or attributes that have been introduced in a later SEMP version than the broker supports.
* Broker versions at a higher SEMP API version level than the Provider can be configured for objects or attributes that are included in the Provider's SEMP version. Objects or attributes that have been introduced in a later SEMP version will be unknown to the Provider. Objects or attributes that have been deprecated in the broker SEMP version may result in configuration failure.

## Objects relationship

Broker inter-object references must be correctly encoded in Terraform configuration to have the apply work. It requires understanding of the PubSub+ event broker objects: it is recommended to consult the [SEMP API reference](https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/software-broker/config/index.htm) and especially "Identifying" attributes that give a hint to required already configured objects. 

## Mapping of SEMP API and Provider names

Terraform uses the [snake case](https://en.wikipedia.org/wiki/Snake_case) naming scheme, while SEMP uses camel case. Resources and datasource are also prefixed with the provider local name, `solacebroker_`.  For example, `solacebroker_msg_vpn` is the message-vpn resource name and `max_subscription_count` is the attribute for the maximum subscription count, since `MsgVpn` is the SEMP API object name and `maxSubscriptionCount` is the name of the SEMP attribute.

## Notes

Following limitations partly come from Terraform and partly from how the broker works.

* Terraform apply will not be atomic.  If interrupted by user, failure, reboot, switchover; the configuration changes may be partly applied and there will be no attempt to rollback anything.
* Terraform must be the authoritative source of configuration.  If there is any overlap between Terraform controlled configuration and either pre-existing configuration or modifications from other management interfaces; the behaviour will be undefined.
* Apply operations may impact broker AD performance; especially large changes.  This can be mitigated by throttling the configuration commands that are executed as part of the apply, but that itself may cause a commit to take a long time.
* Application of configuration may not be hitless.  Brief service interruptions may occur during an apply.  These can include a queue missing a published message, or clients being briefly disconnected.  These outages will be no different than if a current administrator manually makes an equivalent change to a broker.
* Inter object references must be correctly encoded in Terraform configuration to have the apply work.  It may not be possible to detect these problems until during the apply operation.