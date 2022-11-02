# The ID argument (*17) is a MikroTik's internal id.
# It can be obtained via CLI:
#
# [admin@MikroTik] /ppp secret> :put [ find where name=pool-name]
# *17
terraform import mikrotik_ppp_secret.secret '*17'
