package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/payment"
)

var (
	port = flag.Int("port", 50001, "The gRPC server port")
)

func main() {
	flag.Parse()

	if err := payment.RunGrpcServer(port); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
