variable "project_id" {
    type = string
    default = "go-feed-me-123897123"
}

locals {
    services = [
        "basket",
        "browse",
        "container",
        "content",
        "details",
        "random"
        ]
    url_map = {
        "/content" = "content"
        "/apps/basket" = "basket"
        "/apps/browse" = "browse"
        "/apps/details" = "details"
        "/apps/random" = "random"
    }
    default_service = "container"
}