package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/order"
)

var (
	port = flag.Int("port", 50000, "The gRPC server port")
)

func main() {
	flag.Parse()

	if err := order.RunGrpcServer(port); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
