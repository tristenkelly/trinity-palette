variable "instance_type" {
    description = "The type of instance to use"
    type        = string
    default     = "t3.micro"
}

variable "ami_id" {
    description = "The AMI ID for the instance"
    type        = string
}

variable "subnet_id" {
    description = "The subnet ID where the instance will be launched"
    type        = string
}

variable "vpc_id" {
    description = "The VPC ID where the security group will be created"
    type        = string
}

variable "app_port" {
    description = "The port the application runs on"
    type        = number
    default     = 8080
}

variable "user_data" {
    description = "User data script for EC2 initialization"
    type        = string
    default     = ""
}