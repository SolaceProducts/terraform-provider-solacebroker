---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solacebroker_oauth_profile_access_level_group Data Source - solacebroker"
subcategory: ""
description: |-
  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.
  Attribute|Identifying
  :---|:---:
  groupname|x
  oauthprofile_name|x
  A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.
  This has been available since SEMP API version 2.24.
---

# solacebroker_oauth_profile_access_level_group (Data Source)

The name of a group as it exists on the OAuth server being used to authenticate SEMP users.


Attribute|Identifying
:---|:---:
group_name|x
oauth_profile_name|x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since SEMP API version 2.24.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_name` (String) The name of the group.
- `oauth_profile_name` (String) The name of the OAuth profile.

### Read-Only

- `description` (String) A description for the group. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `""`.
- `global_access_level` (String) The global access level for this group. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `"none"`. The allowed values and their meaning are:

<pre>
"none" - User has no access to global data.
"read-only" - User has read-only access to global data.
"read-write" - User has read-write access to most global data.
"admin" - User has read-write access to all global data.
</pre>
- `msg_vpn_access_level` (String) The default message VPN access level for this group. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `"none"`. The allowed values and their meaning are:

<pre>
"none" - User has no access to a Message VPN.
"read-only" - User has read-only access to a Message VPN.
"read-write" - User has read-write access to most Message VPN settings.
</pre>
