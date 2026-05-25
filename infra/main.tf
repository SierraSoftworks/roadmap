variable "app-name" {
  description = "The name of the static web app to create."
  type        = string
  default     = "roadmap"
}

variable "root-domain" {
  description = "The domain name which will be used for this website."
  type        = string
  default     = "sierrasoftworks.com"
}

variable "cloudflare_account_id" {
  description = "The Cloudflare account ID used to query DNS zones."
  type        = string
}

variable "tags" {
  description = "The tags which should apply to the resource."
  type        = map(string)
  default     = {}
}
