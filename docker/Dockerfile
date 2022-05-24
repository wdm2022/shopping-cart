# Builder container image
FROM golang:1.18

RUN apt-get update \
    && apt-get install -y curl protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*
RUN go install github.com/go-task/task/v3/cmd/task@v3.12.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /tmp/go
COPY go.mod go.sum ./
RUN go mod download

WORKDIR /go
