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
	GetPermissions(ctx context.Context) ([]Permission, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repoImpl{db: db}
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
	err := r.db.Model(&Account{}).Where("hashed_email = ?", hashedEmail).First(account).Error
	return account, err
}

func (r *repoImpl) UpsertPermissions(ctx context.Context, permissions []Permission, batchSize int) error {
	return r.db.Model(&Permission{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(permissions, batchSize).Error
}

func (r *repoImpl) GetPermissions(ctx context.Context) ([]Permission, error) {
	permissions := make([]Permission, 0)
	err := r.db.Model(&Permission{}).Scan(&permissions).Error
	return permissions, err
}
