resource "mikrotik_system_license" "foo" {
  account  = "account"
  password = "password"
  level    = "p1"
}

terraform {
  required_providers {
    mikrotik = {
      source = "hashicorp.com/ddelnano/mikrotik"
      version = "0.9.3"
    }
  }
}

provider "mikrotik" {
  # Configuration options
}