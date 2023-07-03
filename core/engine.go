package core

import (
	"fmt"
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
	my.Expressions()
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

	node, err := my.parseObject(&info)
	if err != nil {
		return err
	}
	keys := maps.Keys(node.Fields())

	sortFields := graphql.InputObjectConfigFieldMap{}
	for _, k := range keys {
		sortFields[k] = &graphql.InputObjectFieldConfig{Type: SortDirection}
	}
	sortInput := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: node.Name() + "SortInput", Fields: sortFields,
	})

	whereFields := graphql.InputObjectConfigFieldMap{}
	for k, v := range node.Fields() {
		t := v.Type
		suffix := "Expression"
		if val, ok := t.(*graphql.NonNull); ok {
			t = val.OfType
		}
		if val, ok := t.(*graphql.List); ok {
			t = val.OfType
			suffix = "List" + suffix
		}
		if val, ok := t.(*graphql.NonNull); ok {
			t = val.OfType
		}
		name := t.Name()
		if val, ok := t.(*graphql.Enum); ok {
			t = val
		} else {
			name = name + suffix
		}
		typ, ok := my.types[name]
		if ok {
			whereFields[k] = &graphql.InputObjectFieldConfig{Type: typ}
		} else {
			fmt.Println("not found", name)
		}
	}
	whereInput := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: node.Name() + "WhereInput", Fields: whereFields,
	})

	whereFields["or"] = &graphql.InputObjectFieldConfig{Type: whereInput}
	whereFields["and"] = &graphql.InputObjectFieldConfig{Type: whereInput}
	whereFields["not"] = &graphql.InputObjectFieldConfig{Type: whereInput}

	queryFields := graphql.FieldConfigArgument{
		"id":         {Type: graphql.ID},
		"size":       {Type: graphql.Int},
		"page":       {Type: graphql.Int},
		"top":        {Type: graphql.Int},
		"last":       {Type: graphql.Int},
		"search":     {Type: graphql.String},
		"after":      {Type: Cursor},
		"before":     {Type: Cursor},
		"sort":       {Type: sortInput},
		"where":      {Type: whereInput},
		"distinctOn": {Type: graphql.NewList(graphql.String)},
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
