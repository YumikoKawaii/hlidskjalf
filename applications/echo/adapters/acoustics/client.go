package acoustics

import "context"

type Config struct {
	Endpoint string `json:"endpoint" mapstructure:"endpoint" yaml:"endpoint"`
	Protocol string `json:"protocol" mapstructure:"protocol" yaml:"protocol"`
}

type EntryRequest struct{}

type EntryResponse struct{}

type Client interface {
	Entry(ctx context.Context, request EntryRequest) (EntryResponse, error)
}
