terraform {
  required_providers {
    solacebroker = {
      source = "solaceproducts/solacebroker"
    }
    random = {
      source = "hashicorp/random"
    }
  }
}

resource "random_password" "admin_password" {
  length  = 12
  special = false
}

output "admin_password" {
  value = random_password.admin_password.result
  sensitive = true
}

provider "solacebroker" {
  username = "admin"
  password = random_password.admin_password.result
  url            = "http://localhost:8080"
}

resource "solacebroker_msg_vpn_queue" "q" {
  msg_vpn_name = "default"
  queue_name = "red"
  ingress_enabled = true
  egress_enabled = true
  max_msg_size = 54321
}
