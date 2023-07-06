resource "nexus_repository_go_proxy" "golang_org" {
  name   = "golang-org"
  online = true

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
  }

  proxy {
    remote_url       = "https://proxy.golang.org/"
    content_max_age  = 1440
    metadata_max_age = 1440
  }

  negative_cache_enabled = true
  negative_cache_ttl     = 1440

  http_client {
    blocked    = false
    auto_block = true
  }
}


resource "nexus_repository_go_group" "group" {
  name   = "go-group"
  online = true

  group {
    member_names {
      name  = nexus_repository_go_proxy.golang_org.name
      order = 1
    }
  }

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
  }
}
