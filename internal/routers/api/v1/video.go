package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/service"
	"github.com/kanyuanzhi/web-service/pkg/app"
	"github.com/kanyuanzhi/web-service/pkg/errcode"
)

type Video struct{}

func NewVideo() *Video {
	return &Video{}
}

// List 列出所有监控
func (tg *Video) List(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	videos, err := svc.ListVideos()
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(videos)
	res.ToResponse(resData)
}

// Create 创建新监控
func (tg *Video) Create(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var createVideoParam service.CreateVideoRequest
	err := c.BindJSON(&createVideoParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	video, err := svc.CreateVideo(&createVideoParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(video)
	res.ToResponse(resData)
}

// Update 更新监控信息
func (tg *Video) Update(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var updateVideoParam service.UpdateVideoRequest
	err := c.BindJSON(&updateVideoParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	video, err := svc.UpdateVideo(&updateVideoParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(video)
	res.ToResponse(resData)
}

// Delete 删除监控
func (tg *Video) Delete(c *gin.Context) {
	res := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var deleteVideoParam service.DeleteVideoRequest
	err := c.BindQuery(&deleteVideoParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	err = svc.DeleteVideo(&deleteVideoParam)
	if err != nil {
		global.Log.Error(err)
		res.ToResponse(errcode.ServerError)
		return
	}

	resData := model.NewSuccessResponse(nil)
	res.ToResponse(resData)
}
