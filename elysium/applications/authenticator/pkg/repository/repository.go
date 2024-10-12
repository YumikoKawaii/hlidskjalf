package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	UpsertAccount(ctx context.Context, account *Account) error
	GetAccountById(ctx context.Context, id string) (*Account, error)
	GetAccountByHashedEmail(ctx context.Context, email string) (*Account, error)

	UpsertPermissions(ctx context.Context, permissions []Permission, batchSize int) error
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &repoImpl{}
}

func (r *repoImpl) UpsertAccount(ctx context.Context, account *Account) error {
	return r.db.Model(&Account{}).Clauses(clause.OnConflict{UpdateAll: true}).Create(account).Error
}

func (r *repoImpl) GetAccountById(ctx context.Context, id string) (*Account, error) {
	account := new(Account)
	err := r.db.Model(&Account{}).Where("id = ?", id).Find(account).Error
	return account, err
}

func (r *repoImpl) GetAccountByHashedEmail(ctx context.Context, hashedEmail string) (*Account, error) {
	account := new(Account)
	err := r.db.Model(&Account{}).Where("hashed_email = ?", hashedEmail).Find(account).Error
	return account, err
}

func (r *repoImpl) UpsertPermissions(ctx context.Context, permissions []Permission, batchSize int) error {
	return r.db.Model(&Permission{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(permissions, batchSize).Error
}
