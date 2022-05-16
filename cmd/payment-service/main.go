package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/payment"
)

var (
	port = flag.Int("port", 50001, "gRPC server port")
)

func main() {
	flag.Parse()

	if err := payment.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
