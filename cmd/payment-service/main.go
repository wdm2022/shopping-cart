package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/payment"
	"shopping-cart/pkg/stock"
)

var (
	port             = flag.Int("port", 50001, "gRPC server port")
	orderServiceAddr = flag.String("order-service-addr", "localhost:50000", "address of the order service")
	stockServiceAddr = flag.String("stock-service-addr", "localhost:50002", "address of the stock service")
)

func main() {
	flag.Parse()

	orderServiceConn := order.Connect(orderServiceAddr)
	stockServiceConn := stock.Connect(stockServiceAddr)

	defer func() {
		_ = orderServiceConn.Close()
		_ = stockServiceConn.Close()
	}()

	if err := payment.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
