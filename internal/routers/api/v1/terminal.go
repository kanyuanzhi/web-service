package v1

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Terminal struct {
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (t *Terminal) Create(c *gin.Context) {
	// 接收从前端页面发送过来的终端基本信息
	msg := c.PostForm("message")
	log.Println(msg)

}
