package core

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"golang.org/x/exp/maps"
	"net/http"
	"strings"
)

const (
	apiEndpoint = "/api"
)

type Graphql struct {
	render *Render
	schema graphql.Schema
}

type GraphGroup struct {
	fx.In
	Queries   []graphql.Fields `group:"query"`
	Mutations []graphql.Fields `group:"mutation"`
}

func NewGraphql(g GraphGroup, r *Render) (*Graphql, error) {
	queryFields := make(graphql.Fields)
	mutationFields := make(graphql.Fields)
	for _, q := range g.Queries {
		maps.Copy(queryFields, q)
	}
	for _, m := range g.Mutations {
		maps.Copy(mutationFields, m)
	}
	var config graphql.SchemaConfig
	if len(queryFields) > 0 {
		query := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}
		config.Query = graphql.NewObject(query)
	}
	if len(mutationFields) > 0 {
		mutation := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}
		config.Mutation = graphql.NewObject(mutation)
	}
	s, err := graphql.NewSchema(config)
	if err != nil {
		return nil, err
	}
	return &Graphql{schema: s, render: r}, nil
}

func (my *Graphql) Name() string {
	return "Graphql"
}

func (my *Graphql) Protected() bool {
	return false
}

func (my *Graphql) Init(r chi.Router) {
	r.Handle(apiEndpoint, my)
}

func (my *Graphql) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			_ = my.render.JSON(w, base.ERROR.WithError(err.(error)), WithCode(http.StatusBadRequest))
		}
	}()
	var req gqlRequest
	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()
		req.Query = query.Get("query")
		req.OperationName = query.Get("operationName")
		variables, ok := query["variables"]
		if ok {
			d := json.NewDecoder(strings.NewReader(variables[0]))
			d.UseNumber()
			if err := d.Decode(&req.Variables); err != nil {
				panic(errors.Wrap(err, "Not a valid GraphQL request body"))
			}
		}
	case http.MethodPost:
		d := json.NewDecoder(r.Body)
		d.UseNumber()
		if err := d.Decode(&req); err != nil {
			panic(errors.Wrap(err, "Not a valid GraphQL request body"))
		}
	default:
		panic(errors.New("Unrecognised request method.  Please use GET or POST for GraphQL requests"))
	}
	res := graphql.Do(graphql.Params{
		Context:        r.Context(),
		Schema:         my.schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
		OperationName:  req.OperationName,
	})
	_ = my.render.JSON(w, res)
}

type gqlRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}
