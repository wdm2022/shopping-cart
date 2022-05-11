FROM golang:1.18.2-alpine3.15
ARG GO_SWAGGER_URL=https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64

RUN apk add --no-cache curl

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN curl -L -o /usr/local/bin/swagger ${GO_SWAGGER_URL} \
    && chmod +x /usr/local/bin/swagger