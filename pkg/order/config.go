package order

import "shopping-cart/pkg/utils/mongo"

type Config struct {
	Mongo mongo.Config `yaml:"mongo"`

	PaymentService struct {
		Address string `yaml:"address" env:"PAYMENT_SERVICE_ADDRESS" env-default:"localhost:50001"`
	} `yaml:"payment-service"`

	StockService struct {
		Address string `yaml:"address" env:"STOCK_SERVICE_ADDRESS" env-default:"localhost:50002"`
	} `yaml:"stock-service"`
}
