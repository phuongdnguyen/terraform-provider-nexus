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

resource "nexus_repository_docker_hosted" "docker-hosted-repos-rfrg" {
  name   = "docker-h-rfrg"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
    https_port       = 5001
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_hosted" "docker-hosted-repos-rint" {
  name   = "docker-h-rint"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
    https_port       = 5002
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_hosted" "docker-hosted-repos-rewt" {
  name   = "docker-h-rewt"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
    https_port       = 5003
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_group" "docker-group-repos" {
  name   = "docker-g-rewt"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
  }
  group {
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rfrg.name
      order = "1"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rint.name
      order = "2"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rewt.name
      order = "3"
    }

    # member_names = [
    #   {
    #     name  = nexus_repository_docker_hosted.docker-hosted-repos-rewt.name,
    #     order = "1"
    #   },
    #   {
    #     name  = nexus_repository_docker_hosted.docker-hosted-repos-rint.name,
    #     order = "2"
    #   },
    #   {
    #     name  = nexus_repository_docker_hosted.docker-hosted-repos-rfrg.name,
    #     order = "3"
    #   },
    # ]

    # member_names = [
    #   "docker-h-rewt",
    #   "docker-h-rfrg",
    #   "docker-h-rint",
    # ]
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_group" "docker-group-repos-1" {
  name   = "docker-g-1"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
  }
  group {
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rfrg.name
      order = "1"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rint.name
      order = "2"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rewt.name
      order = "3"
    }
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_group" "docker-group-repos-2" {
  name   = "docker-g-2"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
  }
  group {
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rfrg.name
      order = "1"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rint.name
      order = "2"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rewt.name
      order = "3"
    }
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_group" "docker-group-repos-3" {
  name   = "docker-g-3"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
  }
  group {
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rfrg.name
      order = "1"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rint.name
      order = "2"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rewt.name
      order = "3"
    }
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_group" "docker-group-repos-4" {
  name   = "docker-g-4"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
  }
  group {
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rfrg.name
      order = "1"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rint.name
      order = "2"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rewt.name
      order = "3"
    }
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}

resource "nexus_repository_docker_group" "docker-group-repos-5" {
  name   = "docker-g-5"
  online = true
  docker {
    force_basic_auth = false
    v1_enabled       = false
  }
  group {
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rfrg.name
      order = "1"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rint.name
      order = "2"
    }
    member_names {
      name  = nexus_repository_docker_hosted.docker-hosted-repos-rewt.name
      order = "3"
    }
  }
  storage {
    blob_store_name                = "Default"
    strict_content_type_validation = true
  }
}
