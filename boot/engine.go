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

func (my *Engine) Register(node interface{}) error {
	if node == nil {
		return fmt.Errorf("node can't be nil")
	}
	value := reflect.ValueOf(node)

	hostFunc := value.MethodByName("Host")
	if !hostFunc.IsValid() {
		return fmt.Errorf("missing hostFunc")
	}
	resolveFunc := value.MethodByName("Resolve")
	if !resolveFunc.IsValid() {
		return fmt.Errorf("missing resolve funtion")
	}

	out := resolveFunc.Type().Out(0)
	if out.Kind() == reflect.Func {
		out = out.Out(0)
	}
	outType, err := parseType(out, "result",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	if err != nil {
		return err
	}

	obj := reflect.TypeOf(hostFunc.Call(make([]reflect.Value, 0))[0].Interface())
	objType, err := parseType(obj, "host",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	if err != nil {
		return err
	}

	var name, desc string
	if nameFunc := value.MethodByName("Name"); nameFunc.IsValid() {
		name = nameFunc.Call(make([]reflect.Value, 0))[0].Interface().(string)
	}
	if descFunc := value.MethodByName("Description"); descFunc.IsValid() {
		desc = descFunc.Call(make([]reflect.Value, 0))[0].Interface().(string)
	}

	if name == "" {
		return fmt.Errorf("missing field name")
	}
	host, ok := objType.(*graphql.Object)
	if !ok {
		return fmt.Errorf("invalid host %s", obj)
	}

	var args graphql.FieldConfigArgument
	if _, ok := outType.(*graphql.Object); ok {
		args = graphql.FieldConfigArgument{
			"size":  {Type: graphql.Int},
			"page":  {Type: graphql.Int},
			"sort":  {Type: my.types[outType.Name()+"SortInput"]},
			"where": {Type: my.types[outType.Name()+"WhereInput"]},
		}
	}
	host.AddFieldConfig(name, &graphql.Field{
		Type: wrapType(out, outType), Args: args, Description: desc,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			r := resolveFunc.Call([]reflect.Value{reflect.ValueOf(p)})
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

func (my *Engine) Schema() (graphql.Schema, error) {
	config := graphql.SchemaConfig{}
	if q := my.checkObject("query"); q != nil {
		config.Query = q
	}
	if m := my.checkObject("mutation"); m != nil {
		config.Mutation = m
	}
	if s := my.checkObject("subscription"); s != nil {
		config.Subscription = s
	}
	return graphql.NewSchema(config)
}

func (my *Engine) checkObject(name string) *graphql.Object {
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
