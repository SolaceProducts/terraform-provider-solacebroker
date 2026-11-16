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
  url      = "https://localhost:1943"
}

resource "solacebroker_msg_vpn" "fips" {
  msg_vpn_name = "fips"
  enabled      = true
}

resource "solacebroker_msg_vpn_queue" "fips" {
  msg_vpn_name    = solacebroker_msg_vpn.fips.msg_vpn_name
  queue_name      = "fips"
  ingress_enabled = true
  egress_enabled  = true
}
