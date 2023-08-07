package auth

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/gcms/core"
	"net/http"
)

type Casbin struct {
	enforcer *casbin.Enforcer
}

func NewCasbin(e *casbin.Enforcer) (core.Plugin, error) {
	return &Casbin{enforcer: e}, nil
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
		sub := c.Request.Context().Value(core.UserContextKey)
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
