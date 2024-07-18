package students

import "context"

type Service interface {
	GetStudents(ctx context.Context) ([]Student, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}

func (s *serviceImpl) GetStudents(ctx context.Context) ([]Student, error) {
	return s.repo.GetStudents(ctx)
}
