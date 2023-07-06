resource "nexus_repository_npm_hosted" "internal" {
  name   = "internal"
  online = true

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
    write_policy                   = "ALLOW"
  }
}

resource "nexus_repository_npm_group" "group" {
  name   = "npm-group"
  online = true

  group {
    member_names {
      name  = nexus_repository_npm_hosted.internal.name
      order = 1
    }
  }

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
  }
}
