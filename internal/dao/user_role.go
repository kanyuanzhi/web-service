package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateUserRole(userID uint, roleName string) uint {
	userRole := &model.UserRole{
		UserID:   userID,
		RoleName: roleName,
	}
	return userRole.Create()
}

func (dao *Dao) GetUserRoles(userID uint) []*model.UserRole {
	userRole := &model.UserRole{}
	userRole.UserID = userID
	return userRole.Get()
}
