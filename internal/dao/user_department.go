package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateUserDepartments(userID uint, departmentIDs []uint)[]*model.UserDepartment  {
	userDepartment := &model.UserDepartment{}
	var userDepartments []*model.UserDepartment
	for _,departmentID := range departmentIDs{
		userDepartments = append(userDepartments,
			&model.UserDepartment{UserID: userID,DepartmentID: departmentID})
	}
	return userDepartment.CreateMany(userDepartments)
}

func (dao *Dao) GetUserDepartments(userID uint) []*model.UserDepartment {
	userDepartment := &model.UserDepartment{}
	userDepartment.UserID = userID
	return userDepartment.Get()
}

func (dao *Dao) DeleteUserDepartments(userID uint) {
	userDepartment := &model.UserDepartment{UserID: userID}
	userDepartment.DeleteByUserID()
}
