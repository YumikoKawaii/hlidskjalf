package mimic

import "context"

type Client interface {
	Mimic(ctx context.Context) (Response, error)
}

func NewClient(host string, isUseGrpc bool) Client {
	if isUseGrpc {
		return NewRpcClient(host)
	}

	return NewHttpClient(host)
}

type Primary struct {
	First       uint32      `json:"first,omitempty"`
	Second      uint32      `json:"second,omitempty"`
	Third       uint32      `json:"third,omitempty"`
	Fourth      uint32      `json:"fourth,omitempty"`
	Fifth       uint32      `json:"fifth,omitempty"`
	Sixth       uint32      `json:"sixth,omitempty"`
	Seventh     uint32      `json:"seventh,omitempty"`
	Eighth      uint32      `json:"eighth,omitempty"`
	Ninth       uint32      `json:"ninth,omitempty"`
	Tenth       uint32      `json:"tenth,omitempty"`
	Secondaries []Secondary `json:"secondaries,omitempty"`
}

type Secondary struct {
	First      uint64     `json:"first,omitempty"`
	Second     uint64     `json:"second,omitempty"`
	Third      uint64     `json:"third,omitempty"`
	Fourth     uint64     `json:"fourth,omitempty"`
	Fifth      uint64     `json:"fifth,omitempty"`
	Tertiaries []Tertiary `json:"tertiaries,omitempty"`
}

type Tertiary struct {
	First  float64 `json:"first,omitempty"`
	Second float64 `json:"second,omitempty"`
	Third  float64 `json:"third,omitempty"`
	Fourth float64 `json:"fourth,omitempty"`
	Fifth  float64 `json:"fifth,omitempty"`
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
