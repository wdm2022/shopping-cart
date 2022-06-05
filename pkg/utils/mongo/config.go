package mongo

type Config struct {
	Host             string `yaml:"host" env:"MONGO_HOST" env-default:"localhost"`
	Port             int    `yaml:"port" env:"MONGO_PORT" env-default:"27017"`
	DirectConnection bool   `yaml:"directConnection" env:"MONGO_DIRECT_CONNECTION" env-default:"false"`
	Username         string `yaml:"username" env:"MONGO_USERNAME" env-required:"true"`
	Password         string `yaml:"password" env:"MONGO_PASSWORD" env-required:"true"`
	Database         string `yaml:"database" env:"MONGO_DATABASE" env-required:"true"`
}
