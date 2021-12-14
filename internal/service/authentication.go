package service

import (
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

func (s *Service) FindAuthentication(param *LoginRequest) *model.Authentication {
	return s.dao.FindAuthentication(param.Username)
}

type CreateAuthenticationRequest struct {
	Username      string `json:"username" form:"username" `
	Password      string `json:"password" form:"password"`
	CheckPassword string `json:"check_password" form:"check_password"`
}

func (s *Service) CreateAuthentication(param *CreateAuthenticationRequest) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		global.Log.Error(err)
		return ""
	}
	password := string(hash)
	tokenStr := param.Username + strconv.FormatInt(time.Now().Unix(), 10)
	token := utils.MD5Str(tokenStr)
	s.dao.CreateAuthentication(token, param.Username, password)
	return token
}

type CountAuthenticationRequest struct {
	Username string `json:"username" form:"username" `
}

func (s *Service) CountAuthentication(param *CountAuthenticationRequest) int64 {
	return s.dao.CountAuthentication(param.Username)
}
