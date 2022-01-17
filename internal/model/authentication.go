package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type Authentication struct {
	*DefaultFields
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (auth *Authentication) Find() (*Authentication, error) {
	var authentication Authentication
	err := global.DB.Where(auth).Find(&authentication).Error
	return &authentication, err
}

func (auth *Authentication) Create() (*Authentication, error) {
	err := global.DB.Create(auth).Error
	return auth, err
}

func (auth *Authentication) Count() (int64, error) {
	var count int64
	err := global.DB.Model(auth).Where(auth).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (auth *Authentication) Update() error {
	return global.DB.Model(auth).Where("token = ?", auth.Token).
		Update("password", auth.Password).Error
}

func (auth *Authentication) Delete() error {
	return global.DB.Where(auth).Delete(auth).Error
}
