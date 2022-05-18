package stock

import "shopping-cart/pkg/utils/mongo"

type Config struct {
	Mongo mongo.Config `yaml:"mongo"`

	OrderService struct {
		Address string `yaml:"address" env:"ORDER_SERVICE_ADDRESS" env-default:"localhost:50000"`
	} `yaml:"order-service"`

	PaymentService struct {
		Address string `yaml:"address" env:"PAYMENT_SERVICE_ADDRESS" env-default:"localhost:50001"`
	} `yaml:"payment-service"`
}
