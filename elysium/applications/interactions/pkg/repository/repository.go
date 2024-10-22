package repository

import (
	"context"
	"elysium.com/applications/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	UpsertInteraction(ctx context.Context, interaction *Interaction) error
	GetInteractions(ctx context.Context, params *GetInteractionsParams) ([]Interaction, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repoImpl{
		db: db,
	}
}

func (r *repoImpl) UpsertInteraction(ctx context.Context, interaction *Interaction) error {
	return r.db.Model(&Interaction{}).Clauses(clause.OnConflict{UpdateAll: true}).Create(interaction).Error
}

func (r *repoImpl) GetInteractions(ctx context.Context, params *GetInteractionsParams) ([]Interaction, error) {
	query := r.db.Model(&Interaction{})
	query = query.Where("post_id = ?", params.PostId)

	if params.Page != 0 && params.PageSize != 0 {
		offset := (params.Page - 1) * params.PageSize
		query = query.Limit(int(params.PageSize)).Offset(int(offset))
	}

	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: "created_at",
		},
		Desc: params.Order == utils.DESC,
	})

	interactions := make([]Interaction, 0)
	err := query.Scan(&interactions).Error
	return interactions, err
}
