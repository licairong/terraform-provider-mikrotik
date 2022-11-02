resource "mikrotik_interface_l2tp_client" "l2tp_client" {
  name = "l2tp-out3"
  password = "password"
  connect_to = "192.168.1.1"
  ipsec_secret = "password"
  use_ipsec = "yes"
  user = "admin"
  comment = "terraform acc test"
}