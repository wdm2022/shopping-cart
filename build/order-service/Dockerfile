FROM wdm-shoppingcart-builder AS builder

WORKDIR /project
COPY . .

RUN task order_service:build

FROM alpine:3.15 AS final
EXPOSE 50000

WORKDIR /app
COPY --from=builder /project/dist/order-service .

ENTRYPOINT ["/app/order-service"]