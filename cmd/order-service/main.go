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
	port       = flag.Int("port", 50000, "gRPC server port")
	configFile = flag.String("config-file", "config.yaml", "Path to YAML configuration file")
)

func main() {
	flag.Parse()

	var config order.Config
	err := cleanenv.ReadConfig(*configFile, &config)
	if err != nil {
		log.Printf("Error when loading config: %v", config)
	}

	paymentServiceConn := payment.Connect(config.PaymentService.Address)
	stockServiceConn := stock.Connect(config.StockService.Address)

	defer func() {
		_ = paymentServiceConn.Close()
		_ = stockServiceConn.Close()
	}()

	if err := order.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
