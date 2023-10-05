terraform {
  required_providers {
    solacebroker = {
      source = "registry.terraform.io/solaceproducts/solacebroker"
    }
  }
}

provider "solacebroker" {
  username = "admin"
  password = "admin"
  url      = "http://localhost:8080"
}

resource "solacebroker_msg_vpn" "newone" {
  msg_vpn_name = "new"
  enabled      = true
}

resource "solacebroker_msg_vpn_queue" "q" {
  msg_vpn_name    = solacebroker_msg_vpn.newone.msg_vpn_name
  queue_name      = "red"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue_subscription" "foo" {
  msg_vpn_name       = solacebroker_msg_vpn_queue.q.msg_vpn_name
  queue_name         = solacebroker_msg_vpn_queue.q.queue_name
  subscription_topic = "foo/bar"
}

resource "solacebroker_msg_vpn_client_profile" "bar" {
  msg_vpn_name        = solacebroker_msg_vpn.newone.msg_vpn_name
  client_profile_name = "admin"
  lifecycle {
    create_before_destroy = true
  }
}

resource "solacebroker_msg_vpn_client_username" "username" {
  msg_vpn_name        = solacebroker_msg_vpn.newone.msg_vpn_name
  client_username     = "alice"
  client_profile_name = solacebroker_msg_vpn_client_profile.bar.client_profile_name
  enabled             = false
}
