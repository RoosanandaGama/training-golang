# Stage 1: Build stage
FROM golang:1.22.3-alpine AS build

# Set destination for COPY
WORKDIR /app

# Copy resource code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-hello-world .

# Stage 2: Run stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy binary from build stage
COPY --from=build /app/go-hello-world .

# OPTIONAL: expose port
EXPOSE 8080

# RUN
CMD [ "./go-hello-world" ]