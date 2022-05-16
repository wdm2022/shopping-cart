package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/stock"
)

var (
	port = flag.Int("port", 50002, "gRPC server port")
)

func main() {
	flag.Parse()

	if err := stock.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
