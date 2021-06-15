provider "google" {
  project = var.project_id
  region  = "europe-west2"
  zone    = "europe-west2-a"
  // credentials = "./credentials.json"
}

provider "google-beta" {
  project = var.project_id
  region  = "europe-west2"
  zone    = "europe-west2-a"
  // credentials = "./credentials.json"
}

locals {
  enabled_services = toset([
    "iam.googleapis.com",
    "container.googleapis.com",
    "cloudbuild.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "firebase.googleapis.com",
  ])
}

resource "google_project_service" "services" {
  for_each = local.enabled_services

  disable_dependent_services=false
  disable_on_destroy=false
  project = var.project_id
  service   = each.value
}