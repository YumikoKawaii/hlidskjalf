package clickhouse

type Config struct {
	Addresses string `env:"CLICKHOUSE_ADDRESSES" default:"localhost:9124"`
	Database  string `env:"CLICKHOUSE_DATABASE" default:"pawn"`
	Username  string `env:"CLICKHOUSE_USERNAME" default:"yumiko"`
	Password  string `env:"CLICKHOUSE_PASSWORD" default:"Yumiko1@"`
	Cluster   string `env:"CLICKHOUSE_CLUSTER" default:"hlidskjalf"`
}
