resource "google_compute_region_network_endpoint_group" "cloudrun_neg" {
  for_each              = toset(local.services)
  name                  = "gfm-${each.value}-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "europe-west2"
  cloud_run {
    service = google_cloud_run_service.run_service[each.value].name
  }
}

resource "google_compute_url_map" "url_map" {
  name            = "gfm-url-map"
  default_service = module.lb-http.backend_services[local.default_service].self_link

  host_rule {
    hosts        = ["*"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = module.lb-http.backend_services[local.default_service].self_link

    dynamic "path_rule" {
      for_each = local.url_map
      content {
        paths = [
          "${path_rule.key}",
          "${path_rule.key}/*"
        ]
        service = module.lb-http.backend_services[path_rule.value].self_link
      }
    }
  }
}

locals {
  content_backend_list = [
    {
      contentbackend = {
        description = ""
        groups = [
          {
            group = google_compute_backend_bucket.static_backend.id
          },
        ]
        enable_cdn             = false
        security_policy        = null
        custom_request_headers = null

        iap_config = {
          enable               = false
          oauth2_client_id     = ""
          oauth2_client_secret = ""
        }
        log_config = {
          enable      = false
          sample_rate = null
        }
      }
      }
  ]
  backends_list = concat([
    for s in local.services : {
      "${s}" = {
        description = ""
        groups = [
          {
            group = google_compute_region_network_endpoint_group.cloudrun_neg[s].id
          },
        ]
        enable_cdn             = false
        security_policy        = null
        custom_request_headers = null

        iap_config = {
          enable               = false
          oauth2_client_id     = ""
          oauth2_client_secret = ""
        }
        log_config = {
          enable      = false
          sample_rate = null
        }
      }
    }
  ], local.content_backend_list)

  backends_map = { for item in local.backends_list :
    keys(item)[0] => values(item)[0]
  }
}

module "lb-http" {
  source  = "GoogleCloudPlatform/lb-http/google//modules/serverless_negs"
  version = "~> 4.5"

  url_map = google_compute_url_map.url_map.self_link
  create_url_map    = false
  project = var.project_id
  name    = "gfm-lb"

  managed_ssl_certificate_domains = ["go-feed-me.joe-davidson.co.uk"]
  ssl                             = true
  https_redirect                  = true

  backends = local.backends_map
}
