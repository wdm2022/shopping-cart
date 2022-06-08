package mongo

type Config struct {
	Hosts            []string `yaml:"hosts" env:"MONGO_HOSTS" env-default:"localhost:27017"`
	DirectConnection bool     `yaml:"directConnection" env:"MONGO_DIRECT_CONNECTION" env-default:"false"`
	Username         string   `yaml:"username" env:"MONGO_USERNAME" env-required:"true"`
	Password         string   `yaml:"password" env:"MONGO_PASSWORD" env-required:"true"`
	Database         string   `yaml:"database" env:"MONGO_DATABASE" env-required:"true"`
}
