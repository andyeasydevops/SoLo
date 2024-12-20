provider "aws" {
  region = "us-east-1"  # Adjust the region as needed
}

resource "aws_instance" "quote_app" {
  ami           = "ami-0c55b159cbfafe1f0"  # Choose the appropriate AMI ID for your region
  instance_type = "t2.micro"  # Choose an instance size

  # Allow SSH access to the instance
  security_groups = ["default"]

  tags = {
    Name = "QuoteAppInstance"
  }

  # User data script to install Go and run your app
  user_data = <<-EOF
              #!/bin/bash
              sudo apt update -y
              sudo apt install -y golang-go
              git clone https://github.com/yourusername/quote-app.git
              cd quote-app
              go build -o app .
              nohup ./app &
              EOF
}

output "public_ip" {
  value = aws_instance.quote_app.public_ip
}