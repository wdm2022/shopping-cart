package main

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/stock"
)

var (
	port       = flag.Int("port", 50002, "gRPC server port")
	configFile = flag.String("config-file", "config.yaml", "Path to YAML configuration file")
)

func main() {
	flag.Parse()

	var config stock.Config
	err := cleanenv.ReadConfig(*configFile, &config)
	if err != nil {
		log.Printf("Error when loading config: %v", config)
	}

	orderServiceConn := order.Connect(config.OrderService.Address)
	paymentServiceConn := stock.Connect(config.PaymentService.Address)

	defer func() {
		_ = orderServiceConn.Close()
		_ = paymentServiceConn.Close()
	}()

	if err := stock.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
