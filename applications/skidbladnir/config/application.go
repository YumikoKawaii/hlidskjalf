package config

import (
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type SocketBinding struct {
	Endpoint   string `json:"endpoint" mapstructure:"endpoint" yaml:"endpoint"`
	TargetPort int    `json:"target_port" mapstructure:"target_port" yaml:"target_port"`
}

type OutboundConfig struct {
	Port int `json:"port" mapstructure:"port" yaml:"port"` // iptables OUTPUT redirect target port
}

type InboundConfig struct {
	GRPC SocketBinding `json:"grpc" mapstructure:"grpc" yaml:"grpc"` // inbound gRPC listener (redirected from app's gRPC port)
	HTTP SocketBinding `json:"http" mapstructure:"http" yaml:"http"` // inbound HTTP listener (redirected from app's HTTP port)
}

type RateLimitConfig struct {
	Enabled bool    `json:"enabled" mapstructure:"enabled" yaml:"enabled"`
	RPS     float64 `json:"rps" mapstructure:"rps" yaml:"rps"`
	Burst   int     `json:"burst" mapstructure:"burst" yaml:"burst"`
}

type PeerConfig struct {
	ServiceName string `json:"service_name" mapstructure:"service_name" yaml:"service_name"`
	PodIP       string `json:"pod_ip" mapstructure:"pod_ip" yaml:"pod_ip"`
	Namespace   string `json:"namespace" mapstructure:"namespace" yaml:"namespace"`
	LeaderPort  int    `json:"leader_port" mapstructure:"leader_port" yaml:"leader_port"` // gRPC port to reach the leader's RateLimiter service
}

type Application struct {
	Server       *server.Config        `json:"skidbladnir_server" mapstructure:"skidbladnir_server" yaml:"skidbladnir_server"`
	Outbound     *OutboundConfig       `json:"outbound" mapstructure:"outbound" yaml:"outbound"`
	Inbound      *InboundConfig        `json:"inbound" mapstructure:"inbound" yaml:"inbound"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	RateLimit    *RateLimitConfig      `json:"rate_limit" mapstructure:"rate_limit" yaml:"rate_limit"`
	Peer         *PeerConfig           `json:"peer" mapstructure:"peer" yaml:"peer"`
}

func loadDefault() *Application {
	return &Application{
		Server: &server.Config{
			HTTP: "0.0.0.0:15001",
			GRPC: "0.0.0.0:15443",
		},
		Outbound: &OutboundConfig{
			Port: 15001,
		},
		Inbound: &InboundConfig{
			GRPC: SocketBinding{
				Endpoint:   "0.0.0.0:15002",
				TargetPort: 10443,
			},
			HTTP: SocketBinding{
				Endpoint:   "0.0.0.0:15003",
				TargetPort: 10080,
			},
		},
		TracerConfig: tracer.DefaultConfig(),
		RateLimit:    &RateLimitConfig{Enabled: false, RPS: 100, Burst: 100},
		Peer:         &PeerConfig{LeaderPort: 15443},
	}
}
