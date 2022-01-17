package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateUserDepartmentAssociations(userID uint, departmentIDs []uint) ([]*model.UserDepartmentAssociation, error) {
	association := &model.UserDepartmentAssociation{}
	var associations []*model.UserDepartmentAssociation
	for _, departmentID := range departmentIDs {
		associations = append(associations,
			&model.UserDepartmentAssociation{UserID: userID, DepartmentID: departmentID})
	}
	return association.CreateMany(associations)
}

func (dao *Dao) GetUserDepartmentAssociations(userID uint) ([]*model.UserDepartmentAssociation, error) {
	association := &model.UserDepartmentAssociation{UserID: userID}
	return association.Get()
}

func (dao *Dao) DeleteUserDepartmentAssociationsByUserID(userID uint) error {
	association := &model.UserDepartmentAssociation{UserID: userID}
	return association.Delete()
}

func (dao *Dao) DeleteUserDepartmentAssociationsByDepartmentID(departmentID uint) error {
	association := &model.UserDepartmentAssociation{DepartmentID: departmentID}
	return association.Delete()
}
