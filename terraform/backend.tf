terraform {
  backend "gcs" {
    bucket = "go-feed-me-123897123-tfstate"
    credentials = "./credentials.json"
  }
}