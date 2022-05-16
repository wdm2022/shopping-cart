FROM wdm-shoppingcart-builder AS builder

WORKDIR /project
COPY . .

RUN task payment_service:build

FROM alpine:3.15 AS final
EXPOSE 50001

WORKDIR /app
COPY --from=builder /project/dist/payment-service .

ENTRYPOINT ["/app/payment-service"]