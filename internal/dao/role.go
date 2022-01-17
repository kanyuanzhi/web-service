package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) ListRoles() ([]*model.Role, error) {
	role := &model.Role{}
	return role.List()
}

func (dao *Dao) UpdateRole(id uint, introduction string) (*model.Role, error) {
	role := &model.Role{
		DefaultFields: &model.DefaultFields{ID: id},
		Introduction:  introduction,
	}
	return role.Update()
}
