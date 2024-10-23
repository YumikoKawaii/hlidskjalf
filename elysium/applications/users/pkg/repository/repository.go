package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	UpsertUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context, params *GetUsersParams) ([]User, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repoImpl{
		db: db,
	}
}

func (r *repoImpl) UpsertUser(ctx context.Context, user *User) error {
	return r.db.Model(&User{}).Clauses(clause.OnConflict{UpdateAll: true}).Create(user).Error
}

func (r *repoImpl) GetUsers(ctx context.Context, params *GetUsersParams) ([]User, error) {
	users := make([]User, 0)
	query := r.db.Model(&User{})
	if len(params.Ids) != 0 {
		query = query.Where("id in (?)", params.Ids)
	}

	if params.Page != 0 && params.PageSize != 0 {
		offset := (params.Page - 1) * params.PageSize
		query = query.Offset(int(offset)).Limit(int(params.PageSize))
	}
	err := query.Scan(&users).Error
	return users, err
}
