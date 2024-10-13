package interaction_service

import (
	"context"
	"elysium.com/applications/interactions/pkg/repository"
)

const (
	defaultPage = 1
	defaultPageSize
)

type Service interface {
	UpsertInteraction(ctx context.Context, interaction *repository.Interaction) error
	GetInteractions(ctx context.Context, params *repository.GetInteractionsParams) ([]repository.Interaction, error)
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (s *serviceImpl) UpsertInteraction(ctx context.Context, interaction *repository.Interaction) error {
	return s.repo.UpsertInteraction(ctx, interaction)
}

func (s *serviceImpl) GetInteractions(ctx context.Context, params *repository.GetInteractionsParams) ([]repository.Interaction, error) {
	if params.Page == 0 {
		params.Page = defaultPage
	}
	if params.PageSize == 0 {
		params.PageSize = defaultPageSize
	}

	return s.repo.GetInteractions(ctx, params)
}
