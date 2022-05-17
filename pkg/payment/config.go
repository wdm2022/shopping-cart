package payment

type Config struct {
	OrderService struct {
		Address string `yaml:"address" env:"ORDER_SERVICE_ADDRESS" env-default:"localhost:50000"`
	} `yaml:"order-service"`

	StockService struct {
		Address string `yaml:"address" env:"STOCK_SERVICE_ADDRESS" env-default:"localhost:50002"`
	} `yaml:"stock-service"`
}
