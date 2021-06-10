locals {
    services = [
        "basket",
        "browse",
        "container",
        "content",
        "details",
        "random"]
}

resource "google_cloudbuild_trigger" "service" {
    for_each = toset(local.services)

    name = "gfm-${each.value}-build"

    github {
      owner = "Joe-Davidson1802"
      name  = "go-feed-me"
      push {
        branch = "^main$"
      }
    }

    included_files = ["src/${each.value}/**"]

    substitutions = {
      _NAME = each.value
    }

    filename = "src/${each.value}/cloudbuild.yaml"
}

resource "google_cloudbuild_trigger" "infra" {
    name = "gfm-infra-build"

    github {
      owner = "Joe-Davidson1802"
      name  = "go-feed-me"
      push {
        branch = "^main$"
      }
    }

    included_files = ["terraform/**"]

    filename = "terraform/cloudbuild.yaml"
}