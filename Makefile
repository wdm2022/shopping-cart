all: build

dev_prepare: vendor generate

# Building

build: build_stock_service build_order_service build_payment_service

build_stock_service: dev_prepare
	go build -o build/stockService shoppingcart/cmd/stockService

build_order_service: dev_prepare
	go build -o build/orderService shoppingcart/cmd/orderService

build_payment_service: dev_prepare
	go build -o build/paymentService shoppingcart/cmd/paymentService

# Development

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