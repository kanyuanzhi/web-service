package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
	"github.com/kanyuanzhi/web-service/pkg/errcode"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) Get(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var userParam service.GetUserRequest
	err := c.BindQuery(&userParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	user := svc.GetUser(&userParam)

	userRoleParam := service.GetUserRoleRequest{UserID: user.ID}
	userRoles := svc.GetUserRoles(&userRoleParam)
	var roles []string
	if len(userRoles) == 0 {
		roles = append(roles, "guest")
	} else {
		for _, role := range userRoles {
			roles = append(roles, role.RoleName)
		}
	}

	user.Roles = roles
	user.Departments = []uint{}

	resData := model.NewSuccessResponse(user)
	res.ToResponse(resData)
}

func (u *User) List(c *gin.Context) {
	log.Println(121312421)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (u *User) Register(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var createAuthParam service.CreateAuthenticationRequest
	err := c.BindJSON(&createAuthParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	countAuthParam := service.CountAuthenticationRequest{Username: createAuthParam.Username}

	count := svc.CountAuthentication(&countAuthParam) // 判断用户名是否已经注册
	if count != 0 {
		res.ToResponse(errcode.RepeatUsernameError)
		return
	}

	token := svc.CreateAuthentication(&createAuthParam)
	createUserParam := service.CreateUserRequest{Token: token, Username: createAuthParam.Username}
	userID := svc.CreateUser(&createUserParam)

	createUserRoleParam := service.CreateUserRoleRequest{UserID: userID, RoleName: "guest"}
	svc.CreateUserRole(&createUserRoleParam)

	resData := model.NewSuccessResponse(gin.H{"token": token})
	res.ToResponse(resData)
}

func (u *User) Update(c *gin.Context) {

}

func (u *User) Delete(c *gin.Context) {

}

func (u *User) Login(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var loginParam service.LoginRequest
	err := c.BindJSON(&loginParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	auth := svc.FindAuthentication(&loginParam)
	if auth == nil {
		res.ToResponse(errcode.AuthenticationFailError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(loginParam.Password))
	if err != nil {
		res.ToResponse(errcode.AuthenticationFailError)
		return
	}
	resData := model.NewSuccessResponse(gin.H{"token": auth.Token})
	res.ToResponse(resData)
}

func (u *User) Logout(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var logoutParam service.LogoutRequest
	err := c.BindQuery(&logoutParam)
	if err != nil {
		global.Log.Error(err)
		return
	}
	svc.Logout(&logoutParam)

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
