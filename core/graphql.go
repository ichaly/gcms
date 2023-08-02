package core

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/gcms/base"
	"go.uber.org/fx"
	"net/http"
)

const (
	apiEndpoint = "/graphql"
)

type Graphql struct {
	schema graphql.Schema
}

type SchemaGroup struct {
	fx.In
	All []base.Schema `group:"schema"`
}

type gqlRequest struct {
	Query     string                 `form:"query"`
	Operation string                 `form:"operationName" json:"operationName"`
	Variables map[string]interface{} `form:"variables"`
}

func NewGraphql(e *base.Engine, g SchemaGroup) (*Graphql, error) {
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

func (my *Graphql) Name() string {
	return "Graphql"
}

func (my *Graphql) Init(r *gin.RouterGroup) {
	r.Match([]string{http.MethodGet, http.MethodPost}, apiEndpoint, my.Handler)
}

func (my *Graphql) Protected() bool {
	return true
}

func (my *Graphql) Handler(c *gin.Context) {
	var req gqlRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errors": gqlerrors.FormatErrors(err)})
		return
	}
	res := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         my.schema,
		RequestString:  req.Query,
		OperationName:  req.Operation,
		VariableValues: req.Variables,
	})
	c.JSON(http.StatusOK, res)
}
