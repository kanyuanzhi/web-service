package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) ListRoles() []*model.Role {
	role := &model.Role{}
	return role.List()
}