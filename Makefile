all: build

dev_prepare: vendor generate_protobuf

# Building

build: build_stock_service build_order_service build_payment_service

build_stock_service: dev_prepare
	go build -o build/stockService shoppingCart/cmd/stockService

build_order_service: dev_prepare
	go build -o build/orderService shoppingCart/cmd/orderService

build_payment_service: dev_prepare
	go build -o build/paymentService shoppingCart/cmd/paymentService

# Development

vendor:
	go mod vendor

generate_protobuf:
	protoc --go_out=. \
		--go-grpc_out=. \
		api/proto/*.proto

# Cleanup

clean:
	rm -rf build

clean_all: clean
	rm -rf vendor