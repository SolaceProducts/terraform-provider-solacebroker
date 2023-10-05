
resource "solacebroker_msg_vpn_queue" "q2" {
  msg_vpn_name    = solacebroker_msg_vpn.newone.msg_vpn_name
  queue_name      = "green"
  ingress_enabled = true
  egress_enabled  = true
  max_msg_size    = 12345
}
