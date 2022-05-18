FROM wdm-shoppingcart-builder AS builder

WORKDIR /project
COPY . .

RUN task api_gateway:build

FROM alpine:3.15 AS final
EXPOSE 8080

RUN apk add --no-cache dumb-init

WORKDIR /app
COPY --from=builder /project/out/api-gateway .

# Dumb init necessary to claim the PID 1 and allow Fiber to manage the process IDs
# See: https://github.com/gofiber/fiber/issues/1036
ENTRYPOINT ["/usr/bin/dumb-init", "--", "sh", "-c", "/app/api-gateway --prefork"]