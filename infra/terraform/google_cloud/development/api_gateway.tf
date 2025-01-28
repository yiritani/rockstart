# resource "google_api_gateway_api" "backend_api" {
#   api_id = "${var.service_name}-api"
#   project = var.project_id
# }

# resource "google_api_gateway_api_config" "backend_api_config" {
#   api = google_api_gateway_api.backend_api.id
#   api_config_id = "${var.service_name}-api-config"
  
#   openapi_documents {
#     document {
#       path = "openapi.yaml"
#       contents = file("openapi.yaml")
#     }
#   }
# }

# resource "google_api_gateway_gateway" "backend_gateway" {
#   gateway_id = "${var.service_name}-gateway"
#   api = google_api_gateway_api.backend_api.id
#   api_config = google_api_gateway_api_config.backend_api_config.id
#   project = var.project_id
#   region = var.region
# }
