# https://github.com/datadrivers/terraform-provider-nexus/issues/352
provider "nexus" {
  insecure = true
  password = "123123123"
  url      = "http://127.0.0.1:8081"
  username = "admin"
}

terraform {
  required_version = ">= 0.14"

  required_providers {
    nexus = {
      source = "datadrivers/nexus"
    }
  }
}

resource "nexus_repository_raw_proxy" "github_raw_proxy_1" {
  name = "github-raw-proxy-1"

  storage {
    blob_store_name = "Default"
  }

  proxy {
    remote_url = "https://github.com/"
  }

  http_client {
    auto_block = false
    blocked    = false
  }
}

resource "nexus_repository_raw_proxy" "github_raw_proxy_2" {
  name = "github-raw-proxy-2"

  storage {
    blob_store_name = "Default"
  }

  proxy {
    remote_url = "https://github.com/"
  }

  http_client {
    auto_block = false
    blocked    = false
  }

  negative_cache_enabled = true
  negative_cache_ttl     = 1900
}
