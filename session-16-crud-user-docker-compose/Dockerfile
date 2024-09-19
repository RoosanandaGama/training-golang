# Stage 1: Build stage
FROM golang:1.22.3-alpine AS build

# Install neccessary packages
RUN apk add --no-cache git make protobuf

# Set destination for COPY
WORKDIR /app

# Copy the Go Modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy resource code
COPY . .

# Install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

# RUN the linter with increased timeout
RUN golangci-lint run --timeout 5m

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-crud-user .

# Stage 2: Run stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy binary from build stage
COPY --from=build /app/go-crud-user .

# Install CA certificates
RUN apk --no-cache add ca-certificates

# RUN
CMD [ "./go-crud-user" ]