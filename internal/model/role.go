package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type Role struct {
	*DefaultFields
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

func (r *Role) List() ([]*Role, error) {
	var roles []*Role
	err := global.DB.Find(&roles).Error
	return roles, err
}

func (r *Role) Update() (*Role, error) {
	err := global.DB.Model(r).Updates(r).Error
	return r, err
}
