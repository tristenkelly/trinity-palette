output "bucket_name" {
  description = "Name of the S3 bucket"
  value       = aws_s3_bucket.static_assets.id
}

output "bucket_domain_name" {
  description = "Domain name of the S3 bucket"
  value       = aws_s3_bucket.static_assets.bucket_domain_name
}

output "website_endpoint" {
  description = "Website endpoint of the S3 bucket"
  value       = aws_s3_bucket.static_assets.website_endpoint
}

output "bucket_arn" {
  description = "ARN of the S3 bucket"
  value       = aws_s3_bucket.static_assets.arn
}