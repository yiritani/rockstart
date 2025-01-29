resource "google_api_gateway_api" "api" {
  provider = google-beta
  project = var.project_id
  api_id = "backend-api"
  display_name = "Backend API"
}

resource "google_api_gateway_api_config" "config" {
  depends_on = [google_api_gateway_api.api]
  provider = google-beta
  project = var.project_id
  api = google_api_gateway_api.api.api_id
  api_config_id = "backend-api-config"

  openapi_documents {
    document {
      path = "../../../../apps/backend/src/_generated/openapi/api_definition.swagger.yaml"
      contents = base64encode(
        join("\n", [
          file("../../../../apps/backend/src/_generated/openapi/api_definition.swagger.yaml"),
          "x-google-backend:",
          "  address: ${google_cloud_run_service.backend.status[0].url}",
          "  path_translation: CONSTANT_ADDRESS"
        ])
      )
    }
  }
}

resource "google_api_gateway_gateway" "gateway" {
  provider = google-beta
  project = var.project_id
  gateway_id = "backend-gateway"
  api_config = google_api_gateway_api_config.config.id
  region = "us-central1"
}
