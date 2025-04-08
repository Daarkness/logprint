package service

import (
	""
	"context"
	"github.com/Daarkness/logprint/web-basic/webook/internal/domain"
	"github.com/Daarkness/logprint/web-basic/webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	return svc.repo.Create(ctx, u)
}
