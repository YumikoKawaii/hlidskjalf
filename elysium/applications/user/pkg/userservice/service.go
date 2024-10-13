package userservice

import (
	"context"
	"elysium.com/applications/user/pkg/repository"
)

type Service interface {
	UpsertUser(ctx context.Context, user *repository.User) error
	GetUsersInfoByIds(ctx context.Context, ids []string) ([]repository.User, error)
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (s *serviceImpl) UpsertUser(ctx context.Context, user *repository.User) error {
	return s.repo.UpsertUser(ctx, user)
}

func (s *serviceImpl) GetUsersInfoByIds(ctx context.Context, ids []string) ([]repository.User, error) {
	return s.repo.GetUsersByIds(ctx, ids)
}
