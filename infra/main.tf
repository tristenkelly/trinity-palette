terraform {
  #Comment out backend and run init/plan/apply to create all resources, run init apply again afterward to use the backend
  backend "s3" {
    bucket         = "the-trinity-pallette-terraform-state"
    key            = "tf-infra/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-state-locking"
    encrypt        = true
  }
  required_version = ">= 1.3.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

module "tf-state" {
  source      = "./modules/tf-state"
  bucket_name = "the-trinity-pallette-terraform-state"
}

module "vpc-infra" {
  source = "./modules/vpc"

  vpc_cidr             = local.vpc_cidr
  availability_zones   = local.availability_zones
  public_subnet_cidrs  = local.public_subnet_cidrs
  private_subnet_cidrs = local.private_subnet_cidrs
}

module "rds-infra" {
  source = "./modules/rds"

  db_identifier          = local.db_identifier
  vpc_id                = module.vpc-infra.vpc_id
  private_subnet_ids    = module.vpc-infra.private_subnet_ids
  ec2_security_group_id = module.ec2-infra.security_group_id
  
  depends_on = [module.vpc-infra, module.ec2-infra]
}

module "ec2-infra" {
  source = "./modules/ec2"

  ami_id        = local.ec2_ami_id
  instance_type = local.ec2_instance_type
  subnet_id     = module.vpc-infra.public_subnet_ids[0]  # Changed to public subnet
  vpc_id        = module.vpc-infra.vpc_id
  app_port      = local.app_port
  # key_pair_name = local.key_pair_name  # Uncomment if you want SSH access
}

# Outputs
output "ec2_public_ip" {
  description = "Public IP address of the EC2 instance"
  value       = module.ec2-infra.instance_public_ip
}

output "ec2_private_ip" {
  description = "Private IP address of the EC2 instance"
  value       = module.ec2-infra.instance_private_ip
}

output "rds_endpoint" {
  description = "RDS instance endpoint"
  value       = module.rds-infra.db_endpoint
  sensitive   = true
}
