package serve

import (
	"context"

	"github.com/YumikoKawaii/hlidskjalf/applications/acoustics/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/acoustics/handlers/acoustics"
	"github.com/YumikoKawaii/hlidskjalf/applications/acoustics/handlers/health"
	"github.com/YumikoKawaii/hlidskjalf/applications/acoustics/interceptors"
	"github.com/YumikoKawaii/hlidskjalf/applications/acoustics/workers"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func Server(_ *cobra.Command, _ []string) {
	logger.Info("[あこーすてぃくす] - しょきかちゅう...")
	logger.Info("[あこーすてぃくす] - せっていをよみこみちゅう")
	logger.Info("[あこーすてぃくす] - せっていふぁいるをよんでいます")
	logger.Info("[あこーすてぃくす] - せっていちをかいせきちゅう")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.Info("[あこーすてぃくす] - せっていのよみこみせいこう")
	logger.Info("[あこーすてぃくす] - せっていをけんしょうちゅう")
	logger.Info("[あこーすてぃくす] - せっていのけんしょうかんりょう")
	logger.Info("[あこーすてぃくす] - とれーさーをせっていちゅう")
	logger.Info("[あこーすてぃくす] - おーぷんてれめとりーとれーさーしょきかかんりょう")
	logger.Info("[あこーすてぃくす] - じーあーるぴーしーさーばーをじゅんびちゅう")
	logger.Info("[あこーすてぃくす] - きーぷあらいぶぱらめーたーをせっていちゅう")
	logger.Info("[あこーすてぃくす] - ゆなりーいんたーせぷたーをせっていちゅう")
	logger.Info("[あこーすてぃくす] - すとりーむいんたーせぷたーをせっていちゅう")

	instance := server.Initialize(
		cfg.Server,
		grpc.KeepaliveParams(keepalive.ServerParameters{}),
		grpc.ChainUnaryInterceptor(
			interceptors.UnaryLoggingInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
			grpcvalidator.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			interceptors.StreamLoggingInterceptor(),
			grpcprometheus.StreamServerInterceptor,
			grpcvalidator.StreamServerInterceptor(),
			grpcrecovery.StreamServerInterceptor(),
		),
	)

	logger.Info("[あこーすてぃくす] - じーあーるぴーしーさーばーいんすたんすさくせい")
	logger.Info("[あこーすてぃくす] - ぷろめてうすいんたーせぷたーをとうろくちゅう")
	logger.Info("[あこーすてぃくす] - ばりでーたーいんたーせぷたーをとうろくちゅう")
	logger.Info("[あこーすてぃくす] - りかばりーいんたーせぷたーをとうろくちゅう")
	logger.Info("[あこーすてぃくす] - ろぎんぐいんたーせぷたーをとうろくちゅう")
	logger.Info("[あこーすてぃくす] - すべてのいんたーせぷたーとうろくかんりょう")
	logger.Info("[あこーすてぃくす] - へるすはんどらーをしょきかちゅう")

	healthHandler := health.Initialize()

	logger.Info("[あこーすてぃくす] - へるすはんどらーさくせいかんりょう")
	logger.Info("[あこーすてぃくす] - あこーすてぃくすはんどらーをしょきかちゅう")

	acousticsHandler := acoustics.Initialize()

	logger.Info("[あこーすてぃくす] - あこーすてぃくすはんどらーさくせいかんりょう")
	logger.Info("[あこーすてぃくす] - へるすはんどらーをとうろくちゅう")

	if err := healthHandler.Register(instance); err != nil {
		panic(err)
	}

	logger.Info("[あこーすてぃくす] - へるすはんどらーとうろくかんりょう")
	logger.Info("[あこーすてぃくす] - あこーすてぃくすはんどらーをとうろくちゅう")

	if err := acousticsHandler.Register(instance); err != nil {
		panic(err)
	}

	logger.Info("[あこーすてぃくす] - あこーすてぃくすはんどらーとうろくかんりょう")
	logger.Info("[あこーすてぃくす] - らいぶねすえんどぽいんとじゅんびかんりょう")
	logger.Info("[あこーすてぃくす] - れでぃねすえんどぽいんとじゅんびかんりょう")
	logger.Info("[あこーすてぃくす] - ちゃーじえんどぽいんとじゅんびかんりょう")
	logger.Info("[あこーすてぃくす] - でぃすちゃーじえんどぽいんとじゅんびかんりょう")
	logger.Info("[あこーすてぃくす] - すべてのはんどらーとうろくかんりょう")
	logger.Info("[あこーすてぃくす] - さーばーしょきかかんりょう")
	logger.Info("[あこーすてぃくす] - じーあーるぴーしーさーばーをきどうちゅう")
	logger.Info("[あこーすてぃくす] - えいちてぃーてぃーぴーげーとうぇいをきどうちゅう")
	logger.Info("[あこーすてぃくす] - えらーえみったーをきどうちゅう")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errorEmitter := &workers.ErrorEmitter{Interval: cfg.ErrorEmitter.Interval}
	errorEmitter.Start(ctx)

	logger.Info("[あこーすてぃくす] - えらーえみったーきどうかんりょう")
	logger.Info("[あこーすてぃくす] - せつぞくをまっています...")

	if err := instance.Serve(); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatalf("さーばーきどうえらー")
	}
}
