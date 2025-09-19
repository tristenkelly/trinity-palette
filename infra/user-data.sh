#!/bin/bash
set -e

# Update system
yum update -y

# Install Docker
yum install -y docker
systemctl start docker
systemctl enable docker
usermod -a -G docker ec2-user

# Install AWS CLI v2
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
yum install -y unzip
unzip awscliv2.zip
./aws/install
rm -rf aws awscliv2.zip

# Install Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Create app directory
mkdir -p /opt/trinity-palette
cd /opt/trinity-palette
chown -R ec2-user:ec2-user /opt/trinity-palette

# Create docker-compose.yml
cat > docker-compose.yml << 'EOF'
version: '3.8'
services:
  app:
    image: ${docker_image}
    ports:
      - "${app_port}:${app_port}"
    environment:
      - DATABASE_URL=$${DATABASE_URL}
      - STATIC_BUCKET=${static_bucket}
      - AWS_REGION=us-east-1
      - STATIC_BASE_URL=https://${static_bucket}.s3.amazonaws.com
      - APP_PORT=${app_port}
      - APP_ENV=production
      - JWT_SECRET=$${JWT_SECRET}
    restart: unless-stopped
    logging:
      driver: "json-file"
      options:
        max-size: "100m"
        max-file: "5"
EOF

# Create environment file with auto-fetched credentials
echo "Fetching database credentials from SSM Parameter Store..."
DB_ENDPOINT=$(aws ssm get-parameter --name "/trinity-palette/database/endpoint" --query Parameter.Value --output text --region us-east-1)
DB_PASSWORD=$(aws ssm get-parameter --name "/trinity-palette/database/password" --with-decryption --query Parameter.Value --output text --region us-east-1)
JWT_SECRET=$(aws ssm get-parameter --name "/trinity-palette/app/jwt-secret" --with-decryption --query Parameter.Value --output text --region us-east-1)

cat > .env << EOF
# Database configuration - automatically configured from SSM
DATABASE_URL=postgresql://sadmin:$DB_PASSWORD@$DB_ENDPOINT:5432/postgres?sslmode=require

# JWT Secret - automatically generated
JWT_SECRET=$JWT_SECRET

# S3 Configuration (automatically configured)
STATIC_BUCKET=${static_bucket}
AWS_REGION=us-east-1
STATIC_BASE_URL=https://${static_bucket}.s3.amazonaws.com

# App Configuration
APP_PORT=${app_port}
APP_ENV=production
EOF

# Fix ownership and permissions
chown ec2-user:ec2-user .env
chmod 644 .env
chown -R ec2-user:ec2-user /opt/trinity-palette

# Create systemd service
cat > /etc/systemd/system/trinity-palette.service << 'EOF'
[Unit]
Description=Trinity Palette Application
After=docker.service
Requires=docker.service

[Service]
Type=oneshot
RemainAfterExit=yes
WorkingDirectory=/opt/trinity-palette
ExecStart=/usr/local/bin/docker-compose up -d
ExecStop=/usr/local/bin/docker-compose down
TimeoutStartSec=0

[Install]
WantedBy=multi-user.target
EOF

# Enable and start the service automatically
systemctl daemon-reload
systemctl enable trinity-palette.service
systemctl start trinity-palette.service

# Create deployment script for updates
cat > /opt/trinity-palette/deploy.sh << 'EOF'
#!/bin/bash
cd /opt/trinity-palette
docker-compose pull
docker-compose down
docker-compose up -d
docker image prune -f
EOF

chmod +x /opt/trinity-palette/deploy.sh
chown ec2-user:ec2-user /opt/trinity-palette/deploy.sh

# Log completion
echo "Trinity Palette setup completed at $(date)" >> /var/log/trinity-palette-setup.log
echo "Application should be running on port ${app_port}" >> /var/log/trinity-palette-setup.log
echo "Database automatically configured from SSM Parameter Store" >> /var/log/trinity-palette-setup.log

# Start the service automatically
systemctl start trinity-palette.service
