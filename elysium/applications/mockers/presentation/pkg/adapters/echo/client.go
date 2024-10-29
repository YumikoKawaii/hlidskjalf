package echo

import "context"

type Client interface {
	Echo(ctx context.Context) (Response, error)
}

func NewClient(host string, isUseGrpc bool) Client {
	if isUseGrpc {
		return NewRpcClient(host)
	}

	return NewHttpClient(host)
}

type Stat struct {
	First   string `json:"first,omitempty"`
	Second  string `json:"second,omitempty"`
	Third   string `json:"third,omitempty"`
	Fourth  string `json:"fourth,omitempty"`
	Fifth   string `json:"fifth,omitempty"`
	Sixth   string `json:"sixth,omitempty"`
	Seventh string `json:"seventh,omitempty"`
	Eighth  string `json:"eighth,omitempty"`
	Ninth   string `json:"ninth,omitempty"`
	Tenth   string `json:"tenth,omitempty"`
}

type Response struct {
	Code    uint32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Stats []Stat `json:"stats,omitempty"`
	} `json:"data"`
}
