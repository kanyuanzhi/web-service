package service

import (
	"errors"
	"github.com/kanyuanzhi/web-service/utils"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (s *Service) Login(param *LoginRequest) (string, error) {
	auth, err := s.dao.FindAuthenticationByUsername(param.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(param.Password))
	if err != nil {
		return "", err
	}
	return auth.Token, nil
}

type RegisterRequest struct {
	Username      string `json:"username" form:"username" `
	Password      string `json:"password" form:"password"`
	CheckPassword string `json:"check_password" form:"check_password"`
}

func (s *Service) Register(param *RegisterRequest) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	password := string(hash)
	tokenStr := param.Username + strconv.FormatInt(time.Now().Unix(), 10)
	token := utils.MD5Str(tokenStr)
	_, err = s.dao.CreateAuthentication(token, param.Username, password)
	if err != nil {
		return "", err
	}

	user, err := s.dao.CreateUser(token, param.Username)
	if err != nil {
		return "", err
	}

	_, err = s.dao.CreateUserRoleAssociations(user.ID, []string{"guest"})
	if err != nil {
		return "", err
	}

	return token, nil
}

type UpdatePasswordRequest struct {
	Token            string `json:"token" form:"token"`
	OldPassword      string `json:"old_password" form:"old_password"`
	NewPassword      string `json:"new_password" form:"new_password"`
	CheckNewPassword string `json:"check_new_password" form:"check_new_password"`
}

func (s *Service) UpdateAuthentication(param *UpdatePasswordRequest) error {
	auth, err := s.dao.FindAuthenticationByToken(param.Token)
	if err != nil {
		return errors.New("wrong token")
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(param.OldPassword))
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(param.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	NewPassword := string(hash)
	return s.dao.UpdateAuthenticationByToken(param.Token, NewPassword)
}

func (s *Service) IsUsernameRepeated(username string) (bool, error) {
	count, err := s.dao.CountAuthentication(username)
	if count == 0 {
		return false, err
	} else {
		return true, err
	}
}
