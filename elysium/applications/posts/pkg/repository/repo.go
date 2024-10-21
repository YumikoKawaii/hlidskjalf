package repository

import (
	"context"
	"elysium.com/applications/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	UpsertPost(ctx context.Context, post *Post) error
	GetPosts(ctx context.Context, params *GetPostsParams) ([]Post, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repoImpl{
		db: db,
	}
}

func (r *repoImpl) UpsertPost(ctx context.Context, post *Post) error {
	return r.db.Model(&Post{}).Clauses(clause.OnConflict{UpdateAll: true}).Create(post).Error
}

func (r *repoImpl) GetPosts(ctx context.Context, params *GetPostsParams) ([]Post, error) {
	query := r.db.Model(&Post{})
	if len(params.Ids) != 0 {
		query = query.Where("id in (?)", params.Ids)
	}

	if len(params.Author) != 0 {
		query = query.Where("author = ?", params.Author)
	}

	if params.Page != 0 && params.PageSize != 0 {
		offset := (params.Page - 1) * params.PageSize
		query = query.Limit(int(params.PageSize)).Offset(int(offset))
	}

	posts := make([]Post, 0)

	err := query.Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: "created_at",
		},
		Desc: params.Order == utils.DESC,
	}).Scan(&posts).Error
	return posts, err
}
