package auth

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/util"
	"net/http"
)

type Graphql struct {
	enforcer *casbin.Enforcer
}

func NewGraphql(e *casbin.Enforcer) base.Plugin {
	return &Graphql{enforcer: e}
}

func (my *Graphql) Base() string {
	return "/graphql"
}

func (my *Graphql) Init(r gin.IRouter) {
	r.Use(my.handler)
}

func (my *Graphql) handler(c *gin.Context) {
	var req = struct {
		Query string `form:"query"`
	}{}
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err)})
		return
	}
	doc, _ := parser.Parse(parser.ParseParams{Source: req.Query})
	sub, _ := c.Request.Context().Value(base.UserContextKey).(uint64)
	for _, node := range doc.Definitions {
		switch d := node.(type) {
		case ast.TypeSystemDefinition:
			for _, s := range d.GetSelectionSet().Selections {
				switch f := s.(type) {
				case *ast.Field:
					ok, err := my.enforcer.Enforce(util.FormatLong(int64(sub)), f.Name.Value, d.GetOperation())
					if err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err)})
						return
					}
					if !ok {
						c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": gqlerrors.FormatErrors(errors.New("无权限"))})
						return
					}
				}
			}
		}
	}
	c.Next()
}
