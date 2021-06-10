resource "google_service_account" "default" {
  account_id   = "service-account-id"
  display_name = "Service Account"
}

data "google_client_config" "provider" {}

resource "google_container_cluster" "primary" {
  name               = "gfm-cluster"
  initial_node_count = 3
  node_config {
    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.default.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
  timeouts {
    create = "30m"
    update = "40m"
  }
}

provider "kubectl" {
  load_config_file       = false
  host  = "https://${google_container_cluster.primary.endpoint}"
  token = data.google_client_config.provider.access_token
  cluster_ca_certificate = base64decode(
    google_container_cluster.primary.master_auth[0].cluster_ca_certificate,
  )
}

data "kubectl_filename_list" "manifests" {
    pattern = "../kubernetes/*.yaml"
}

resource "kubectl_manifest" "gfm" {
    count = length(data.kubectl_filename_list.manifests.matches)
    yaml_body = file(element(data.kubectl_filename_list.manifests.matches, count.index))
}