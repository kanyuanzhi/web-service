package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
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
