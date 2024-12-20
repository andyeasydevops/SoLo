
Here’s a detailed set of instructions on how to build and run the Go application locally:

Step 1: Install Go

	1.	Download Go:
Go to the official Go website and download the installer for your operating system.
	2.	Install Go:
Follow the instructions in the installer to complete the installation. On macOS, you can also install Go using Homebrew:

brew install go


	3.	Verify the installation:
Open a terminal and type:

go version

You should see the installed version of Go.

Step 2: Clone the Project

	1.	Open the terminal and navigate to the directory where you want to store the project.
For example:

cd ~/Projects


	2.	Clone the repository containing the project code:

git clone <repository-url>

Replace <repository-url> with the actual URL of your GitHub repository.

	3.	Navigate into the project directory:

cd <project-directory>



Step 3: Initialize the Go Module

	1.	Ensure you’re in the project directory. Then initialize a Go module:

go mod init <module-name>

Replace <module-name> with a relevant name (e.g., quotes-app). This will create a go.mod file.

	2.	Download the required dependencies:

go get github.com/PuerkitoBio/goquery



Step 4: Build the Application

	1.	Build the Go application:

go build -o quotes-app main.go

This will create an executable file named quotes-app in your project directory.

Step 5: Run the Application

	1.	Run the application directly:

./quotes-app


	2.	The server will start, and you’ll see an output like:

Scraping page 1: https://quotes.toscrape.com/page/1/
Scraping page 2: https://quotes.toscrape.com/page/2/
Listening on :8080


	3.	Open your browser or use a tool like curl or Postman to test the endpoint:

http://localhost:8080/quotes



Step 6: Verify Output

The application will return a JSON array containing 100 quotes. Each quote will include the text, author, and tags.

Example output:

[
  {
    "text": "The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking.",
    "author": "Albert Einstein",
    "tags": ["change", "deep-thoughts", "thinking", "world"]
  },
  {
    "text": "It is our choices, Harry, that show what we truly are, far more than our abilities.",
    "author": "J.K. Rowling",
    "tags": ["abilities", "choices"]
  }
]

Troubleshooting

	•	Missing dependencies? Run:

go mod tidy

This ensures all required dependencies are downloaded.

	•	Error with ports? Ensure port 8080 is not in use. You can change the port in the code:

log.Fatal(http.ListenAndServe(":8080", nil))

Replace 8080 with a different port (e.g., 9090).

	•	Permission error on macOS? Ensure the binary is executable:

chmod +x quotes-app



This process will help you set up, build, and run the Go application successfully on your local machine.
Automated Deployment of Quote-App on AWS with Terraform and Docker

This project automates the deployment of the Quote-App using Terraform for infrastructure management and Docker for containerizing the application. It also includes AWS CLI configuration to manage the application deployment on AWS.

Prerequisites

	•	Docker installed on your machine.
	•	Terraform installed on your machine.
	•	AWS CLI installed and configured.
	•	A valid AWS account for infrastructure deployment.

1. Docker Configuration for Scalability

This app is containerized using Docker to allow for scalable deployments. The Dockerfile ensures the application is packaged and ready for deployment. Additionally, Docker Compose can be used to scale the number of application instances running.

Dockerfile

Here is the Dockerfile used to containerize the Go application:

# Build step
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app ./main.go

# Final stage for running the app
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/app /app/
CMD ["/app/app"]

This Dockerfile builds the application from the source code and then packages it into a minimal image using the Distroless base image.

Docker Compose for Scalability

To scale the application, you can use Docker Compose. Below is an example docker-compose.yml that deploys multiple instances of the app for load balancing:

version: '3'
services:
  quote-app:
    build: .
    ports:
      - "8080:8080"
    deploy:
      replicas: 3  # This will run 3 instances of the app for scalability

This will allow you to run the app on multiple containers, improving performance and availability.

2. AWS CLI Configuration

To deploy the application on AWS using Terraform, you need to configure AWS CLI with the correct credentials.

Installing AWS CLI

If AWS CLI is not installed, you can follow the installation instructions here: AWS CLI Installation.

Configuring AWS CLI

Run the following command to configure AWS CLI with your AWS credentials:

aws configure

You will be prompted to enter the following details:

	•	AWS Access Key ID: Your AWS access key.
	•	AWS Secret Access Key: Your AWS secret key.
	•	Default region name: The AWS region to deploy your infrastructure (e.g., us-east-1).
	•	Default output format: The output format (e.g., json).

Verifying AWS CLI

To verify the configuration is successful, run:

aws sts get-caller-identity

This command should return the IAM user details if everything is configured correctly.

3. Terraform Configuration

Terraform Files

This project includes several Terraform configuration files to manage AWS infrastructure.

	•	main.tf: Contains the main configuration for AWS resources like EC2 instances.
	•	variables.tf: Declares variables that can be customized for your infrastructure.
	•	outputs.tf: Defines the output values, such as the public IP of the EC2 instance.

Example of main.tf for AWS EC2 Instance

provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "quote_app_instance" {
  ami           = "ami-12345678"  # Replace with a valid AMI ID
  instance_type = "t2.micro"
  key_name      = "your-key"      # Replace with your SSH key name

  tags = {
    Name = "QuoteAppInstance"
  }
}

Example of variables.tf for AWS Configuration

variable "aws_region" {
  description = "The AWS region to deploy the infrastructure"
  default     = "us-east-1"
}

Initializing and Applying Terraform

Run the following commands to initialize Terraform, review the infrastructure plan, and apply the changes to deploy the app on AWS.

	1.	Initialize Terraform:

terraform init


	2.	Review the Plan:

terraform plan


	3.	Apply the Configuration:

terraform apply

Terraform will provision the AWS resources such as EC2 instances, and after successful execution, it will display the IP address of the instance.

Output Example

Once the infrastructure is deployed, you can get the public IP of the EC2 instance by referencing the output:

output "instance_ip" {
  value = aws_instance.quote_app_instance.public_ip
}

4. Running the Application on AWS

Once the infrastructure is created, the Quote-App will be running on an EC2 instance. You can access the app using the public IP address provided by Terraform. Open your browser and navigate to:

http://<EC2_PUBLIC_IP>:8080/quotes

This will display the Quote-App’s interface where users can interact with the application.


