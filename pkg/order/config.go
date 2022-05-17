package order

type Config struct {
	PaymentService struct {
		Address string `yaml:"address" env:"PAYMENT_SERVICE_ADDRESS" env-default:"localhost:50001"`
	} `yaml:"payment-service"`

	StockService struct {
		Address string `yaml:"address" env:"STOCK_SERVICE_ADDRESS" env-default:"localhost:50002"`
	} `yaml:"stock-service"`
}
