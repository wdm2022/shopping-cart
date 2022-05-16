FROM wdm-shoppingcart-builder AS builder

WORKDIR /project
COPY . .

RUN task api_gateway:build

FROM alpine:3.15 AS final
EXPOSE 8080

WORKDIR /app
COPY --from=builder /project/dist/api-gateway .

ENTRYPOINT ["/app/api-gateway"]