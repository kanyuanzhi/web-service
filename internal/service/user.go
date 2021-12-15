package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type LogoutRequest struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) Logout(param *LogoutRequest) {
	// todo: record logout event
}

type CreateUserRequest struct {
	Token    string `json:"token" form:"token"`
	Username string `json:"username" form:"username"`
}

func (s *Service) CreateUser(param *CreateUserRequest) uint {
	return s.dao.CreateUser(param.Token, param.Username)
}

type GetUserRequest struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) GetUser(param *GetUserRequest) *model.User {
	return s.dao.GetUser(param.Token)
}

func (s *Service) ListUsers() []*model.User {
	return s.dao.ListUsers()
}

type UpdateUserAccountRequest struct {
	// UserAccount包括两部分，用户基本表users和用户部门表users_departments
	// UpdateUser
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Contact      string `json:"contact"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar,omitempty"` //用户自己更新时avatar有值，管理员更新时avatar无值

	// UpdateUserDepartments(delete and create)
	Departments []uint `json:"departments"`
}

func (s *Service) UpdateUser(param *UpdateUserAccountRequest) *model.User {
	return s.dao.UpdateUser(param.ID, param.Name, param.Contact, param.Introduction, param.Avatar)
}

type UpdateUserDepartmentsRequest struct {
	UserID        uint   `json:"user_id" form:"user_id"`
	DepartmentIDs []uint `json:"departments" form:"departments"`
}

func (s *Service) UpdateUserDepartments(param *UpdateUserDepartmentsRequest) []*model.UserDepartment {
	deleteUserDepartmentsParam := DeleteUserDepartmentsRequest{UserID: param.UserID}
	s.DeleteUserDepartments(&deleteUserDepartmentsParam)
	if len(param.DepartmentIDs) == 0 {
		return nil
	} else {
		createUserDepartmentsParam := CreateUserDepartmentsRequest{UserID: param.UserID, DepartmentIDs: param.DepartmentIDs}
		return s.CreateUserDepartments(&createUserDepartmentsParam)
	}
}

type UpdateUserRolesRequest struct {
	// UpdateUserRoles
	UserID    uint     `json:"id"`
	RoleNames []string `json:"roles"`
}

func (s *Service) UpdateUserRoles(param *UpdateUserRolesRequest) []*model.UserRole {
	deleteUserRolesRequest := DeleteUserRolesRequest{UserID: param.UserID}
	s.DeleteUserRoles(&deleteUserRolesRequest)
	createUserRolesRequest := CreateUserRolesRequest{UserID: param.UserID, RoleNames: param.RoleNames}
	return s.CreateUserRoles(&createUserRolesRequest)
}
