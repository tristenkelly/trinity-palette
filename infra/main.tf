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

module "s3-static" {
  source      = "./modules/s3-static"
  bucket_name = "the-trinity-pallette-static-assets"
  environment = "production"
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

  db_identifier      = local.db_identifier
  vpc_id            = module.vpc-infra.vpc_id
  private_subnet_ids = module.vpc-infra.private_subnet_ids
  
  depends_on = [module.vpc-infra]
}

module "ec2-infra" {
  source = "./modules/ec2"

  ami_id        = local.ec2_ami_id
  instance_type = local.ec2_instance_type
  subnet_id     = module.vpc-infra.public_subnet_ids[0]
  vpc_id        = module.vpc-infra.vpc_id
  app_port      = local.app_port
  user_data     = local.user_data
}

# Store RDS credentials in Systems Manager Parameter Store
resource "aws_ssm_parameter" "rds_endpoint" {
  name  = "/trinity-palette/database/endpoint"
  type  = "String"
  value = module.rds-infra.db_endpoint
  
  tags = {
    Environment = "production"
    Project     = "trinity-palette"
  }
  
  depends_on = [module.rds-infra]
}

resource "aws_ssm_parameter" "rds_password" {
  name  = "/trinity-palette/database/password"
  type  = "SecureString"
  value = module.rds-infra.db_password
  
  tags = {
    Environment = "production" 
    Project     = "trinity-palette"
  }
  
  depends_on = [module.rds-infra]
}

resource "aws_ssm_parameter" "jwt_secret" {
  name  = "/trinity-palette/app/jwt-secret"
  type  = "SecureString"
  value = "your-super-secret-jwt-key-${random_id.jwt_suffix.hex}"
  
  tags = {
    Environment = "production"
    Project     = "trinity-palette"
  }
}

resource "random_id" "jwt_suffix" {
  byte_length = 16
}
