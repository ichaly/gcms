package boot

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

type Engine struct {
	types map[string]graphql.Type
}

func NewEngine() *Engine {
	return &Engine{
		types: map[string]graphql.Type{
			Query: q, Mutation: m, Subscription: s,
		},
	}
}

func (my *Engine) AddTo(
	resolver interface{}, objectName, fieldName, description string,
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
	typ := reflect.TypeOf(resolver)
	if typ.Kind() != reflect.Func {
		return fmt.Errorf("resolve prototype should be a function")
	}
	if typ.NumOut() != 2 {
		return fmt.Errorf("resolve prototype should return 2 values")
	}
	out := typ.Out(0)
	src, err := my.buildObject(out)
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
		"sort":       {Type: my.types[src.Name()+"SortInput"]},
		"where":      {Type: my.types[src.Name()+"WhereInput"]},
	}
	obj.AddFieldConfig(fieldName, &graphql.Field{
		Name: fieldName, Type: wrapType(out, src), Args: queryArgs, Description: description,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return func() (interface{}, error) {
				f := reflect.ValueOf(resolver)
				v := []reflect.Value{reflect.ValueOf(p)}
				r := f.Call(v)
				res := r[0].Interface()
				err, _ := r[1].Interface().(error)
				return res, err
			}, nil
		},
	})
	return nil
}

func (my *Engine) AddType(prototype interface{}) (*graphql.Object, error) {
	typ := reflect.TypeOf(prototype)
	return my.buildObject(typ)
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
