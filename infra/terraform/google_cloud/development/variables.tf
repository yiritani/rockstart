variable "project_id" {
  type        = string
}

variable "region" {
  description = "GCPのリージョン"
  type        = string
  default     = "us-central1"
}

variable "dockerfile_backend" {
  description = "バックエンドのDockerfileパス"
  type        = string
  default     = "apps/backend/Dockerfile"
}

variable "service_name" {
  description = "サービス名"
  type        = string
  default     = "rockstart"
}

variable "image_repo" {
  description = "イメージリポジトリ名"
  type        = string
  default     = "rockstart"
}

variable "support_email" {
  description = "The support email for IAP"
  type        = string
  default     = "yusuke.iritani0909@gmail.com"
}

variable "access_user_email" {
  description = "The email address of the user to grant access to"
  type        = string
  default     = "yusuke.iritani0909@gmail.com"
}
