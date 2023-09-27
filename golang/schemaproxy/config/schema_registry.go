package config

type SchemaRegistryConfig struct {
	URL string `name:"schema-registry-url" env:"SCHEMA_REGISTRY_URL" default:""`
}
