package boot

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

type Engine struct {
	Names map[string]reflect.Type
	Types map[string]graphql.Type
}

func NewEngine() *Engine {
	return &Engine{
		Names: map[string]reflect.Type{},
		Types: map[string]graphql.Type{
			Query: q, Mutation: m, Subscription: s,
		},
	}
}

func (my *Engine) Register(prototype interface{}) (*graphql.Object, error) {
	typ := reflect.TypeOf(prototype)
	obj, err := my.buildObject(typ)
	if err != nil {
		my.Names[obj.Name()] = typ
	}
	return obj, err
}

func (my *Engine) Schema() (graphql.Schema, error) {
	config := graphql.SchemaConfig{}
	if len(q.Fields()) > 0 {
		config.Query = q
	}
	if len(m.Fields()) > 0 {
		config.Mutation = m
	}
	if len(s.Fields()) > 0 {
		config.Subscription = s
	}
	return graphql.NewSchema(config)
}
