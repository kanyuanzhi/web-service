package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
	"github.com/kanyuanzhi/web-service/pkg/errcode"
	"os"
	"path/filepath"
	"time"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

// Get 获取单个用户信息
func (u *User) Get(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var userParam service.GetUserRequest
	err := c.BindQuery(&userParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	user, err := svc.GetUser(&userParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(user)
	res.ToResponse(resData)
}

// List 列出所有用户信息
func (u *User) List(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	users, err := svc.ListUsers()
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(users)
	res.ToResponse(resData)
}

// Register 用户注册
func (u *User) Register(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var registerParam service.RegisterRequest
	err := c.BindJSON(&registerParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	// 判断用户名是否已经注册
	isUsernameRepeated, err := svc.IsUsernameRepeated(registerParam.Username)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}
	if isUsernameRepeated == true {
		res.ToResponse(errcode.RepeatUsernameError)
		return
	}

	token, err := svc.Register(&registerParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(gin.H{"token": token})
	res.ToResponse(resData)
}

// UpdatePassword 更新用户密码
func (u *User) UpdatePassword(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updatePasswordParam service.UpdatePasswordRequest
	err := c.BindJSON(&updatePasswordParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	err = svc.UpdateAuthentication(&updatePasswordParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.AuthenticationFailError)
		return
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}

// UpdateAccount 更新用户信息
func (u *User) UpdateAccount(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updateUserAccountParam service.UpdateUserAccountRequest
	err := c.BindJSON(&updateUserAccountParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	user, err := svc.UpdateUserAccount(&updateUserAccountParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(user)
	res.ToResponse(resData)
}

// UpdateRoles 更新用户权限
func (u *User) UpdateRoles(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updateUserRoleAssociationsRequest service.UpdateUserRoleAssociationsRequest
	err := c.BindJSON(&updateUserRoleAssociationsRequest)
	if err != nil {
		global.Log.Error(err)
		return
	}

	if len(updateUserRoleAssociationsRequest.RoleNames) == 0 {
		res.ToResponse(errcode.EmptyRolesError)
		return
	}

	err = svc.UpdateUserRoleAssociations(&updateUserRoleAssociationsRequest)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}

// Delete 删除用户
func (u *User) Delete(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var deleteUserParam service.DeleteUserRequest
	err := c.BindQuery(&deleteUserParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	err = svc.DeleteUser(&deleteUserParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}

// Login 用户登录
func (u *User) Login(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var loginParam service.LoginRequest
	err := c.BindJSON(&loginParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	token, err := svc.Login(&loginParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.AuthenticationFailError)
		return
	}
	resData := model.NewSuccessResponse(gin.H{"token": token})
	res.ToResponse(resData)
}

// Logout 用户登出
func (u *User) Logout(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var logoutParam service.LogoutRequest
	err := c.BindQuery(&logoutParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	err = svc.Logout(&logoutParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}
	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}

func (u *User) UploadAvatar(c *gin.Context) {
	file, _ := c.FormFile("avatar")
	username := c.PostForm("username")
	dir, _ := os.Getwd()
	timestamp := time.Now().UnixMilli()
	avatarPath := filepath.Join(dir, "public/avatars", fmt.Sprintf("%s_%d.png", username, timestamp))
	err := c.SaveUploadedFile(file, avatarPath)
	if err != nil {
		global.Log.Error(err)
	}
	avatarNetPath := fmt.Sprintf("http://%s:%d/api/v1/user/avatar/%s_%d.png", global.Object.Host, global.Object.Port, username, timestamp)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": avatarNetPath,
	})
}

func (u *User) DownloadAvatar(c *gin.Context) {
	image := c.Param("image")
	dir, _ := os.Getwd()
	avatarPath := filepath.Join(dir, "public/avatars", image)
	c.File(avatarPath)
}
