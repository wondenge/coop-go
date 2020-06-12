FROM golang:1.14.3-alpine3.11 AS builder

LABEL maintainer="William Ondenge <ondengew@gmail.com>"

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git mercurial 

# Move to working directory /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download all dependencies.
# Dependencies will be cached if the go.mod and go.sum
# files are not changed.
RUN go mod download

# Copy the source code from the current directory to the 
# working directory inside the container.
COPY . .

# Build the server application
RUN go build ./cmd/coop

# Build a small image
FROM scratch

COPY --from=builder /app /

# This container exposes port 8080 to the outside world.
EXPOSE 8000 

# Command to run
CMD ["./coop", "-host", "docker"]