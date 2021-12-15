package service

import (
	"errors"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/utils"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (s *Service) Login(param *LoginRequest) *model.Authentication {
	return s.dao.FindAuthenticationByUsername(param.Username)
}

type RegisterRequest struct {
	Username      string `json:"username" form:"username" `
	Password      string `json:"password" form:"password"`
	CheckPassword string `json:"check_password" form:"check_password"`
}

func (s *Service) Register(param *RegisterRequest) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		global.Log.Error(err)
		return ""
	}
	password := string(hash)
	tokenStr := param.Username + strconv.FormatInt(time.Now().Unix(), 10)
	token := utils.MD5Str(tokenStr)
	s.dao.CreateAuthentication(token, param.Username, password)

	createUserParam := CreateUserRequest{Token: token, Username: param.Username}
	userID := s.CreateUser(&createUserParam)

	createUserRolesParam := CreateUserRolesRequest{UserID: userID, RoleNames: []string{"guest"}} // 关联默认权限
	s.CreateUserRoles(&createUserRolesParam)

	return token
}

type UpdatePasswordRequest struct {
	Token            string `json:"token" form:"token"`
	OldPassword      string `json:"old_password" form:"old_password"`
	NewPassword      string `json:"new_password" form:"new_password"`
	CheckNewPassword string `json:"check_new_password" form:"check_new_password"`
}

func (s *Service) UpdateAuthentication(param *UpdatePasswordRequest) error{
	auth := s.dao.FindAuthenticationByToken(param.Token)
	if auth == nil{
		return errors.New("wrong token")
	}
	err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(param.OldPassword))
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(param.NewPassword), bcrypt.DefaultCost)
	NewPassword := string(hash)
	return s.dao.UpdateAuthentication(param.Token, NewPassword)
}

type CountAuthenticationRequest struct {
	Username string `json:"username" form:"username" `
}

func (s *Service) CountAuthentication(param *CountAuthenticationRequest) int64 {
	return s.dao.CountAuthentication(param.Username)
}
