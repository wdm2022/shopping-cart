package main

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/payment"
	"shopping-cart/pkg/stock"
)

var (
	port       = flag.Int("port", 50001, "gRPC server port")
	configFile = flag.String("config-file", "config.yaml", "Path to YAML configuration file")
)

func main() {
	flag.Parse()

	var config payment.Config
	err := cleanenv.ReadConfig(*configFile, &config)
	if err != nil {
		log.Printf("Error when loading config: %v", config)
	}

	orderServiceConn := order.Connect(config.OrderService.Address)
	stockServiceConn := stock.Connect(config.StockService.Address)

	defer func() {
		_ = orderServiceConn.Close()
		_ = stockServiceConn.Close()
	}()

	if err := payment.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
