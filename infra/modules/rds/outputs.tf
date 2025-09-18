output "db_instance_id" {
  description = "The RDS instance ID"
  value       = aws_db_instance.default.id
}

output "db_endpoint" {
  description = "RDS instance endpoint"
  value       = aws_db_instance.default.endpoint
}

output "db_port" {
  description = "RDS instance port"
  value       = aws_db_instance.default.port
}

output "security_group_id" {
  description = "ID of the RDS security group"
  value       = aws_security_group.rds_sg.id
}