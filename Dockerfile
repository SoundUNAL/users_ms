# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /app

# Download Go modules

COPY go.mod go.sum ./
RUN go mod download

# Copy the local package files to the container's workspace
COPY . /app


# Build the Go app
#RUN go build -o /app/cmd/main 

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
# CMD ["./app/cmd/main"]
