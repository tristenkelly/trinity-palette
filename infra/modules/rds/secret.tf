resource "random_password" "master"{
  length           = 16
  special          = true
  override_special = "_!%^"
}

resource "aws_secretsmanager_secret" "password" {
  name = "db-master-password"
}

resource "aws_secretsmanager_secret_version" "password" {
  secret_id = sadmin
  secret_string = random_password.master.result
}