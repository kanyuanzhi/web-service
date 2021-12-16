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

func (dao *Dao) ListUsers() []*model.User {
	user := &model.User{}
	return user.List()
}

func (dao *Dao) UpdateUser(id uint, name string, contact string, introduction string, avatar string) *model.User {
	user := &model.User{
		DefaultFields: &model.DefaultFields{ID: id},
		Name:          name,
		Contact:       contact,
		Introduction:  introduction,
		Avatar:        avatar,
	}
	return user.Update()
}

func (dao *Dao) DeleteUser(token string) error {
	user := &model.User{
		Token: token,
	}
	return user.Delete()
}
