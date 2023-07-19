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
	return &Engine{types: map[string]graphql.Type{}}
}

func (my *Engine) Schema() (graphql.Schema, error) {
	config := graphql.SchemaConfig{}
	if q := my.verifyObject("query"); q != nil {
		config.Query = q
	}
	if m := my.verifyObject("mutation"); m != nil {
		config.Mutation = m
	}
	if s := my.verifyObject("subscription"); s != nil {
		config.Subscription = s
	}
	return graphql.NewSchema(config)
}

func (my *Engine) Register(node interface{}) error {
	if node == nil {
		return fmt.Errorf("node can't be nil")
	}

	val := reflect.ValueOf(node)
	reflectType := reflect.TypeOf(node)

	base, err := unwrap(reflectType)
	if err != nil {
		return err
	}
	var name, desc = base.Name(), ""
	if _, o := reflectType.MethodByName("Host"); !o {
		return fmt.Errorf("missing host")
	}
	if _, o := reflectType.MethodByName("Resolve"); !o {
		return fmt.Errorf("missing resolve funtion")
	}

	if _, o := reflectType.MethodByName("Name"); o {
		name = val.MethodByName("Name").Call(make([]reflect.Value, 0))[0].Interface().(string)
	}
	if _, o := reflectType.MethodByName("Description"); o {
		desc = val.MethodByName("Description").Call(make([]reflect.Value, 0))[0].Interface().(string)
	}

	if name == "" {
		return fmt.Errorf("missing field name")
	}

	host := val.MethodByName("Host").Call(make([]reflect.Value, 0))[0].Interface()
	graphqlType, err := parseType(reflect.TypeOf(host), "host",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	obj, ok := graphqlType.(*graphql.Object)
	if !ok {
		return fmt.Errorf("invalid object %s", host)
	}

	resolver := val.MethodByName("Resolve")
	out := resolver.Type().Out(0)
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
			r := resolver.Call([]reflect.Value{reflect.ValueOf(p)})
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

func (my *Engine) verifyObject(name string) *graphql.Object {
	val, ok := my.types[name]
	if !ok {
		return nil
	}
	obj, ok := val.(*graphql.Object)
	if !ok {
		return nil
	}
	if len(obj.Fields()) <= 0 {
		return nil
	}
	return obj
}
