package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) FindAuthenticationByUsername(username string) (*model.Authentication, error) {
	auth := &model.Authentication{
		Username: username,
	}
	return auth.Find()
}

func (dao *Dao) FindAuthenticationByToken(token string) (*model.Authentication, error) {
	auth := &model.Authentication{
		Token: token,
	}
	return auth.Find()
}

func (dao *Dao) CreateAuthentication(token string, username string, password string) (*model.Authentication, error) {
	auth := &model.Authentication{
		Username: username,
		Password: password,
		Token:    token,
	}
	return auth.Create()
}

func (dao *Dao) CountAuthentication(username string) (int64, error) {
	auth := &model.Authentication{
		Username: username,
	}
	return auth.Count()
}

func (dao *Dao) UpdateAuthenticationByToken(token string, password string) error {
	auth := &model.Authentication{
		Password: password,
		Token:    token,
	}
	return auth.Update()
}

func (dao *Dao) DeleteAuthenticationByToken(token string) error {
	auth := &model.Authentication{
		Token: token,
	}
	return auth.Delete()
}
