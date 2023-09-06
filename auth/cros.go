package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ichaly/gcms/base"
	"net/http"
)

type Cros struct{}

func NewCros() base.Plugin {
	return &Cros{}
}

func (my *Cros) Base() string {
	return "/"
}

func (my *Cros) Init(r gin.IRouter) {
	r.Use(my.handler)
}

// 处理跨域请求,支持options访问
func (my *Cros) handler(c *gin.Context) {
	origin := c.GetHeader("Origin")
	if len(origin) == 0 {
		c.Next()
		return
	}

	// 同源直接过
	host := c.GetHeader("Host")
	if origin == "http://"+host || origin == "https://"+host {
		c.Next()
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	// OPTIONS 过
	method := c.Request.Method
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		c.Abort()
	}
	c.Next()
}