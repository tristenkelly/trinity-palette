variable "db_identifier" {
    description = "The RDS instance identifier"
    type        = string
}

variable "vpc_id" {
    description = "The VPC ID where RDS will be deployed"
    type        = string
}

variable "private_subnet_ids" {
    description = "List of private subnet IDs for RDS"
    type        = list(string)
}