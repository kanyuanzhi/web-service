package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateUserRoleAssociations(userID uint, roleNames []string) ([]*model.UserRoleAssociation, error) {
	association := &model.UserRoleAssociation{}
	var associations []*model.UserRoleAssociation
	for _, roleName := range roleNames {
		associations = append(associations,
			&model.UserRoleAssociation{UserID: userID, RoleName: roleName})
	}
	return association.CreateMany(associations)
}

func (dao *Dao) GetUserRoleAssociationsByUserID(userID uint) ([]*model.UserRoleAssociation, error) {
	association := &model.UserRoleAssociation{UserID: userID}
	return association.Get()
}

func (dao *Dao) DeleteUserRoleAssociationsByUserID(userID uint) error {
	association := &model.UserRoleAssociation{UserID: userID}
	return association.Delete()
}
