package core

import (
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"reflect"
)

type Engine struct {
	types        map[string]graphql.Type
	fieldParsers []fieldParser
}

func NewEngine() (*Engine, error) {
	my := &Engine{types: map[string]graphql.Type{
		Query: q, Mutation: m, Subscription: s,
	}}
	my.fieldParsers = []fieldParser{
		my.asBuiltinScalarField,
		my.asCustomScalarField,
		my.asIdField,
		my.asEnumField,
		my.asObjectField,
	}
	return my, nil
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

func (my *Engine) AddTo(source interface{}, target string) error {
	info, err := unwrap(reflect.TypeOf(source))
	if err != nil {
		return err
	}

	val, ok := my.types[target]
	if !ok {
		return errors.New("target type not registered")
	}

	obj, ok := val.(*graphql.Object)
	if !ok {
		return errors.New("source type not an object")
	}

	node, err := my.buildObject(&info)
	if err != nil {
		return err
	}

	queryArgs := graphql.FieldConfigArgument{
		"id":         {Type: graphql.ID},
		"size":       {Type: graphql.Int},
		"page":       {Type: graphql.Int},
		"top":        {Type: graphql.Int},
		"last":       {Type: graphql.Int},
		"after":      {Type: Cursor},
		"before":     {Type: Cursor},
		"search":     {Type: graphql.String},
		"distinctOn": {Type: graphql.NewList(graphql.String)},
		"sort":       {Type: my.types[node.Name()+"SortInput"]},
		"where":      {Type: my.types[node.Name()+"WhereInput"]},
	}
	obj.AddFieldConfig(strcase.ToLowerCamel(node.Name()), &graphql.Field{
		Type: node, Args: queryArgs, Description: node.Description(),
	})
	return nil
}

func (my *Engine) AddQuery(source interface{}) error {
	return my.AddTo(source, Query)
}

func (my *Engine) AddMutation(source interface{}) error {
	return my.AddTo(source, Mutation)
}

func (my *Engine) AddSubscription(source interface{}) error {
	return my.AddTo(source, Subscription)
}
