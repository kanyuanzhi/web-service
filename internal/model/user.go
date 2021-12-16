package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type User struct {
	*DefaultFields
	Token        string `json:"token"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Contact      string `json:"contact"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`

	Roles       []string `json:"roles" gorm:"-"`
	Departments []uint   `json:"departments" gorm:"-"`
}

func (u *User) Create() uint {
	err := global.DB.Create(u).Error
	if err != nil {
		global.Log.Error(err)
	}
	return u.ID
}

func (u *User) Get() *User {
	var user User
	err := global.DB.Where(u).Find(&user).Error
	if err != nil {
		global.Log.Error(err)
	}
	return &user
}

func (u *User) List() []*User {
	var users []*User
	err := global.DB.Find(&users).Error
	if err != nil {
		global.Log.Error(err)
	}
	return users
}

func (u *User) Update() *User {
	err := global.DB.Model(u).Updates(u).Error
	if err != nil {
		global.Log.Error(err)
	}
	return u
}

func (u *User) Delete() error {
	err := global.DB.Where("token = ?", u.Token).Delete(u).Error
	if err != nil {
		return err
	}
	return nil
}
