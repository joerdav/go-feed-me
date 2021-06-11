variable "project_id" {
    type = string
    default = "go-feed-me-123897123"
}

locals {
    services = [
        "basket",
        "browse",
        "container",
        "details",
        "random"
        ]
    url_map = {
        "/apps/basket" = "basket"
        "/apps/browse" = "browse"
        "/apps/details" = "details"
        "/apps/random" = "random"
    }
    default_service = "container"
}