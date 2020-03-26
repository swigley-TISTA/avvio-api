# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Sam Wigley <swigley@tistatech.com>"

# Set the Current Working Directory inside the container
ENV GOPATH /go

RUN export GOPATH=/go
RUN echo $GOPATH
WORKDIR /go

COPY . /go/src/avvio-api

WORKDIR /go/src/avvio-api

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
COPY Gopkg.toml Gopkg.lock ./

RUN pwd
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN dep ensure --vendor-only

# Copy the source from the current directory to the Working Directory inside the container

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
