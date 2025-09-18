locals {
  availability_zones = ["us-east-1a", "us-east-1b"]

  vpc_cidr             = "10.0.0.0/16"
  public_subnet_cidrs  = ["10.0.1.0/24", "10.0.2.0/24"]
  private_subnet_cidrs = ["10.0.3.0/24", "10.0.4.0/24"]

  db_identifier = "tp-database"

  ec2_instance_type = "t2.micro"
  ec2_ami_id        = "ami-0c02fb55956c7d316"  # Amazon Linux 2 AMI - updated
  app_port          = 8080
  
  # Docker user data script
  user_data = base64encode(templatefile("${path.module}/user-data.sh", {
    static_bucket = "the-trinity-pallette-static-assets"
    app_port      = 8080
    docker_image  = "your-dockerhub-username/trinity-palette:latest"  # Update this
  }))
}