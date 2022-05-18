package main

import (
	"context"
	"flag"
	"log"
	"shopping-cart/pkg/order"
	"shopping-cart/pkg/stock"
	"shopping-cart/pkg/stock/mongo"
	configUtils "shopping-cart/pkg/utils/config"
)

var (
	port       = flag.Int("port", 50002, "gRPC server port")
	configFile = flag.String("config-file", "dev/stock-service/config.yaml", "Path to YAML configuration file")
	helpConfig = flag.Bool("help-config", false, "Display configuration")
)

func main() {
	flag.Parse()

	var config stock.Config
	if *helpConfig {
		configUtils.PrintConfigHelp(config)
	}
	configUtils.ReadConfig(*configFile, &config)

	orderServiceConn := order.Connect(config.OrderService.Address)
	paymentServiceConn := stock.Connect(config.PaymentService.Address)

	mongoClient := mongo.Connect(&config.Mongo)

	defer func() {
		_ = orderServiceConn.Close()
		_ = paymentServiceConn.Close()
		_ = mongoClient.Disconnect(context.Background())
	}()

	if err := stock.RunGrpcServer(port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
