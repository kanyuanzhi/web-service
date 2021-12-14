package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type Role struct {
	*DefaultFields
	Name         string
	Introduction string
}

func (r *Role) Get() *Role {
	var role Role
	err := global.DB.Where(r).Find(&role).Error
	if err != nil {
		global.Log.Error(err)
	}
	return &role
}
