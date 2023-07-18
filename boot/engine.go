package boot

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

type Engine struct {
	types    map[string]graphql.Type
	builders []chainBuilder
}

func NewEngine() *Engine {
	return &Engine{types: map[string]graphql.Type{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: Query, Fields: graphql.Fields{},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: Mutation, Fields: graphql.Fields{},
		}),
		Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name: Subscription, Fields: graphql.Fields{},
		}),
	}}
}

func (my *Engine) Schema() (graphql.Schema, error) {
	for _, b := range my.builders {
		if err := b.build(my); err != nil {
			panic(err)
		}
	}
	config := graphql.SchemaConfig{}
	if q := my.types[Query].(*graphql.Object); len(q.Fields()) > 0 {
		config.Query = q
	}
	if m := my.types[Mutation].(*graphql.Object); len(m.Fields()) > 0 {
		config.Mutation = m
	}
	if s := my.types[Subscription].(*graphql.Object); len(s.Fields()) > 0 {
		config.Subscription = s
	}
	return graphql.NewSchema(config)
}

func (my *Engine) Register(resolver interface{}, host, name, desc string) error {
	if resolver == nil {
		return fmt.Errorf("missing resolve funtion")
	}
	if name == "" {
		name = getFuncName(resolver)
	}
	if name == "" {
		return fmt.Errorf("missing field name")
	}
	val, ok := my.types[host]
	if !ok {
		return fmt.Errorf("missing object %s", host)
	}
	obj, ok := val.(*graphql.Object)
	if !ok {
		return fmt.Errorf("invalid object %s", host)
	}
	typ := reflect.TypeOf(resolver)
	if typ.Kind() != reflect.Func {
		return fmt.Errorf("resolve prototype should be a function")
	}
	if typ.NumOut() == 0 {
		return fmt.Errorf("resolve prototype should return 1 or 2 values")
	}
	out := typ.Out(0)
	if out.Kind() == reflect.Func {
		out = out.Out(0)
	}
	src, err := parseType(out, "result",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	if err != nil {
		return err
	}
	var queryArgs graphql.FieldConfigArgument
	if _, ok := src.(*graphql.Object); ok {
		queryArgs = graphql.FieldConfigArgument{
			"size":  {Type: graphql.Int},
			"page":  {Type: graphql.Int},
			"sort":  {Type: my.types[src.Name()+"SortInput"]},
			"where": {Type: my.types[src.Name()+"WhereInput"]},
		}
		if desc == "" {
			desc = src.Description()
		}
	}
	obj.AddFieldConfig(name, &graphql.Field{
		Name: name, Type: wrapType(out, src), Args: queryArgs, Description: desc,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			f := reflect.ValueOf(resolver)
			v := []reflect.Value{reflect.ValueOf(p)}
			r := f.Call(v)
			if r[0].Type().Kind() == reflect.Func {
				return func() (interface{}, error) {
					return r[0].Call(nil)[0].Interface(), nil
				}, nil
			} else {
				res := r[0].Interface()
				err, _ := r[1].Interface().(error)
				return res, err
			}
		},
	})
	return nil
}
