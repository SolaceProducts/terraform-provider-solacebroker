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

resource "solacebroker_msg_vpn" "another" {
  msg_vpn_name = "another"
  enabled      = true
}

resource "solacebroker_msg_vpn_queue" "qa" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "red"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qbds" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "blue"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qca" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "orange"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qds" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "purple"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qedddss" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "green"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qeaa" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "yellow"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qesdss" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "indigo"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}
resource "solacebroker_msg_vpn_queue" "qedsdss" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "violet"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qedsd" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "cyan"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qeaad" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "bruge"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qed" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "pink"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "qes" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "pruplepink"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "one"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok2" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "two"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}
resource "solacebroker_msg_vpn_queue" "ok3" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "three"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok4" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "four"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok5" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "five"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok6" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "five"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok7" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "six"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok8" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "eight"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}

resource "solacebroker_msg_vpn_queue" "ok9" {
  msg_vpn_name    = solacebroker_msg_vpn.another.msg_vpn_name
  queue_name      = "nine"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54321
}
