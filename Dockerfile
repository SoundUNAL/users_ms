
# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy

# Build
RUN  go build -o /userMS

EXPOSE 8080

# Run
CMD ["./userMS"]

#docker build -t users_ms .
#docker run -d -t -i -p 8080:8080 --name users_ms users_ms



