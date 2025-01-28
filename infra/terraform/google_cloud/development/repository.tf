resource "google_artifact_registry_repository" "artifact_registry" {
  description   = "rockstart"
  format        = "DOCKER"
  location      = var.region
  repository_id = var.image_repo
}