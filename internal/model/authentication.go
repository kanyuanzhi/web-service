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
	if err != nil {
		return nil, err
	}
	return &authentication, nil
}

func (auth *Authentication) Create() {
	err := global.DB.Create(auth).Error
	if err != nil {
		global.Log.Error(err)
	}
}

func (auth *Authentication) Count() int64 {
	var count int64
	err := global.DB.Model(auth).Where(auth).Count(&count).Error
	if err != nil {
		global.Log.Error(err)
		return -1
	}
	return count
}

func (auth *Authentication) Update() error {
	return global.DB.Model(auth).Where("token = ?", auth.Token).
		Update("password", auth.Password).Error
}

func (auth *Authentication) Delete() error {
	return global.DB.Where("token = ?", auth.Token).
		Delete(auth).Error
}
