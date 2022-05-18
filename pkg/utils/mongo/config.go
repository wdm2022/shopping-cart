package mongo

type Config struct {
	Host     string `yaml:"host" env:"MONGO_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"MONGO_PORT" env-default:"27017"`
	Username string `yaml:"username" env-required:"MONGO_USERNAME"`
	Password string `yaml:"password" env-required:"MONGO_PASSWORD"`
	Database string `yaml:"database" env-required:"MONGO_DATABASE"`
}
