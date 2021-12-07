package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/docs"
	"github.com/kanyuanzhi/web-service/internal/routers/apis"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := apis.NewTag()
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", tag.Get)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
