package auth

import (
	"context"
	"elysium.com/applications/authenticator/pkg/repository"
	"elysium.com/applications/utils"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	updatePermissionBatchSize = 100
)

type Service interface {
	Signup(ctx context.Context, email, password string) (*repository.Account, error)
	Login(ctx context.Context, email, password string) (*repository.Account, error)

	UpsertPermissions(ctx context.Context, permissions []repository.Permission) error
}

type Secret struct {
	HashKey    string
	EncryptKey string
}

type serviceImpl struct {
	repo   repository.Repository
	secret Secret
}

func NewService() Service {
	return &serviceImpl{}
}

func (s *serviceImpl) Signup(ctx context.Context, email, password string) (*repository.Account, error) {

	hashedEmail := utils.HashString(email, s.secret.HashKey)
	_, err := s.repo.GetAccountByHashedEmail(ctx, hashedEmail)
	if err == nil {
		return nil, ErrExistAccount
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInternal
	}

	encryptEmail, err := utils.Encrypt([]byte(email), s.secret.EncryptKey)
	if err != nil {
		return nil, ErrInternal
	}

	hashedPassword := utils.HashString(password, s.secret.HashKey)

	account := &repository.Account{
		Id:             uuid.New().String(),
		HashedEmail:    hashedEmail,
		EncryptEmail:   string(encryptEmail),
		HashedPassword: hashedPassword,
	}

	err = s.repo.UpsertAccount(ctx, account)
	return account, err
}

func (s *serviceImpl) Login(ctx context.Context, email, password string) (*repository.Account, error) {

	hashedEmail := utils.HashString(email, s.secret.HashKey)
	account, err := s.repo.GetAccountByHashedEmail(ctx, hashedEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotExistAccount
		}

		return nil, ErrInternal
	}

	hashedPassword := utils.HashString(password, s.secret.HashKey)
	if account.HashedPassword != hashedPassword {
		return nil, ErrWrongAccountInfo
	}

	return account, nil
}

func (s *serviceImpl) UpsertPermissions(ctx context.Context, permissions []repository.Permission) error {
	return s.repo.UpsertPermissions(ctx, permissions, updatePermissionBatchSize)
}
