# Cloud DNS
# resource "google_project_service" "enable_dns" {
#   project = var.project_id
#   service = "dns.googleapis.com"
# }
# resource "google_dns_managed_zone" "next_go_zone" {
#   name        = "rockstart-zone"
#   dns_name    = "${var.domain_name}."
#   description = "Managed zone for rockstart project"
# }
# resource "google_dns_record_set" "frontend_dns" {
#   depends_on  = [google_project_service.enable_dns, google_dns_managed_zone.next_go_zone]
#   name        = "${var.domain_name}."
#   type        = "A"
#   ttl         = 300
#   managed_zone = google_dns_managed_zone.next_go_zone.name
#   rrdatas     = [google_compute_global_address.frontend_ip.address]
# }