package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct{}

func NewIndex() (*Index, error) {
	return &Index{}, nil
}

func (my *Index) Base() string {
	return "/"
}

func (my *Index) Init(r gin.IRouter) {
	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
	})
}
