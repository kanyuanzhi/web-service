package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateUserRoles(userID uint, roleNames []string) []*model.UserRole {
	userRole := &model.UserRole{}
	var userRoles []*model.UserRole
	for _, roleName := range roleNames {
		userRoles = append(userRoles,
			&model.UserRole{UserID: userID, RoleName: roleName})
	}
	return userRole.CreateMany(userRoles)
}

func (dao *Dao) GetUserRoles(userID uint) []*model.UserRole {
	userRole := &model.UserRole{UserID: userID}
	return userRole.Get()
}

func (dao *Dao) DeleteUserRoles(userID uint) {
	userRole := &model.UserRole{UserID: userID}
	userRole.DeleteByUserID()
}
