package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type Role struct {
	*DefaultFields
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

func (r *Role) List() []*Role {
	var roles []*Role
	err := global.DB.Find(&roles).Error
	if err != nil {
		global.Log.Error(err)
	}
	return roles
}
