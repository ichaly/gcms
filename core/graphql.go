package core

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/boot"
	"github.com/pkg/errors"
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

type gqlRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func NewGraphql(r *Render, e *boot.Engine, _ SchemaGroup) (*Graphql, error) {
	s, err := e.Schema()
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
