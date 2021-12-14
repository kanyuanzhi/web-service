package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	c *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		c: ctx,
	}
}

func (res *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res.c.JSON(http.StatusOK, data)
}
