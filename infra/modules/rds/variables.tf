variable "db_identifier" {
    description = "The RDS instance identifier"
    type        = string
}

variable "vpc_id" {
    description = "The VPC ID where RDS will be deployed"
    type        = string
}

variable "private_subnet_ids" {
    description = "List of private subnet IDs for the DB subnet group"
    type        = list(string)
}

variable "ec2_security_group_id" {
    description = "Security group ID of the EC2 instance that needs access to RDS"
    type        = string
    default     = ""
}