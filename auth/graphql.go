package auth

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/ichaly/gcms/base"
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
	var rb = struct {
		Query string `form:"query"`
	}{}
	err := c.ShouldBind(&rb)
	if err != nil {
		return
	}
	doc, _ := parser.Parse(parser.ParseParams{Source: rb.Query})
	for _, node := range doc.Definitions {
		switch d := node.(type) {
		case ast.TypeSystemDefinition:
			for _, s := range d.GetSelectionSet().Selections {
				switch f := s.(type) {
				case *ast.Field:
					sub := c.Request.Context().Value(base.UserContextKey)
					ok, err := my.enforcer.Enforce(sub, f.Name.Value, d.GetOperation())
					if err != nil {
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
