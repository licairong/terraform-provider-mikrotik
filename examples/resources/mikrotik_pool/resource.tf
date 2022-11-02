resource "mikrotik_pool" "pool" {
  name    = "pool-name"
  ranges  = "172.16.0.6-172.16.0.12"
  comment = "ip pool with range specified"
}

terraform {
  required_providers {
    mikrotik = {
      source = "hashicorp.com/ddelnano/mikrotik"
      version = "0.9.1"
    }
  }
}

provider "mikrotik" {
  # Configuration options
}