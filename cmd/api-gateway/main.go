package main

import (
	"flag"
	"log"
	"shopping-cart/pkg/api"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/payment"
	"shopping-cart/pkg/stock"
	configUtils "shopping-cart/pkg/utils/config"
)

var (
	port       = flag.Int("port", 8080, "HTTP server port")
	prefork    = flag.Bool("prefork", false, "Spawn multiple listener processes")
	configFile = flag.String("config-file", "config.yaml", "Path to YAML configuration file")
	helpConfig = flag.Bool("help-config", false, "Display configuration")
)

func main() {
	flag.Parse()

	var config api.Config
	if *helpConfig {
		configUtils.PrintConfigHelp(config)
	}
	configUtils.ReadConfig(*configFile, &config)

	orderServiceConn := order.Connect(config.OrderService.Address)
	paymentServiceConn := payment.Connect(config.PaymentService.Address)
	stockServiceConn := stock.Connect(config.StockService.Address)

	defer func() {
		_ = orderServiceConn.Close()
		_ = paymentServiceConn.Close()
		_ = stockServiceConn.Close()
	}()

	if err := api.RunHttpServer(port, prefork); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
