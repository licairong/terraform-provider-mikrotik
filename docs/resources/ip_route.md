# mikrotik_ip_route (Resource)
Add an ip route.

## Example Usage
```terraform
data "mikrotik_ip_addresss" "ipaddrs" {}

resource "mikrotik_ip_route" "lan" {
  dst_address    = "192.168.2.0/24"
  gateway        = data.mikrotik_ip_addresss.ipaddrs.items.1.interface
  comment = "ip route"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dst_address` (String) Destination address.
- `gateway` (String) Gateway interface.

### Optional

- `comment` (String) The comment for the IP route assignment.
- `disabled` (Boolean) Whether to disable IP route. Default: `false`.

### Read-Only

- `id` (String) The ID of this resource.

