package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/stock"
)

var (
	port               = flag.Int("port", 50002, "gRPC server port")
	orderServiceAddr   = flag.String("order-service-addr", "localhost:50000", "address of the order service")
	paymentServiceAddr = flag.String("payment-service-addr", "localhost:50001", "address of the payment service")
)

func main() {
	flag.Parse()

	orderServiceConn := order.Connect(orderServiceAddr)
	paymentServiceConn := stock.Connect(paymentServiceAddr)

	defer func() {
		_ = orderServiceConn.Close()
		_ = paymentServiceConn.Close()
	}()

	if err := stock.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
