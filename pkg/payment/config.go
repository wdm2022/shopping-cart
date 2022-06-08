package payment

import "shopping-cart/pkg/utils/mongo"

type Config struct {
	Mongo mongo.Config `yaml:"mongo"`

	OrderService struct {
		Address string `yaml:"address" env:"ORDER_SERVICE_ADDRESS" env-default:"localhost:50000"`
	} `yaml:"order-service"`

	StockService struct {
		Address string `yaml:"address" env:"STOCK_SERVICE_ADDRESS" env-default:"localhost:50002"`
	} `yaml:"stock-service"`
}
