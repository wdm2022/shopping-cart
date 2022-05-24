FROM local.local/shopping-cart-builder:latest AS builder

WORKDIR /project
COPY . .

RUN task order_service:build

FROM alpine:3.15 AS final
EXPOSE 50000

WORKDIR /app
COPY --from=builder /project/out/order-service .

ENTRYPOINT ["/app/order-service"]