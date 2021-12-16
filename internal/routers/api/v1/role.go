package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
	"github.com/kanyuanzhi/web-service/pkg/errcode"
)

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

func (r *Role) List(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	roles := svc.ListRoles()

	resData := model.NewSuccessResponse(roles)
	res.ToResponse(resData)
}

func (r *Role) Update(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updateRoleParam service.UpdateRoleRequest
	err := c.BindJSON(&updateRoleParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	role, err := svc.UpdateRole(&updateRoleParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
	}

	resData := model.NewSuccessResponse(role)
	res.ToResponse(resData)
}
