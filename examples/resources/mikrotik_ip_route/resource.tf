data "mikrotik_ip_addresss" "ipaddrs" {}

output "ipaddrs" {
  value = data.mikrotik_ip_addresss.ipaddrs
}

resource "mikrotik_ip_route" "route" {
  dst_address    = "192.168.2.0/24"
  gateway        = data.mikrotik_ip_addresss.ipaddrs.items.1.interface
  comment = "ip route acc test"
}
