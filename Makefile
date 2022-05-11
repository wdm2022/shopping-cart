.PHONY: vendor build

all: build

# Building

build: build_stock_service build_order_service build_payment_service build_api_gateway

build_stock_service: dev_prepare
	go build -o build/stock-service shopping-cart/cmd/stock-service

build_order_service: dev_prepare
	go build -o build/order-service shopping-cart/cmd/order-service

build_payment_service: dev_prepare
	go build -o build/payment-service shopping-cart/cmd/payment-service

build_api_gateway: dev_prepare
	go build -o build/api-gateway shopping-cart/cmd/api-gateway

# Running

run_stock_service:
	go run shopping-cart/cmd/stock-service

run_order_service:
	go run shopping-cart/cmd/order-service

run_payment_service:
	go run shopping-cart/cmd/payment-service

run_api_gateway:
	go run shopping-cart/cmd/api-gateway

# Development

dev_prepare: vendor generate

vendor:
	go mod vendor

generate:
	go generate ./...

# Cleanup

clean:
	rm -rf build

clean_protobuf:
	find api/proto -type f -name "*.pb.go" -delete

clean_all: clean clean_protobuf
	rm -rf vendor