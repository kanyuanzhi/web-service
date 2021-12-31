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

	// 关联roles
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
	if roles == nil {
		user.Roles = []string{}
	} else {
		user.Roles = roles
	}

	// 关联departments
	userDepartmentParam := service.GetUserDepartmentsRequest{UserID: user.ID}
	userDepartments := svc.GetUserDepartments(&userDepartmentParam)
	var departments []uint
	for _, userDepartment := range userDepartments {
		departments = append(departments, userDepartment.DepartmentID)
	}
	if departments == nil {
		user.Departments = []uint{}
	} else {
		user.Departments = departments
	}

	resData := model.NewSuccessResponse(user)
	res.ToResponse(resData)
}

func (u *User) List(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	users := svc.ListUsers()

	for _, user := range users {
		// 根据userID查询对应权限
		userRoleParam := service.GetUserRoleRequest{UserID: user.ID}
		userRoles := svc.GetUserRoles(&userRoleParam)
		var roles []string
		for _, userRole := range userRoles {
			roles = append(roles, userRole.RoleName)
		}
		if roles == nil {
			// 令json格式在用户没有任何权限时，departments字段为[]，否则为null
			user.Roles = []string{}
		} else {
			user.Roles = roles
		}
		// 根据userID查询对应部门
		userDepartmentParam := service.GetUserDepartmentsRequest{UserID: user.ID}
		userDepartments := svc.GetUserDepartments(&userDepartmentParam)
		var departments []uint
		for _, userDepartment := range userDepartments {
			departments = append(departments, userDepartment.DepartmentID)
		}
		if departments == nil {
			// 令json格式在用户没有加入任何部门时，departments字段为[]，否则为null
			user.Departments = []uint{}
		} else {
			user.Departments = departments
		}
	}

	resData := model.NewSuccessResponse(users)
	res.ToResponse(resData)
}

func (u *User) Register(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var registerParam service.RegisterRequest
	err := c.BindJSON(&registerParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	countAuthParam := service.CountAuthenticationRequest{Username: registerParam.Username}

	count := svc.CountAuthentication(&countAuthParam) // 判断用户名是否已经注册
	if count != 0 {
		res.ToResponse(errcode.RepeatUsernameError)
		return
	}

	token := svc.Register(&registerParam)

	resData := model.NewSuccessResponse(gin.H{"token": token})
	res.ToResponse(resData)
}

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

func (u *User) UpdateAccount(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	// 更新基本信息
	var updateUserAccountParam service.UpdateUserAccountRequest
	err := c.BindJSON(&updateUserAccountParam)
	if err != nil {
		global.Log.Error(err)
		return
	}
	user := svc.UpdateUser(&updateUserAccountParam)

	// 更新部门
	updateUserDepartmentRequest := service.UpdateUserDepartmentsRequest{UserID: updateUserAccountParam.ID, DepartmentIDs: updateUserAccountParam.Departments}
	userDepartments := svc.UpdateUserDepartments(&updateUserDepartmentRequest)
	if userDepartments == nil {
		user.Departments = []uint{}
	} else {
		var departments []uint
		for _, userDepartment := range userDepartments {
			departments = append(departments, userDepartment.DepartmentID)
		}
		user.Departments = departments
	}

	resData := model.NewSuccessResponse(user)
	res.ToResponse(resData)
}

func (u *User) UpdateRoles(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updateUserRolesRequest service.UpdateUserRolesRequest
	err := c.BindJSON(&updateUserRolesRequest)
	if err != nil {
		global.Log.Error(err)
		return
	}

	if len(updateUserRolesRequest.RoleNames) == 0 {
		res.ToResponse(errcode.EmptyRolesError)
		return
	}

	userRoles := svc.UpdateUserRoles(&updateUserRolesRequest)
	var roles []string
	for _, userRole := range userRoles {
		roles = append(roles, userRole.RoleName)
	}
	resData := model.NewSuccessResponse(gin.H{"roles": roles})
	res.ToResponse(resData)
}

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
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
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

	auth, err := svc.Login(&loginParam)
	if err == nil {
		global.Log.Error(err)
		res.ToResponse(errcode.AuthenticationFailError)
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(loginParam.Password))
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.AuthenticationFailError)
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
