package dao

import (
	"github.com/Daarkness/logprint/web-basic/webook/internal/domain"

	"context"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}
func (ud *UserDao) Insert(ctx context.Context, u domain.User) error {
	return nil

}

type User struct {
	Id    int64
	Name  string
	Email string
	Phone string
}
