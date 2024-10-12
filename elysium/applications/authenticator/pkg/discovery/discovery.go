package discovery

import (
	"context"
	"elysium.com/applications/authenticator/pkg/repository"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

type Storage interface {
	Operate(ctx context.Context, interval time.Duration)
}

type storageImpl struct {
	redisClient *redis.Client
	repo        repository.Repository
}

func NewStorage(client *redis.Client, repo repository.Repository) Storage {
	return &storageImpl{
		redisClient: client,
		repo:        repo,
	}
}

func (s *storageImpl) Operate(ctx context.Context, interval time.Duration) {

	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			s.updateCache(ctx)
		}
	}

}

func (s *storageImpl) updateCache(ctx context.Context) {
	// fetch all permission
	permissions, err := s.repo.GetPermissions(ctx)
	if err != nil {
		logrus.Errorf("error fetch permissions: %s", err.Error())
	}

	// group by maps
	permissionMap := make(map[string][]string)
	for _, permission := range permissions {
		permissionMap[permission.AccountId] = append(permissionMap[permission.AccountId], permission.Route)
	}

	// update redis
	for accountId, pers := range permissionMap {
		if err := s.redisClient.SAdd(ctx, accountId, pers); err != nil {
			logrus.Errorf("error setting value to redis: %s", err)
		}
	}
}
