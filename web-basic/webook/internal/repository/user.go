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
	d := dao.NewUserDao()
	d.Insert(ctx, dao.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	})
	return nil
}

//
//type UserSer
