resource "mikrotik_ppp_secret" "secret" {
  name    = "test"
  password = "password"
  profile = "l2tp"
  routes = "192.168.1.0/24,192.168.2.0/24"
  service = "l2tp"
  comment = "ppp secret acc test"
}
