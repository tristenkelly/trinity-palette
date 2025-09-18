variable "bucket_name" {
  description = "Name of the S3 bucket for static assets"
  type        = string
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}