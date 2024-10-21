package post_service

import (
	"context"
	"elysium.com/applications/posts/pkg/repository"
)

const (
	defaultGetPostsPage     = 1
	defaultGetPostsPageSize = 50
)

type Service interface {
	UpsertPost(ctx context.Context, post *repository.Post) error
	GetPosts(ctx context.Context, params *repository.GetPostsParams) ([]repository.Post, error)
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) UpsertPost(ctx context.Context, post *repository.Post) error {
	return s.repo.UpsertPost(ctx, post)
}

func (s *serviceImpl) GetPosts(ctx context.Context, params *repository.GetPostsParams) ([]repository.Post, error) {
	return s.repo.GetPosts(ctx, params)
}
