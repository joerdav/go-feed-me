resource "google_storage_bucket" "static" {
  name          = "gfm-static-assets"
  location      = "EU"
  force_destroy = true

  uniform_bucket_level_access = true

  cors {
    origin          = ["https://go-feed-me.joe-davidson.co.uk"]
    method          = ["GET", "HEAD", "PUT", "POST", "DELETE"]
    response_header = ["*"]
    max_age_seconds = 3600
  }
  depends_on = [ google_project_service.services ]
}

resource "google_storage_bucket_iam_member" "member" {
  bucket = google_storage_bucket.static.name
  role = "roles/storage.objectViewer"
  member = "allUsers"
}