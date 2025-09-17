resource "aws_vpc" "tpVPC" {
    cidr_block = var.vpc_cidr
    instance_tenancy = "default"
    tags = {
        Name = "the-trinity-pallette-vpc"
        Project = "the-trinity-pallette"
    }
}

resource "aws_internet_gateway" "tpIGW" {
    vpc_id = aws_vpc.tpVPC.id
    tags = {
        Name = "the-trinity-pallette-igw"
        Project = "the-trinity-pallette"
    }
}

resource "aws_eip" "tpNatGatewayEIP1" {
  tags = {
    Name    = "tpNatGatewayEIP1"
    Project = "the-trinity-pallette"
  }
}

resource "aws_nat_gateway" "tpNatGateway1" {
  allocation_id = aws_eip.tpNatGatewayEIP1.id
  subnet_id     = aws_subnet.tpPublicSubnet1.id
  tags = {
    Name    = "tpNatGateway1"
    Project = "the-trinity-pallette"
  }
}

resource "aws_subnet" "tpPublicSubnet1" {
  vpc_id            = aws_vpc.tpVPC.id
  cidr_block        = var.public_subnet_cidrs[0]
  availability_zone = var.availability_zones[0]
  tags = {
    Name    = "tpPublicSubnet1"
    Project = "the-trinity-pallette"
  }
}

resource "aws_eip" "tpNatGatewayEIP2" {
  tags = {
    Name    = "tpNatGatewayEIP2"
    Project = "the-trinity-pallette"
  }
}

resource "aws_nat_gateway" "tpNatGateway2" {
  allocation_id = aws_eip.tpNatGatewayEIP2.id
  subnet_id     = aws_subnet.tpPublicSubnet1.id
  tags = {
    Name    = "tpNatGateway2"
    Project = "the-trinity-pallette"
  }
}

resource "aws_subnet" "tpPublicSubnet2" {
  vpc_id            = aws_vpc.tpVPC.id
  cidr_block        = var.public_subnet_cidrs[1]
  availability_zone = var.availability_zones[1]
  tags = {
    Name    = "tpPublicSubnet2"
    Project = "the-trinity-pallette"
  }
}

resource "aws_subnet" "tpPrivateSubnet1" {
  vpc_id            = aws_vpc.tpVPC.id
  cidr_block        = var.private_subnet_cidrs[0]
  availability_zone = var.availability_zones[0]
  tags = {
    Name    = "tpPrivateSubnet1"
    Project = "the-trinity-pallette"
  }
}

resource "aws_subnet" "tpPrivateSubnet2" {
  vpc_id            = aws_vpc.tpVPC.id
  cidr_block        = var.private_subnet_cidrs[1]
  availability_zone = var.availability_zones[1]
  tags = {
    Name    = "tpPrivateSubnet2"
    Project = "the-trinity-pallette"
  }
}

resource "aws_route_table" "tpPublicRT" {
  vpc_id = aws_vpc.tpVPC.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.tpIGW.id
  }
  tags = {
    Name    = "tpPublicRT"
    Project = "the-trinity-pallette"
  }
}

resource "aws_route_table" "tpPrivateRT1" {
  vpc_id = aws_vpc.tpVPC.id
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.tpNatGateway1.id
  }
  tags = {
    Name    = "tpPrivateRT1"
    Project = "the-trinity-pallette"
  }
}

resource "aws_route_table" "tpPrivateRT2" {
  vpc_id = aws_vpc.tpVPC.id
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.tpNatGateway2.id
  }
  tags = {
    Name    = "tpPrivateRT2"
    Project = "the-trinity-pallette"
  }
}

resource "aws_route_table_association" "tpPublicRTassociation1" {
  subnet_id      = aws_subnet.tpPublicSubnet1.id
  route_table_id = aws_route_table.tpPublicRT.id
}

resource "aws_route_table_association" "tpPublicRTassociation2" {
  subnet_id      = aws_subnet.tpPublicSubnet2.id
  route_table_id = aws_route_table.tpPublicRT.id
}

resource "aws_route_table_association" "tpPrivateRTassociation1" {
  subnet_id      = aws_subnet.tpPrivateSubnet1.id
  route_table_id = aws_route_table.tpPrivateRT1.id
}

resource "aws_route_table_association" "tpPrivateRTassociation2" {
  subnet_id      = aws_subnet.tpPrivateSubnet2.id
  route_table_id = aws_route_table.tpPrivateRT2.id
}