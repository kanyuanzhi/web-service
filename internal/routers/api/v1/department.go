package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
)

type Department struct {
	*model.DefaultFields
	Name         string `json:"name,omitempty"`
	Introduction string `json:"introduction,omitempty"`
}

func NewDepartment() *Department {
	return &Department{}
}

func (d *Department) List(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	departments := svc.ListDepartments()
	resData := model.NewSuccessResponse(departments)
	res.ToResponse(resData)
}

func (d *Department) Create(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	createDepartmentParam := &service.CreateDepartmentRequest{}
	err := c.BindJSON(createDepartmentParam)
	if err != nil {
		global.Log.Error(err)
		return
	}
	department := svc.CreateDepartment(createDepartmentParam)

	resData := model.NewSuccessResponse(department)
	res.ToResponse(resData)
}
