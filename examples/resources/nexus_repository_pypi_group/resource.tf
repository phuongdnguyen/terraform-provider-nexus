resource "nexus_repository_pypi_hosted" "internal" {
  name   = "internal"
  online = true

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
    write_policy                   = "ALLOW"
  }
}

resource "nexus_repository_pypi_proxy" "pypi_org" {
  name   = "pypi-org"
  online = true

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
  }

  proxy {
    remote_url       = "https://pypi.org"
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

resource "nexus_repository_pypi_group" "group" {
  name   = "pypi-group"
  online = true

  group {
    member_names {
      name  = nexus_repository_pypi_hosted.internal.name
      order = 1
    }
    member_names {
      name  = nexus_repository_pypi_proxy.pypi_org.name
      order = 2
    }
  }

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
  }
}
