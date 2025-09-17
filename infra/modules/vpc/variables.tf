variable "vpc_cidr" {
    description = "The CIDR block for the VPC"
    type        = string
    default     = "10.0.0.0/16"
}

variable "availability_zones" {
    description = "A list of availability zones in the region"
    type        = list(string)
}

variable "public_subnet_cidrs" {
    description = "A list of public subnet CIDR blocks"
    type        = list(string)
}

variable "private_subnet_cidrs" {
    description = "A list of private subnet CIDR blocks"
    type        = list(string)
}