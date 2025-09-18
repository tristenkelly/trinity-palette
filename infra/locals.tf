locals {
  availability_zones = ["us-east-1a", "us-east-1b"]

  vpc_cidr             = "10.0.0.0/16"
  public_subnet_cidrs  = ["10.0.1.0/24", "10.0.2.0/24"]
  private_subnet_cidrs = ["10.0.3.0/24", "10.0.4.0/24"]

  db_identifier = "tp-database"

  ec2_instance_type = "t2.micro"
  ec2_ami_id        = "ami-0c55b159cbfafe1f0"
  app_port          = 8080
}