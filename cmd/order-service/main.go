package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/order"
)

var (
	port = flag.Int("port", 50000, "gRPC server port")
)

func main() {
	flag.Parse()

	if err := order.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
