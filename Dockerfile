# Step 1: Build stage
FROM golang:1.23-alpine AS builder

# Set the working directory to /App
WORKDIR /App

# Copy the files into the /App directory
COPY . .

# Run go mod tidy to handle dependencies
RUN go mod tidy

# Build the app (ensure the binary name is correct)
RUN go build -o quote-app main.go

# Check if the 'quote-app' is generated
RUN ls -l /App

# Step 2: Final production image
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /app

# Copy the compiled app from the builder step
COPY --from=builder /App/quote-app .

# Check if the file is present in the final image
RUN ls -l /app

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./quote-app"]