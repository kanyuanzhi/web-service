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

	RoleNames     []string `json:"roles" form:"roles" gorm:"-"`
	DepartmentIDs []uint   `json:"departments" form:"departments" gorm:"-"`
}

func (u *User) Create() (*User, error) {
	err := global.DB.Create(u).Error
	return u, err
}

func (u *User) Get() (*User, error) {
	var user User
	err := global.DB.Where(u).Find(&user).Error
	return &user, err
}

func (u *User) List() ([]*User, error) {
	var users []*User
	err := global.DB.Find(&users).Error
	return users, err
}

func (u *User) Update() (*User, error) {
	err := global.DB.Model(u).Updates(u).Error
	return u, err
}

func (u *User) Delete() error {
	err := global.DB.Where(u).Delete(u).Error
	return err
}
