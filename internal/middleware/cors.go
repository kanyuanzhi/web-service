package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"HEAD", "GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"*"},
		//ExposedHeaders:   []string{"*"},
		Debug: false,
	})
}
