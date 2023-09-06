package auth

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/gcms/base"
	"net/http"
)

type Casbin struct {
	enforcer *casbin.Enforcer
}

func NewCasbin(e *casbin.Enforcer) base.Plugin {
	return &Casbin{enforcer: e}
}

func (my *Casbin) Base() string {
	return "/"
}

func (my *Casbin) Init(r gin.IRouter) {
	r.Use(my.handler())
}

func (my *Casbin) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断策略中是否存在
		sub := c.Request.Context().Value(base.UserContextKey)
		ok, err := my.enforcer.Enforce(sub, c.Request.URL.RequestURI(), c.Request.Method)
		if err != nil {
			return
		}
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": gqlerrors.FormatErrors(errors.New("无权限"))})
			return
		}
		c.Next()
	}
}
