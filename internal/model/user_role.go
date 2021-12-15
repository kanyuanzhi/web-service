package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type UserRole struct {
	*DefaultFields
	UserID   uint   `json:"user_id"`
	RoleName string `json:"role_name"`
}

func (ur *UserRole) TableName() string {
	return "users_roles"
}

func (ur *UserRole) CreateMany(userRoles []*UserRole) []*UserRole {
	err := global.DB.Create(&userRoles).Error
	if err != nil {
		global.Log.Error(err)
	}
	return userRoles
}

func (ur *UserRole) Get() []*UserRole {
	var userRoles []*UserRole
	err := global.DB.Where(ur).Find(&userRoles).Error
	if err != nil {
		global.Log.Error(err)
	}
	return userRoles
}

func (ur *UserRole) DeleteByUserID() {
	err := global.DB.Where(ur).Delete(ur).Error
	if err != nil {
		global.Log.Error(err)
	}
}
