resource "google_project_service" "cloud_run_api" {
  service = "run.googleapis.com"
  project = var.project_id

  disable_on_destroy = false
}

resource "google_cloud_run_service" "backend" {
  depends_on = [google_project_service.cloud_run_api]

  name     = "${var.service_name}-cloudrun-backend"
  location = var.region

  template {
    spec {
      service_account_name = google_service_account.cloudrun_service_account.email

      containers {
        # TODO: こうすることで初回のterraform apply時にcloudbuildとの相互参照を回避できる
        # 2回目以降は正しいimageに変える
        image = "gcr.io/cloudrun/hello"
        # image = "${var.region}-docker.pkg.dev/${var.project_id}/${var.image_repo}/backend"
        ports {
          container_port = 50051
        }
      }
    }
  }

  metadata {
    annotations = {
      "run.googleapis.com/client-name" = "terraform"
    }
  }
}

output "backend_url" {
  value = google_cloud_run_service.backend.status[0].url
}
