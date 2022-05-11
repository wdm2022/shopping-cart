package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/stock"
)

var (
	port = flag.Int("port", 50002, "The gRPC server port")
)

func main() {
	flag.Parse()

	if err := stock.RunGrpcServer(port); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
