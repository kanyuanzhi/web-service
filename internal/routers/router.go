package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/web-service/docs"
	"github.com/kanyuanzhi/web-service/internal/middleware"
	"github.com/kanyuanzhi/web-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	user := v1.NewUser()
	role := v1.NewRole()
	department := v1.NewDepartment()

	terminal := v1.NewTerminal()

	docs.SwaggerInfo.BasePath = "/api/v1"
	apiv1 := r.Group("/api/v1")
	{
		// finished
		apiv1.POST("/user/login", user.Login)
		apiv1.POST("/user/logout", user.Logout)
		apiv1.POST("/user/register", user.Register)
		apiv1.GET("/user", user.Get)
		apiv1.GET("/users", user.List)
		apiv1.PUT("/user/account", user.UpdateAccount)
		apiv1.PUT("/user/roles", user.UpdateRoles)
		apiv1.PUT("/user/password", user.UpdatePassword)
		apiv1.DELETE("/user", user.Delete)

		apiv1.POST("/department", department.Create)
		apiv1.GET("/departments", department.List)
		apiv1.PUT("/department", department.Update)
		apiv1.DELETE("/department", department.Delete)

		apiv1.GET("/roles", role.List)
		apiv1.PUT("/role", role.Update)

		apiv1.POST("/terminal", terminal.Create)
		apiv1.GET("/terminals", terminal.List)
		apiv1.PUT("/terminal", terminal.Update)
		apiv1.DELETE("/terminal", terminal.Delete)

		// unfinished
		apiv1.POST("/user/avatar", user.UploadAvatar)
		apiv1.GET("/user/avatar/:image", user.DownloadAvatar)
		//apiv1.PUT("/user/:token", user.Update)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
