package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type UserDepartmentAssociation struct {
	*DefaultFields
	UserID       uint `json:"user_id"`
	DepartmentID uint `json:"department_id"`
}

func (uda *UserDepartmentAssociation) TableName() string {
	return "user_department_associations"
}

func (uda *UserDepartmentAssociation) CreateMany(userDepartmentAssociations []*UserDepartmentAssociation) ([]*UserDepartmentAssociation, error) {
	err := global.DB.Create(&userDepartmentAssociations).Error
	return userDepartmentAssociations, err
}

func (uda *UserDepartmentAssociation) Get() ([]*UserDepartmentAssociation, error) {
	var UserDepartments []*UserDepartmentAssociation
	err := global.DB.Where(uda).Find(&UserDepartments).Error
	return UserDepartments, err
}

func (uda *UserDepartmentAssociation) Delete() error {
	err := global.DB.Where(uda).Delete(uda).Error
	return err
}
