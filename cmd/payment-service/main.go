package main

import (
	"context"
	"flag"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/payment"
	"shopping-cart/pkg/payment/mongo"
	"shopping-cart/pkg/stock"
	configUtils "shopping-cart/pkg/utils/config"
)

var (
	port       = flag.Int("port", 50001, "gRPC server port")
	configFile = flag.String("config-file", "dev/payment-service/config.yaml", "Path to YAML configuration file")
	helpConfig = flag.Bool("help-config", false, "Display configuration")
)

func main() {
	flag.Parse()

	var config payment.Config
	if *helpConfig {
		configUtils.PrintConfigHelp(config)
	}
	configUtils.ReadConfig(*configFile, &config)

	orderServiceConn := order.Connect(config.OrderService.Address)
	stockServiceConn := stock.Connect(config.StockService.Address)

	mongoClient := mongo.Connect(&config.Mongo)

	defer func() {
		_ = orderServiceConn.Close()
		_ = stockServiceConn.Close()
		_ = mongoClient.Disconnect(context.Background())
	}()

	if err := payment.RunGrpcServer(&mongoClient, port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
