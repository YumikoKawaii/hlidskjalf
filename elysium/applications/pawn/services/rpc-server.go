package services

import (
	"context"
	pb "elysium.com/pb/pawn"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NewConfig return a optional config with grpc port and http port.
func NewConfig(grpcPort, httpPort int) Config {
	return Config{
		GRPC: Listen{
			Host: "0.0.0.0",
			Port: grpcPort,
		},
		HTTP: Listen{
			Host: "0.0.0.0",
			Port: httpPort,
		},
	}
}

// Config hold http/grpc server config
type Config struct {
	GRPC Listen `json:"grpc" mapstructure:"gprc" yaml:"grpc"`
	HTTP Listen `json:"http" mapstructure:"http" yaml:"grpc"`
}

// Listen config for host/port socket listener
type Listen struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

// String return socket listen DSN
func (l Listen) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

// Server structure
type Server struct {
	gRPC    *grpc.Server
	mux     *runtime.ServeMux
	cfg     Config
	httpMux *mux.Router
}

// NewServer ...
func NewServer(cfg Config, opt ...grpc.ServerOption) *Server {
	return &Server{
		gRPC:    grpc.NewServer(opt...),
		mux:     runtime.NewServeMux(),
		cfg:     cfg,
		httpMux: mux.NewRouter(),
	}
}

func (s *Server) Register(grpcServer ...interface{}) error {
	for _, srv := range grpcServer {
		switch _srv := srv.(type) {
		case pb.GreetServiceServer:
			pb.RegisterGreetServiceServer(s.gRPC, _srv)
			if err := pb.RegisterGreetServiceHandlerFromEndpoint(
				context.Background(), s.mux, s.cfg.GRPC.String(),
				[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
			); err != nil {
				return err
			}

		case pb.PerformanceServiceServer:
			pb.RegisterPerformanceServiceServer(s.gRPC, _srv)
			if err := pb.RegisterPerformanceServiceHandlerFromEndpoint(
				context.Background(), s.mux, s.cfg.GRPC.String(),
				[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
			); err != nil {
				return err
			}

		default:
			return xerrors.Errorf("error unknown GRPC service to register: %#v", srv)
		}
	}
	return nil
}

func (s *Server) Serve() error {

	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	httpMux := http.NewServeMux()
	httpMux.Handle("/", s.mux)
	httpServer := http.Server{
		Addr:    s.cfg.HTTP.String(),
		Handler: httpMux,
	}

	go func() {
		log.Printf("[Pawn] - Serving on %s\n", s.cfg.HTTP)
		if err := httpServer.ListenAndServe(); err != nil {
			errch <- err
		}
	}()

	go func() {
		lis, err := net.Listen("tcp", s.cfg.GRPC.String())
		if err != nil {
			errch <- err
			return
		}
		if err := s.gRPC.Serve(lis); err != nil {
			errch <- err
		}
	}()

	for {
		select {
		case <-stop:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()
			s.gRPC.GracefulStop()
			if err := httpServer.Shutdown(ctx); err != nil {
				log.Printf("failed to stop server: %w", err)
			}
		case err := <-errch:
			return err
		}
	}

}
