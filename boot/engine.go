package boot

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

type Engine struct {
	names map[string]reflect.Type
	types map[string]graphql.Type
}

func NewEngine() *Engine {
	return &Engine{
		names: map[string]reflect.Type{},
		types: map[string]graphql.Type{
			Query: q, Mutation: m, Subscription: s,
		},
	}
}

func (my *Engine) AddTo(
	resolver func(graphql.ResolveParams) (interface{}, error),
	objectName, fieldName, description string, tags ...string,
) error {
	if resolver == nil {
		return fmt.Errorf("missing resolve funtion")
	}
	if fieldName == "" {
		fieldName = getFuncName(resolver)
	}
	if fieldName == "" {
		return fmt.Errorf("missing field name")
	}
	val, ok := my.types[objectName]
	if !ok {
		return fmt.Errorf("missing object %s", objectName)
	}
	obj, ok := val.(*graphql.Object)
	if !ok {
		return fmt.Errorf("invalid object %s", objectName)
	}
	objectType := my.types[fieldName]
	obj.AddFieldConfig(fieldName, &graphql.Field{
		Name: fieldName, Type: objectType, Description: description,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return func() (interface{}, error) {
				return resolver(p)
			}, nil
		},
	})
	return nil
}

func (my *Engine) AddType(prototype interface{}) (*graphql.Object, error) {
	typ := reflect.TypeOf(prototype)
	obj, err := my.buildObject(typ)
	if err == nil {
		my.names[obj.Name()] = typ
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
