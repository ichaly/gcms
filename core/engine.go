package core

import (
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"reflect"
)

type Engine struct {
	types        map[string]graphql.Type
	fieldParsers []fieldParser
}

func NewEngine() (*Engine, error) {
	my := &Engine{types: map[string]graphql.Type{
		"Query": q, "Mutation": m, "Subscription": s,
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
	name := info.baseType.Name()

	_, ok := my.types[name]
	if ok {
		return errors.New("source type already registered")
	}

	val, ok := my.types[target]
	if !ok {
		return errors.New("target type not registered")
	}

	obj, ok := val.(*graphql.Object)
	if !ok {
		return errors.New("source type not an object")
	}

	node, err := my.parseObject(&info)
	if err != nil {
		return err
	}
	keys := maps.Keys(node.Fields())
	sortFields := graphql.InputObjectConfigFieldMap{}
	for _, k := range keys {
		sortFields[k] = &graphql.InputObjectFieldConfig{Type: SortDirection}
	}
	queryFields := graphql.FieldConfigArgument{
		"id":         {Type: graphql.ID},
		"size":       {Type: graphql.Int},
		"page":       {Type: graphql.Int},
		"top":        {Type: graphql.Int},
		"last":       {Type: graphql.Int},
		"search":     {Type: graphql.String},
		"after":      {Type: Cursor},
		"before":     {Type: Cursor},
		"distinctOn": {Type: graphql.NewList(graphql.String)},
		"sort": {Type: graphql.NewInputObject(graphql.InputObjectConfig{
			Name: node.Name() + "SortInput", Fields: sortFields,
		})},
		"where": {Type: graphql.NewInputObject(graphql.InputObjectConfig{
			Name: node.Name() + "WhereInput", Fields: sortFields,
		})},
	}
	obj.AddFieldConfig(strcase.ToLowerCamel(node.Name()), &graphql.Field{
		Type: node, Args: queryFields, Description: node.Description(),
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
