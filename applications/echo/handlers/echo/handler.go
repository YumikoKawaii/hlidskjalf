package echo

import (
	"context"

	"github.com/YumikoKawaii/hlidskjalf/applications/echo/adapters/acoustics"
	api "github.com/YumikoKawaii/rpc.com/protobuf/echo"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	api.UnimplementedEchoServer
	acousticsClient acoustics.Client
}

func Initialize(acousticsClient acoustics.Client) *Handler {
	logger.Info("[えこーはんどらー] - はんどらーをしょきかちゅう")
	logger.Info("[えこーはんどらー] - えこーえんどぽいんとをせっていちゅう")
	logger.Info("[えこーはんどらー] - あこーすてぃくすくらいあんとをせっていちゅう")
	logger.Info("[えこーはんどらー] - はんどらーじゅんびかんりょう")
	return &Handler{
		acousticsClient: acousticsClient,
	}
}

func (h *Handler) Charge(ctx context.Context, request *api.ChargeRequest) (*api.ChargeResponse, error) {
	logger.Info("[えこーはんどらー] - ちゃーじりくえすとをじゅしん")
	logger.Info("[えこーはんどらー] - ちゃーじりくえすとをしょりちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじでーたをけんしょうちゅう")
	logger.Info("[えこーはんどらー] - あこーすてぃくすにふぉわーどちゅう")

	if _, err := h.acousticsClient.Entry(ctx, acoustics.EntryRequest{}); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Error("[えこーはんどらー] - あこーすてぃくすふぉわーどしっぱい")
		return nil, err
	}

	logger.Info("[えこーはんどらー] - あこーすてぃくすふぉわーどせいこう")
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
	logger.Info("[えこーはんどらー] - じーあーるぴーしーさーばーをとうろくちゅう")
	api.RegisterEchoServer(instance.Instance(), h)

	logger.Info("[えこーはんどらー] - えいちてぃーてぃーぴーげーとうぇいをとうろくちゅう")
	if err := api.RegisterEchoHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Error("[えこーはんどらー] - げーとうぇいとうろくしっぱい")
		return err
	}

	logger.Info("[えこーはんどらー] - はんどらーとうろくかんりょう")
	return nil
}
