package postgres

type Config struct {
	Database     string `env:"POSTGRES_DATABASE" default:"hlidskjalf"`
	User         string `env:"POSTGRES_USER" default:"yumiko"`
	Password     string `env:"POSTGRES_PASSWORD" default:"Yumiko1@"`
	Host         string `env:"POSTGRES_HOST" default:"localhost"`
	Port         int    `env:"POSTGRES_PORT" default:"5432"`
	SSL          string `env:"POSTGRES_SSL" default:"verify-full"`
	TimeoutInSec int    `env:"POSTGRES_CONNECTION_TIMEOUT_IN_SEC" default:"60"`
	TimeZone     string `env:"POSTGRES_TIMEZONE" default:"UTC"`
}

func LoadConfig() *Config {
	return &Config{}
}
