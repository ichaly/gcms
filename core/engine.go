package core

import (
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"reflect"
)

type Engine struct {
	config       graphql.SchemaConfig
	types        map[reflect.Type]graphql.Type
	fieldParsers []fieldParser
}

func NewEngine() (*Engine, error) {
	my := &Engine{
		config: graphql.SchemaConfig{
			Types: []graphql.Type{Cursor, Void, OrderDirection},
		},
		types: map[reflect.Type]graphql.Type{
			_queryType: q, _mutationType: m, _subscriptionType: s,
		},
	}
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
	if len(q.Fields()) > 0 {
		my.config.Query = q
	}
	if len(m.Fields()) > 0 {
		my.config.Mutation = m
	}
	if len(s.Fields()) > 0 {
		my.config.Subscription = s
	}
	return graphql.NewSchema(my.config)
}

func (my *Engine) AddTo(source interface{}, target reflect.Type) error {
	src := reflect.TypeOf(source)
	_, ok := my.types[src]
	if ok {
		return errors.New("source type already registered")
	}

	if src == target {
		return errors.New("source and target are the same")
	}

	val, ok := my.types[target]
	if !ok {
		return errors.New("target type not registered")
	}

	obj, ok := val.(*graphql.Object)
	if !ok {
		return errors.New("source type not an object")
	}

	info, err := unwrap(src)
	if err != nil {
		return err
	}
	node, err := my.parseObject(&info)
	if err != nil {
		return err
	}
	name := strcase.ToLowerCamel(node.Name())
	keys := maps.Keys(node.Fields())
	fields := graphql.InputObjectConfigFieldMap{}
	for _, k := range keys {
		fields[k] = &graphql.InputObjectFieldConfig{Type: OrderDirection}
	}
	orderByType := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: node.Name() + "OrderByInput", Fields: fields,
	})
	obj.AddFieldConfig(name, &graphql.Field{
		Type: node, Description: node.Description(),
		Args: map[string]*graphql.ArgumentConfig{
			"id":         {Type: graphql.ID},
			"limit":      {Type: graphql.Int},
			"offset":     {Type: graphql.Int},
			"first":      {Type: graphql.Int},
			"last":       {Type: graphql.Int},
			"after":      {Type: Cursor},
			"before":     {Type: Cursor},
			"search":     {Type: graphql.String},
			"orderBy":    {Type: orderByType},
			"where":      {Type: graphql.String},
			"distinctOn": {Type: graphql.NewList(graphql.String)},
		},
	})
	return nil
}

func (my *Engine) AddQuery(source interface{}) error {
	return my.AddTo(source, _queryType)
}

func (my *Engine) AddMutation(source interface{}) error {
	return my.AddTo(source, _mutationType)
}

func (my *Engine) AddSubscription(source interface{}) error {
	return my.AddTo(source, _subscriptionType)
}
