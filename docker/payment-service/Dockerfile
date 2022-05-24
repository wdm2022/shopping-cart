FROM local.local/shopping-cart-builder:latest AS builder

WORKDIR /project
COPY . .

RUN task payment_service:build

FROM alpine:3.15 AS final
EXPOSE 50001

WORKDIR /app
COPY --from=builder /project/out/payment-service .

ENTRYPOINT ["/app/payment-service"]