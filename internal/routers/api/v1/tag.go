package v1

import (
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() *Tag {
	return &Tag{}
}

// @BasePath /api/v1

// Get PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /tags [get]
func (t *Tag) Get(c *gin.Context) {

}
