package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type Casbin struct {
	enforcer *casbin.Enforcer
}

func NewCasbin(e *casbin.Enforcer) (*Casbin, error) {
	return &Casbin{enforcer: e}, nil
}

func (my *Casbin) Name() string {
	return "Casbin"
}

func (my *Casbin) Init(r *gin.RouterGroup) {
	r.Group("/api").Use(my.handler())
}

func (my *Casbin) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断策略中是否存在
		ok, err := my.enforcer.Enforce("admin", c.Request.URL.RequestURI(), c.Request.Method)
		if err != nil {
			return
		}
		if !ok {
			c.Abort()
			return
		}
		c.Next()
	}
}
