resource "azurerm_dns_cname_record" "cname" {
  name                = var.app-name
  resource_group_name = "dns"
  zone_name           = var.root-domain
  ttl                 = 300
  target_resource_id  = azurerm_static_web_app.website.id

  lifecycle {
    prevent_destroy = true
  }

  depends_on = [
    azurerm_static_web_app.website
  ]
}

data "cloudflare_zones" "root" {
  account = {
    id = var.cloudflare_account_id
  }

  name = var.root-domain
}

resource "cloudflare_dns_record" "dnsauth" {
  zone_id = data.cloudflare_zones.root.result[0].id
  name    = "_dnsauth.${var.app-name}"
  type    = "TXT"
  content = azurerm_static_web_app_custom_domain.domain.validation_token
  ttl     = 300

  lifecycle {
    prevent_destroy = true
  }
}

resource "cloudflare_dns_record" "cname" {
  zone_id = data.cloudflare_zones.root.result[0].id
  name    = var.app-name
  type    = "CNAME"
  content = azurerm_static_web_app.website.default_host_name
  ttl     = 300
  proxied = true

  lifecycle {
    prevent_destroy = true
  }

  depends_on = [
    azurerm_static_web_app.website
  ]
}
