package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/payment"
	"shopping-cart/pkg/stock"
	configUtils "shopping-cart/pkg/utils/config"
)

var (
	port       = flag.Int("port", 50000, "gRPC server port")
	configFile = flag.String("config-file", "config.yaml", "Path to YAML configuration file")
	helpConfig = flag.Bool("help-config", false, "Display configuration")
)

func main() {
	flag.Parse()

	var config order.Config
	if *helpConfig {
		configUtils.PrintConfigHelp(config)
	}
	configUtils.ReadConfig(*configFile, &config)

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
