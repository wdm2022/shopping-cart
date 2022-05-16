package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/api"
)

var (
	port               = flag.Int("port", 8080, "HTTP server port")
	prefork            = flag.Bool("prefork", false, "Spawn multiple listener processes")
	orderServiceAddr   = flag.String("order-service-addr", "localhost:50000", "address of the order service")
	paymentServiceAddr = flag.String("payment-service-addr", "localhost:50001", "address of the payment service")
	stockServiceAddr   = flag.String("stock-service-addr", "localhost:50002", "address of the stock service")
)

func main() {
	flag.Parse()

	if err := api.RunHttpServer(port, prefork); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
