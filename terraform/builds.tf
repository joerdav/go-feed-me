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
      owner = "joe-davidson1802"
      name  = "go-feed-me"
      push {
        branch = "^main$"
      }
    }

    included_files = ["src/${each.value}/**"]

    build {
      step {
        name = "gcr.io/cloud-builders/docker"
        args = [
          "buid", 
          "--tag=gcr.io/$PROJECT_ID/gfm-$${_NAME}:latest",
          "--tag=gcr.io/$PROJECT_ID/gfm-$${_NAME}:$COMMIT_SHA",
          "."
          ]
        timeout = "120s"
      }

      step {
        name = "gcr.io/cloud-builders/docker"
        args = [
          "push", 
          "gcr.io/$PROJECT_ID/gfm-$${_NAME}"
          ]
        timeout = "120s"
      }
      
      substitutions = {
        _NAME = each.value
      }

      artifacts {
        images = [
            "gcr.io/$PROJECT_ID/gfm-$${_NAME}:$COMMIT_SHA",
            "gcr.io/$PROJECT_ID/gfm-$${_NAME}:$latest"
        ]
      }
    }
}