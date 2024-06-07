terraform {
  required_providers {
    solacebroker = {
      source = "registry.terraform.io/solaceproducts/solacebroker"
    }
  }
}

provider solacebroker {
  username       = "admin"
  password       = "admin"
  # url            = "http://192.168.132.23:8080"
  url            = "http://localhost:8080"
  # skip_api_check  = true
  # request_min_interval     = "1s"
  insecure_skip_verify = true
  retries        = 10
  retry_min_interval = "2s"
  retry_max_interval = "20s"
  request_timeout_duration = "120s"
}

resource "solacebroker_msg_vpn" "newone" {
  msg_vpn_name = "new2"
  # set to default
  authentication_basic_enabled = true
  replication_queue_max_msg_spool_usage = 60000
  # set to non-default
  enabled      = true
  jndi_enabled = true
  #jndi_enabled = false
  replication_queue_reject_msg_to_sender_on_discard_enabled = false
  event_transaction_count_threshold = { "clear_percent" = 20, "set_percent" = 80 }
  # set to broker default
  # max_connection_count = 100
  # not set to broker default
  # max_subscription_count = 27
  event_egress_msg_rate_threshold     = { "clear_value" = 40, "set_value" = 50 }
  event_ingress_msg_rate_threshold    = { "clear_value" = 40, "set_value" = 50 }
}

resource "solacebroker_msg_vpn_queue" "q" {
  msg_vpn_name    = solacebroker_msg_vpn.newone.msg_vpn_name
  queue_name      = "aperfectly/$/valid/$topic/$$"
  # ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 54322
}

resource "solacebroker_msg_vpn_queue_subscription" "foo" {
  msg_vpn_name       = solacebroker_msg_vpn_queue.q.msg_vpn_name
  queue_name         = solacebroker_msg_vpn_queue.q.queue_name
  subscription_topic = "foo/bar"
}
