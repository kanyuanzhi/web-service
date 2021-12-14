package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) GetRole(id uint) *model.Role {
	user := &model.Role{}
	user.ID = id
	return user.Get()
}
