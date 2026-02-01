package serve

import (
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/adapters/acoustics"
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/handlers/echo"
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/handlers/health"
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/interceptors"
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
	logger.Info("[えこー] - しょきかちゅう...")
	logger.Info("[えこー] - せっていをよみこみちゅう")
	logger.Info("[えこー] - せっていふぁいるをよんでいます")
	logger.Info("[えこー] - せっていちをかいせきちゅう")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.Info("[えこー] - せっていのよみこみせいこう")
	logger.Info("[えこー] - せっていをけんしょうちゅう")
	logger.Info("[えこー] - せっていのけんしょうかんりょう")
	logger.Info("[えこー] - とれーさーをせっていちゅう")
	logger.Info("[えこー] - おーぷんてれめとりーとれーさーしょきかかんりょう")
	logger.Info("[えこー] - じーあーるぴーしーさーばーをじゅんびちゅう")
	logger.Info("[えこー] - きーぷあらいぶぱらめーたーをせっていちゅう")
	logger.Info("[えこー] - ゆなりーいんたーせぷたーをせっていちゅう")
	logger.Info("[えこー] - すとりーむいんたーせぷたーをせっていちゅう")

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

	logger.Info("[えこー] - じーあーるぴーしーさーばーいんすたんすさくせい")
	logger.Info("[えこー] - ぷろめてうすいんたーせぷたーをとうろくちゅう")
	logger.Info("[えこー] - ばりでーたーいんたーせぷたーをとうろくちゅう")
	logger.Info("[えこー] - りかばりーいんたーせぷたーをとうろくちゅう")
	logger.Info("[えこー] - ろぎんぐいんたーせぷたーをとうろくちゅう")
	logger.Info("[えこー] - すべてのいんたーせぷたーとうろくかんりょう")
	logger.Info("[えこー] - へるすはんどらーをしょきかちゅう")

	healthHandler := health.Initialize()

	logger.Info("[えこー] - へるすはんどらーさくせいかんりょう")
	logger.Info("[えこー] - あこーすてぃくすくらいあんとをしょきかちゅう")

	var acousticsClient acoustics.Client
	switch cfg.Acoustics.Protocol {
	case "http":
		acousticsClient, err = acoustics.NewHTTPClient(cfg.Acoustics.Endpoint)
	default:
		acousticsClient, err = acoustics.NewGRPCClient(cfg.Acoustics.Endpoint)
	}
	if err != nil {
		panic(err)
	}

	logger.Info("[えこー] - あこーすてぃくすくらいあんとさくせいかんりょう")
	logger.Info("[えこー] - えこーはんどらーをしょきかちゅう")

	echoHandler := echo.Initialize(acousticsClient)

	logger.Info("[えこー] - えこーはんどらーさくせいかんりょう")
	logger.Info("[えこー] - へるすはんどらーをとうろくちゅう")

	if err := healthHandler.Register(instance); err != nil {
		panic(err)
	}

	logger.Info("[えこー] - へるすはんどらーとうろくかんりょう")
	logger.Info("[えこー] - えこーはんどらーをとうろくちゅう")

	if err := echoHandler.Register(instance); err != nil {
		panic(err)
	}

	logger.Info("[えこー] - えこーはんどらーとうろくかんりょう")
	logger.Info("[えこー] - らいぶねすえんどぽいんとじゅんびかんりょう")
	logger.Info("[えこー] - れでぃねすえんどぽいんとじゅんびかんりょう")
	logger.Info("[えこー] - ちゃーじえんどぽいんとじゅんびかんりょう")
	logger.Info("[えこー] - でぃすちゃーじえんどぽいんとじゅんびかんりょう")
	logger.Info("[えこー] - すべてのはんどらーとうろくかんりょう")
	logger.Info("[えこー] - さーばーしょきかかんりょう")
	logger.Info("[えこー] - じーあーるぴーしーさーばーをきどうちゅう")
	logger.Info("[えこー] - えいちてぃーてぃーぴーげーとうぇいをきどうちゅう")
	logger.Info("[えこー] - せつぞくをまっています...")

	if err := instance.Serve(); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatalf("さーばーきどうえらー")
	}
}
