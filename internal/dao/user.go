package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateUser(token string, username string) uint {
	user := &model.User{
		Token:    token,
		Username: username,
	}
	return user.Create()
}

func (dao *Dao) GetUser(token string) *model.User {
	user := &model.User{
		Token: token,
	}
	return user.Get()
}
