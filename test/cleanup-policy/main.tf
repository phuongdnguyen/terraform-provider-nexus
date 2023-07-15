terraform {
  required_providers {
    nexus = {
      source  = "nduyphuong/nexus"
      version = "1.23.0"
    }
  }
}

provider "nexus" {
  insecure = true
  password = "123123123"
  url      = "http://127.0.0.1:8081"
  username = "admin"
}

resource "nexus_cleanup_policy" "this" {
  name   = "test"
  format = "docker"
  criteria {
    last_downloaded = 180
    last_blob_updated = 180
    regex = "123"
  }
}
