output "tp_vpc_id" {
    description = "VPC id"
    value = aws_vpc.tpVPC.id
}

output "tp_public_subnet_ids" {
    description = "List of public subnet ids"
    value = [
        aws_subnet.tpPublicSubnet1.id,
        aws_subnet.tpPublicSubnet2.id
        ]
}

output "tp_private_subnet_ids" {
    description = "List of private subnet ids"
    value = [
        aws_subnet.tpPrivateSubnet1.id,
        aws_subnet.tpPrivateSubnet2.id
        ]
}