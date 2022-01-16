package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
	"github.com/kanyuanzhi/web-service/pkg/errcode"
)

type TerminalGroup struct{}

func NewTerminalGroup() *TerminalGroup {
	return &TerminalGroup{}
}

func (tg *TerminalGroup) List(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	terminalGroups, err := svc.ListTerminalGroups()
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}
	resData := model.NewSuccessResponse(terminalGroups)
	res.ToResponse(resData)
}

func (tg *TerminalGroup) Create(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var createTerminalGroupParam service.CreateTerminalGroupRequest
	err := c.BindJSON(&createTerminalGroupParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	terminalGroup, err := svc.CreateTerminalGroup(&createTerminalGroupParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(terminalGroup)
	res.ToResponse(resData)
}

func (tg *TerminalGroup) Update(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updateTerminalGroupParam service.UpdateTerminalGroupRequest
	err := c.BindJSON(&updateTerminalGroupParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	terminalGroup, err := svc.UpdateTerminalGroup(&updateTerminalGroupParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(terminalGroup)
	res.ToResponse(resData)
}

func (tg *TerminalGroup) UpdateMembers(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var UpdateTerminalGroupAssociationsParam service.UpdateTerminalGroupAssociationsRequest
	err := c.BindJSON(&UpdateTerminalGroupAssociationsParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	err = svc.UpdateTerminalGroupAssociations(&UpdateTerminalGroupAssociationsParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}

func (tg *TerminalGroup) Delete(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var deleteTerminalGroupParam service.DeleteTerminalGroupRequest
	err := c.BindQuery(&deleteTerminalGroupParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	err = svc.DeleteTerminalGroup(&deleteTerminalGroupParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}
