resource "google_project_service" "artifact_registry_api" {
  service = "artifactregistry.googleapis.com"
  project = var.project_id
  disable_on_destroy = false
}


resource "google_artifact_registry_repository" "artifact_registry" {
  depends_on = [google_project_service.artifact_registry_api]
  description   = "rockstart"
  format        = "DOCKER"
  location      = var.region
  repository_id = var.image_repo
}