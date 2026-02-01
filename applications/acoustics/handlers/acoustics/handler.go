package acoustics

import (
	"context"

	api "github.com/YumikoKawaii/rpc.com/protobuf/acoustics"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	api.UnimplementedAcousticsServer
}

func Initialize() *Handler {
	logger.Info("[えこーはんどらー] - はんどらーをしょきかちゅう")
	logger.Info("[えこーはんどらー] - えこーえんどぽいんとをせっていちゅう")
	logger.Info("[えこーはんどらー] - はんどらーじゅんびかんりょう")
	return &Handler{}
}

func (h *Handler) Entry(ctx context.Context, request *api.EntryRequest) (*api.EntryResponse, error) {
	logger.Info("[えこーはんどらー] - ちゃーじりくえすとをじゅしん")
	logger.Info("[えこーはんどらー] - ちゃーじりくえすとをしょりちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじでーたをけんしょうちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじしょりをじっこうちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじれすぽんすをさくせいちゅう")
	logger.Info("[えこーはんどらー] - ちゃーじかんりょう")
	return &api.EntryResponse{}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterAcousticsServer(instance.Instance(), h)
	return api.RegisterAcousticsHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
}
