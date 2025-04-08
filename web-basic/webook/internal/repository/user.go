package repository

import (
	"context"
	"github.com/Daarkness/logprint/web-basic/webook/internal/domain"
	"github.com/Daarkness/logprint/web-basic/webook/internal/repository/dao"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) Create(ctx context.Context, u domain.User) error {
	dao := user.NewUserDao()
	dao.INSERT(ctx, u)
	return nil
}

//
//type UserSer
