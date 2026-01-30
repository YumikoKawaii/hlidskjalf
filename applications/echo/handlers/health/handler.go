package health

import (
	"context"
	"time"

	api "github.com/YumikoKawaii/rpc.com/protobuf/health"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	api.UnimplementedHealthServer
}

func Initialize() *Handler {
	logger.Info("[へるす] - はんどらーをしょきかちゅう")
	logger.Info("[へるす] - へるすちぇっくえんどぽいんとをせっていちゅう")
	logger.Info("[へるす] - はんどらーじゅんびかんりょう")
	return &Handler{}
}

func (h *Handler) Liveness(ctx context.Context, request *api.LivenessRequest) (*api.LivenessResponse, error) {
	logger.Info("[へるす] - らいぶねすちぇっくをじゅしん")
	logger.Info("[へるす] - らいぶねすりくえすとをしょりちゅう")
	logger.Info("[へるす] - さーびすのじょうたいをかくにんちゅう")
	logger.Info("[へるす] - さーびすはいきています")
	logger.Info("[へるす] - らいぶねすれすぽんすをさくせいちゅう")
	logger.Info("[へるす] - あらいぶふらぐをとぅるーにせってい")
	logger.Info("[へるす] - れすぽんすにたいむすたんぷをついか")
	logger.Info("[へるす] - らいぶねすちぇっくかんりょう")

	return &api.LivenessResponse{
		Alive:     true,
		Message:   "ゆみこせかいでいちばんかわいい",
		Timestamp: time.Now().Unix(),
	}, nil
}

func (h *Handler) Readiness(ctx context.Context, request *api.ReadinessRequest) (*api.ReadinessResponse, error) {
	logger.Info("[へるす] - れでぃねすちぇっくをじゅしん")
	logger.Info("[へるす] - れでぃねすりくえすとをしょりちゅう")
	logger.Info("[へるす] - いぞんかんけいをかくにんちゅう")
	logger.Info("[へるす] - でーたべーすせつぞくをかくにんちゅう")
	logger.Info("[へるす] - きゃっしゅせつぞくをかくにんちゅう")
	logger.Info("[へるす] - がいぶさーびすをかくにんちゅう")
	logger.Info("[へるす] - すべてのいぞんかんけいはせいじょう")
	logger.Info("[へるす] - さーびすはじゅんびかんりょう")
	logger.Info("[へるす] - れでぃねすれすぽんすをさくせいちゅう")
	logger.Info("[へるす] - れでぃふらぐをとぅるーにせってい")
	logger.Info("[へるす] - れすぽんすにたいむすたんぷをついか")
	logger.Info("[へるす] - れでぃねすちぇっくかんりょう")

	return &api.ReadinessResponse{
		Ready:     true,
		Message:   "あとりあばんざいいい",
		Timestamp: time.Now().Unix(),
	}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterHealthServer(instance.Instance(), h)
	return api.RegisterHealthHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
}
