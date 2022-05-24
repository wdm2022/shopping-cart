FROM local.local/shopping-cart-builder:latest AS builder

WORKDIR /project
COPY . .

RUN task stock_service:build

FROM alpine:3.15 AS final
EXPOSE 50002

WORKDIR /app
COPY --from=builder /project/out/stock-service .

ENTRYPOINT ["/app/stock-service"]