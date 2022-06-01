###########
# BUILDER #
###########

FROM golang:1.18 AS builder

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

###################
# BUILDING STAGES #
###################

FROM builder AS api-gateway-build

WORKDIR /project
COPY . .

RUN task api_gateway:build

FROM builder AS order-service-build

WORKDIR /project
COPY . .

RUN task order_service:build

FROM builder AS payment-service-build

WORKDIR /project
COPY . .

RUN task payment_service:build

FROM builder AS stock-service-build

WORKDIR /project
COPY . .

RUN task stock_service:build

################
# FINAL STAGES #
################

FROM alpine:3.15 AS api-gateway
EXPOSE 8080

RUN apk add --no-cache dumb-init

WORKDIR /app
COPY --from=api-gateway-build /project/out/api-gateway .

# Dumb init necessary to claim the PID 1 and allow Fiber to manage the process IDs
# See: https://github.com/gofiber/fiber/issues/1036
ENTRYPOINT ["/usr/bin/dumb-init", "--", "sh", "-c", "/app/api-gateway --prefork"]

FROM alpine:3.15 AS order-service
EXPOSE 50000

WORKDIR /app
COPY --from=order-service-build /project/out/order-service .

ENTRYPOINT ["/app/order-service"]

FROM alpine:3.15 AS payment-service
EXPOSE 50001

WORKDIR /app
COPY --from=payment-service-build /project/out/payment-service .

ENTRYPOINT ["/app/payment-service"]

FROM alpine:3.15 AS stock-service
EXPOSE 50002

WORKDIR /app
COPY --from=stock-service-build /project/out/stock-service .

ENTRYPOINT ["/app/stock-service"]
