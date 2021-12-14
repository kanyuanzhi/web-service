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

func (auth *Authentication) Find() *Authentication {
	var authentication Authentication
	err := global.DB.Where(auth).Find(&authentication).Error
	if err != nil {
		global.Log.Error(err)
		return nil
	}
	return &authentication
}

func (auth *Authentication) Create() uint {
	err := global.DB.Create(auth).Error
	if err != nil {
		global.Log.Error(err)
	}
	return auth.ID
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
