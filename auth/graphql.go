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
	"net/http"
	"strings"
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
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	var req = struct {
		Query string `form:"query"`
	}{}
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		panic(err)
	}
	doc, _ := parser.Parse(parser.ParseParams{Source: req.Query})
	sub, _ := c.Request.Context().Value(base.UserContextKey).(string)
	for _, node := range doc.Definitions {
		switch d := node.(type) {
		case *ast.OperationDefinition:
			if d.Name == nil {
				panic(errors.New("must have operation name"))
			}
			if ok, err := my.enforcer.Enforce(sub, d.Name.Value, d.GetOperation()); err != nil {
				panic(err)
			} else if !ok {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": gqlerrors.FormatErrors(errors.New("无权限"))})
				return
			}
		}
	}
	c.Next()
}

func unfoldSelection(set *ast.SelectionSet, prefix ...string) (res []string) {
	if len(prefix) > 0 {
		res = append(res, strings.Join(prefix, "."))
	}
	for _, s := range set.Selections {
		switch f := s.(type) {
		case *ast.Field:
			nodes := append(prefix, f.Name.Value)
			if f.GetSelectionSet() == nil {
				res = append(res, strings.Join(nodes, "."))
			} else {
				res = append(res, unfoldSelection(f.GetSelectionSet(), nodes...)...)
			}
		}
	}
	return
}
