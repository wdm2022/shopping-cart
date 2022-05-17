package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/payment"
	"shopping-cart/pkg/stock"
)

var (
	port               = flag.Int("port", 50000, "gRPC server port")
	paymentServiceAddr = flag.String("payment-service-addr", "localhost:50001", "address of the payment service")
	stockServiceAddr   = flag.String("stock-service-addr", "localhost:50002", "address of the stock service")
)

func main() {
	flag.Parse()

	paymentServiceConn := payment.Connect(paymentServiceAddr)
	stockServiceConn := stock.Connect(stockServiceAddr)

	defer func() {
		_ = paymentServiceConn.Close()
		_ = stockServiceConn.Close()
	}()

	if err := order.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
