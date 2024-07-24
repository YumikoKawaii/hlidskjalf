package types

type Prototype interface {
	// clickhouse related settings
	ToClickhouseTableName() string
	GetOrderKey() string
	GetPartitionKey() string
	GetShardingKey() string
	GetClickhouseEngine() string
	GetClickhouseTTL() string
	// mysql related settings
	ToMySQLTableName() string
	// postgres related settings
	ToPostgresTableName() string

	// serialize data
	ToBytes() []byte
	Analyze([]byte) (Prototype, error)
	GetSubject() string
}

type Primordial struct {
	Subject string `json:"subject"`
}
