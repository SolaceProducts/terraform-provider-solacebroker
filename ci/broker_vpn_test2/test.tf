terraform {
  required_providers {
    solacebroker = {
      source = "registry.terraform.io/solaceproducts/solacebroker"
    }
  }
}

provider solacebroker {
  url            = "http://localhost:8080"
}

resource "solacebroker_msg_vpn" "test" {
  msg_vpn_name = "test"
}