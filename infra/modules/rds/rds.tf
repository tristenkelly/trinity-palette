data "aws_secretsmanager_secret" "password" {
  name = "db-master-password"
}

data "aws_secretsmanager_secret_version" "password" {
  secret_id = data.aws_secretsmanager_secret.password.id
}

resource "aws_db_instance" "default" {
  identifier           = var.db_identifier
  allocated_storage    = 20
  storage_type         = "gp2"
  engine               = "postgres"
  engine_version       = "13.22" 
  instance_class       = "db.t3.micro" 
  db_subnet_group_name = aws_db_subnet_group.db_subnet_group.name
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  username             = "sadmin"
  password             = data.aws_secretsmanager_secret_version.password.secret_string
  skip_final_snapshot  = true  
}

resource "aws_db_subnet_group" "db_subnet_group" {
  name       = "tp-db-subnet-group"
  subnet_ids = var.private_subnet_ids
  tags = {
    Name    = "tp-db-subnet-group"
    Project = "the-trinity-pallette"
  }
}

resource "aws_security_group" "rds_sg" {
  name        = "tp-rds-sg"
  description = "Security group for RDS instance"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"]  # VPC CIDR block
  }

  tags = {
    Name    = "tp-rds-sg"
    Project = "the-trinity-pallette"
  }
}