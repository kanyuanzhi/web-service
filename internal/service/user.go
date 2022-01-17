package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type LogoutRequest struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) Logout(param *LogoutRequest) error {
	// todo: record logout event
	return nil
}

type GetUserRequest struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) GetUser(param *GetUserRequest) (*model.User, error) {
	user, err := s.dao.GetUser(param.Token)
	if err != nil {
		return nil, err
	}

	// todo 待优化
	// 关联roles
	userRoleAssociations, err := s.dao.GetUserRoleAssociationsByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	roleNames := []string{}
	for _, association := range userRoleAssociations {
		roleNames = append(roleNames, association.RoleName)
	}
	user.RoleNames = roleNames

	// 关联departments
	userDepartmentAssociations, err := s.dao.GetUserDepartmentAssociations(user.ID)
	if err != nil {
		return nil, err
	}
	departmentIDs := []uint{}
	for _, association := range userDepartmentAssociations {
		departmentIDs = append(departmentIDs, association.DepartmentID)
	}
	user.DepartmentIDs = departmentIDs

	return user, nil
}

func (s *Service) ListUsers() ([]*model.User, error) {
	users, err := s.dao.ListUsers()
	if err != nil {
		return nil, err
	}
	// todo 待优化
	for _, user := range users {
		// 关联roles
		userRoleAssociations, _ := s.dao.GetUserRoleAssociationsByUserID(user.ID)
		roleNames := []string{}
		for _, association := range userRoleAssociations {
			roleNames = append(roleNames, association.RoleName)
		}
		user.RoleNames = roleNames

		// 关联departments
		userDepartmentAssociations, _ := s.dao.GetUserDepartmentAssociations(user.ID)
		departmentIDs := []uint{}
		for _, association := range userDepartmentAssociations {
			departmentIDs = append(departmentIDs, association.DepartmentID)
		}
		user.DepartmentIDs = departmentIDs
	}
	return users, nil
}

type UpdateUserAccountRequest struct {
	// UserAccount包括两部分，用户基本表users和用户部门关联表user_department_associations
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Contact       string `json:"contact"`
	Introduction  string `json:"introduction"`
	Avatar        string `json:"avatar,omitempty"` //用户自己更新时avatar有值，管理员更新时avatar无值
	DepartmentIDs []uint `json:"departments" form:"departments"`
}

func (s *Service) UpdateUserAccount(param *UpdateUserAccountRequest) (*model.User, error) {
	// 更新users表
	user, err := s.dao.UpdateUser(param.ID, param.Name, param.Contact, param.Introduction, param.Avatar)
	if err != nil {
		return nil, err
	}

	// 更新user_department_associations
	err = s.dao.DeleteUserDepartmentAssociationsByUserID(param.ID)
	if err != nil {
		return nil, err
	}
	if len(param.DepartmentIDs) != 0 {
		_, err = s.dao.CreateUserDepartmentAssociations(param.ID, param.DepartmentIDs)
		if err != nil {
			return nil, err
		}
	}
	user.DepartmentIDs = param.DepartmentIDs

	return user, nil
}

type UpdateUserRoleAssociationsRequest struct {
	UserID    uint     `json:"id"`
	RoleNames []string `json:"roles"`
}

func (s *Service) UpdateUserRoleAssociations(param *UpdateUserRoleAssociationsRequest) error {
	err := s.dao.DeleteUserRoleAssociationsByUserID(param.UserID)
	if err != nil {
		return err
	}
	_, err = s.dao.CreateUserRoleAssociations(param.UserID, param.RoleNames)
	if err != nil {
		return err
	}
	return nil
}

type DeleteUserRequest struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) DeleteUser(param *DeleteUserRequest) error {
	user, err := s.dao.GetUser(param.Token)
	if err != nil {
		return err
	}

	//删除user表中的记录
	err = s.dao.DeleteUserByToken(param.Token)
	if err != nil {
		return err
	}

	//删除authentication表中的记录
	err = s.dao.DeleteAuthenticationByToken(param.Token)
	if err != nil {
		return err
	}

	//删除UserRoleAssociation表中的记录
	err = s.dao.DeleteUserRoleAssociationsByUserID(user.ID)
	if err != nil {
		return err
	}

	//删除UserDepartmentAssociation表中的记录
	err = s.dao.DeleteUserDepartmentAssociationsByUserID(user.ID)
	if err != nil {
		return err
	}

	return nil
}
