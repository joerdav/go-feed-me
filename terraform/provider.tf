provider "google" {
  project = var.project_id
  region  = "europe-west2"
  zone    = "europe-west2-a"
  credentials = "./credentials.json"
}

resource "google_project_service" "iam" {
  project = var.project_id
  service   = "iam.googleapis.com"
}

resource "google_project_service" "container" {
  project = var.project_id
  service   = "container.googleapis.com"
}

resource "google_project_service" "cloudbuild" {
  project = var.project_id
  service   = "cloudbuild.googleapis.com"
}

resource "google_project_service" "cloudresourcemanager" {
  project = var.project_id
  service   = "cloudresourcemanager.googleapis.com"
}