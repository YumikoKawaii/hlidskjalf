package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	UpsertUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUsersByIds(ctx context.Context, ids []string) ([]User, error)
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

func (r *repoImpl) GetUserById(ctx context.Context, id string) (*User, error) {
	user := new(User)
	err := r.db.Model(&User{}).Where("id = ?", id).First(user).Error
	return user, err
}

func (r *repoImpl) GetUsersByIds(ctx context.Context, ids []string) ([]User, error) {
	users := make([]User, 0)
	err := r.db.Model(&User{}).Where("id in (?)", ids).Scan(users).Error
	return users, err
}
