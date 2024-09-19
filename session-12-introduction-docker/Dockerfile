FROM golang:1.22.3-alpine

# Set destination for COPY
WORKDIR /app

# Download GO modules
COPY go.mod ./
RUN go mod download

# Copy resource code
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-hello-world

# OPTIONAL: expose port
EXPOSE 8080

# RUN
CMD [ "./go-hello-world" ]