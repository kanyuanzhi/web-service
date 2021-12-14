package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) FindAuthentication(username string) *model.Authentication {
	auth := &model.Authentication{
		Username: username,
	}
	return auth.Find()
}

func (dao *Dao) CreateAuthentication(token string, username string, password string) uint {
	auth := &model.Authentication{
		Username: username,
		Password: password,
		Token:    token,
	}
	return auth.Create()
}

func (dao *Dao) CountAuthentication(username string) int64 {
	auth := &model.Authentication{
		Username: username,
	}
	return auth.Count()
}
