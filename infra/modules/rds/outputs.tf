output "db_endpoint" {
  description = "The RDS instance endpoint"
  value       = aws_db_instance.default.endpoint
}

output "db_password" {
  description = "The master password for the database"
  value       = random_password.master.result
  sensitive   = true
}

output "db_username" {
  description = "The master username for the database"
  value       = aws_db_instance.default.username
}

output "db_name" {
  description = "The database name"
  value       = aws_db_instance.default.name
}
