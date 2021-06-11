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
}

resource "google_compute_backend_bucket" "static_backend" {
  name        = "static-backend-bucket"
  description = ""
  bucket_name = google_storage_bucket.static.name
  enable_cdn  = false
}