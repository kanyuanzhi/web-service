package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type UserRoleAssociation struct {
	*DefaultFields
	UserID   uint   `json:"user_id"`
	RoleName string `json:"role_name"`
}

func (ura *UserRoleAssociation) TableName() string {
	return "user_role_associations"
}

func (ura *UserRoleAssociation) CreateMany(associations []*UserRoleAssociation) ([]*UserRoleAssociation, error) {
	err := global.DB.Create(&associations).Error
	return associations, err
}

func (ura *UserRoleAssociation) Get() ([]*UserRoleAssociation, error) {
	var userRoles []*UserRoleAssociation
	err := global.DB.Where(ura).Find(&userRoles).Error
	return userRoles, err
}

func (ura *UserRoleAssociation) Delete() error {
	err := global.DB.Where(ura).Delete(ura).Error
	return err
}
