package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
	"github.com/kanyuanzhi/web-service/pkg/errcode"
)

type Terminal struct{}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (t *Terminal) Create(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var createTerminalParam service.CreateTerminalRequest
	err := c.BindJSON(&createTerminalParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	terminal, err := svc.CreateTerminal(&createTerminalParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
	}

	resData := model.NewSuccessResponse(terminal)
	res.ToResponse(resData)
}

func (t *Terminal) List(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	terminals, err := svc.ListTerminals()
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
	}
	resData := model.NewSuccessResponse(terminals)
	res.ToResponse(resData)
}

func (t *Terminal) Update(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updateTerminalParam service.UpdateTerminalRequest
	err := c.BindJSON(&updateTerminalParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	terminal, err := svc.UpdateTerminal(&updateTerminalParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
	}

	resData := model.NewSuccessResponse(terminal)
	res.ToResponse(resData)
}

func (t *Terminal) Delete(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var deleteTerminalParam service.DeleteTerminalRequest
	err := c.BindQuery(&deleteTerminalParam)
	if err != nil {
		global.Log.Error(err)
		return
	}

	err = svc.DeleteTerminal(&deleteTerminalParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}
