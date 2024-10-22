package service

import (
	"context"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/center"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
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

// Controller handles HTTP request.
// This should only be used for HTTP-only request.
type Controller interface {
	ApplyRoutes(*mux.Router)
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

func httpHeaderMatcher(key string) (string, bool) {
	if strings.ToLower(key) == "x-api-key" {
		return key, true
	}
	return runtime.DefaultHeaderMatcher(key)
}

// NewServer ...
func NewServer(cfg Config, opt ...grpc.ServerOption) *Server {
	return &Server{
		gRPC: grpc.NewServer(opt...),
		mux: runtime.NewServeMux(
			utils.StandardisedProtoJSONMarshaler(),
			runtime.WithIncomingHeaderMatcher(httpHeaderMatcher),
		),
		cfg:     cfg,
		httpMux: mux.NewRouter(),
	}
}

func (s *Server) Register(grpcServer ...interface{}) error {
	for _, srv := range grpcServer {
		switch _srv := srv.(type) {
		case pb.CenterServiceServer:
			pb.RegisterCenterServiceServer(s.gRPC, _srv)
			if err := pb.RegisterCenterServiceHandlerFromEndpoint(context.Background(), s.mux, s.cfg.GRPC.String(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
				return err
			}
		default:
			return xerrors.Errorf("error unknown GRPC service to register: %#v", srv)
		}
	}
	return nil
}

func isRunningInDockerContainer() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	logrus.Infof("preflight request for %s", r.URL.Path)
}

// Serve server listen for HTTP and GRPC
func (s *Server) Serve() error {
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s.httpMux.PathPrefix("/").Handler(s.mux)
	httpServer := http.Server{
		Addr:    s.cfg.HTTP.String(),
		Handler: allowCORS(s.httpMux),
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			errch <- err
		}
	}()
	go func() {
		listener, err := net.Listen("tcp", s.cfg.GRPC.String())
		if err != nil {
			errch <- err
			return
		}
		if err := s.gRPC.Serve(listener); err != nil {
			errch <- fmt.Errorf("error serve gRPC server: %w", err)
		}
	}()
	for {
		select {
		case <-stop:
			ctx, cancelFn := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancelFn()
			s.gRPC.GracefulStop()
			if err := httpServer.Shutdown(ctx); err != nil {
				logrus.Errorf("failed to stop server: %w", err)
			}
			if !isRunningInDockerContainer() {
				fmt.Println("Shutting down. Wait for 5 seconds")
				time.Sleep(5 * time.Second)
			}
			return nil
		case err := <-errch:
			return err
		}
	}
}
