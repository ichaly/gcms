package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ichaly/gcms/base"
	"net/http"
)

type Index struct{}

func NewIndex() base.Plugin {
	return &Index{}
}

func (my *Index) Base() string {
	return "/"
}

func (my *Index) Init(r gin.IRouter) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
	})
}
