package base

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/gcms/core"
	"go.uber.org/fx"
	"net/http"
)

type Graphql struct {
	schema graphql.Schema
}

type SchemaGroup struct {
	fx.In
	All []core.Schema `group:"schema"`
}

type gqlRequest struct {
	Query     string                 `form:"query"`
	Operation string                 `form:"operationName" json:"operationName"`
	Variables map[string]interface{} `form:"variables"`
}

func NewGraphql(e *core.Engine, g SchemaGroup) (*Graphql, error) {
	for _, v := range g.All {
		err := e.Register(v)
		if err != nil {
			return nil, err
		}
	}
	s, err := e.Schema()
	if err != nil {
		return nil, err
	}
	return &Graphql{schema: s}, nil
}

func (my *Graphql) Base() string {
	return "/graphql"
}

func (my *Graphql) Init(r gin.IRouter) {
	r.Match([]string{http.MethodGet, http.MethodPost}, "/", my.handler)
}

func (my *Graphql) handler(c *gin.Context) {
	var req gqlRequest
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errors": gqlerrors.FormatErrors(err)})
		return
	}
	res := graphql.Do(graphql.Params{
		Context:        c.Request.Context(),
		Schema:         my.schema,
		RequestString:  req.Query,
		OperationName:  req.Operation,
		VariableValues: req.Variables,
	})
	c.JSON(http.StatusOK, res)
}
