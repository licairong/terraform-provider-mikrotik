# The ID argument (*17) is a MikroTik's internal id.
# It can be obtained via CLI:
#
# [admin@MikroTik] /ip pool> :put [ find where name=pool-name]
# *17
terraform import mikrotik_interface_l2tp_client.l2tp_clien '*17'
