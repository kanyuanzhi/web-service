package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type UserDepartment struct {
	*DefaultFields
	UserID       uint `json:"user_id"`
	DepartmentID uint `json:"department_id"`
}

func (ud *UserDepartment) TableName() string {
	return "users_departments"
}

func (ud *UserDepartment) CreateMany(userDepartments []*UserDepartment) []*UserDepartment {
	err := global.DB.Create(&userDepartments).Error
	if err != nil {
		global.Log.Error(err)
	}
	return userDepartments
}

func (ud *UserDepartment) Get() []*UserDepartment {
	var UserDepartments []*UserDepartment
	err := global.DB.Where(ud).Find(&UserDepartments).Error
	if err != nil {
		global.Log.Error(err)
	}
	return UserDepartments
}

func (ud *UserDepartment) Delete() {
	err := global.DB.Where(ud).Delete(ud).Error
	if err != nil {
		global.Log.Error(err)
	}
}

