data "cloudflare_zones" "root" {
  account = {
    id = var.cloudflare_account_id
  }

  name = var.root-domain
}

resource "cloudflare_dns_record" "cname" {
  zone_id = data.cloudflare_zones.root.result[0].id
  name    = var.app-name
  type    = "CNAME"
  content = "sierrasoftworks.github.io"
  ttl     = 1
  proxied = false
}
