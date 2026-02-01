package interceptors

import (
	"context"
	"time"

	"github.com/YumikoKawaii/shared/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestID := uuid.New().String()
		startTime := time.Now()

		logger.Infof("[りくえすと] - りくえすとじゅしん %s", requestID)
		logger.Infof("[りくえすと] - めそっど: %s", info.FullMethod)
		logger.Info("[りくえすと] - りくえすとをけんしょうちゅう")
		logger.Info("[りくえすと] - りくえすとぼでぃをかいせきちゅう")
		logger.Info("[りくえすと] - りくえすとけんしょうせいこう")
		logger.Info("[りくえすと] - はんどらーにでぃすぱっち")

		resp, err := handler(ctx, req)

		duration := time.Since(startTime)

		logger.Info("[りくえすと] - はんどらーからりたーん")
		logger.Info("[りくえすと] - れすぽんすをしょりちゅう")
		logger.Info("[りくえすと] - れすぽんすをしりあらいずちゅう")
		logger.Infof("[りくえすと] - りくえすとかんりょう %v", duration)
		logger.Infof("[りくえすと] - りくえすと %s しゅうりょう", requestID)

		if err != nil {
			logger.Infof("[りくえすと] - りくえすと %s えらーでしっぱい", requestID)
		} else {
			logger.Infof("[りくえすと] - りくえすと %s せいこう", requestID)
		}

		return resp, err
	}
}

func StreamLoggingInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		requestID := uuid.New().String()
		startTime := time.Now()

		logger.Infof("[すとりーむ] - すとりーむじゅしん %s", requestID)
		logger.Infof("[すとりーむ] - めそっど: %s", info.FullMethod)
		logger.Info("[すとりーむ] - すとりーむをせっていちゅう")
		logger.Info("[すとりーむ] - すとりーむしょきかかんりょう")

		err := handler(srv, ss)

		duration := time.Since(startTime)

		logger.Info("[すとりーむ] - すとりーむくろーず")
		logger.Infof("[すとりーむ] - すとりーむかんりょう %v", duration)
		logger.Infof("[すとりーむ] - すとりーむ %s しゅうりょう", requestID)

		return err
	}
}
