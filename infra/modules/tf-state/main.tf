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