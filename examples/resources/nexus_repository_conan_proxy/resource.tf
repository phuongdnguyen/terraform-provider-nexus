resource "nexus_repository_conan_proxy" "conan_center" {
  name   = "conan-center"
  online = true

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
  }

  proxy {
    remote_url       = "https://center.conan.io"
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
