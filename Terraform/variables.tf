variable "aws_region" {
  description = "The AWS region to deploy resources"
  default     = "us-east-1"
}

variable "instance_type" {
  description = "Type of EC2 instance to use"
  default     = "t2.micro"
}