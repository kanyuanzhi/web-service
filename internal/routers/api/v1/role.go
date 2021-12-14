package v1

import (
	"github.com/gin-gonic/gin"
)

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

func (r *Role) Get(c *gin.Context) {

}
