package echo

import (
	"context"

	api "github.com/YumikoKawaii/rpc.com/protobuf/echo"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	api.UnimplementedEchoServer
}

func Initialize() *Handler {
	logger.Info("[えこーはんどらー] - はんどらーをしょきかちゅう")
	logger.Info("[えこーはんどらー] - えこーえんどぽいんとをせっていちゅう")
	logger.Info("[えこーはんどらー] - はんどらーじゅんびかんりょう")
	return &Handler{}
}

func (h *Handler) Charge(ctx context.Context, request *api.ChargeRequest) (*api.ChargeResponse, error) {
	logger.Info("[えこーはんどらー] - ちゃーじりくえすとをじゅしん")
	logger.Info("[えこーはんどらー] - ちゃーじりくえすとをしょりちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじでーたをけんしょうちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじしょりをじっこうちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじれすぽんすをさくせいちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじかんりょう")
	return &api.ChargeResponse{}, nil
}

func (h *Handler) Discharge(ctx context.Context, request *api.DischargeRequest) (*api.DischargeResponse, error) {
	logger.Info("[えこーはんどらー] - でぃすちゃーじりくえすとをじゅしん")
	logger.Info("[えこーはんどらー] - でぃすちゃーじりくえすとをしょりちゅう")
	logger.Info("[えこーはんどらー] - でぃすちゃーじでーたをけんしょうちゅう")
	logger.Info("[えこーはんどらー] - でぃすちゃーじしょりをじっこうちゅう")
	logger.Info("[えこーはんどらー] - でぃすちゃーじれすぽんすをさくせいちゅう")
	logger.Info("[えこーはんどらー] - でぃすちゃーじかんりょう")
	return &api.DischargeResponse{}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterEchoServer(instance.Instance(), h)
	return api.RegisterEchoHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
}
