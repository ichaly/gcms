package core

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

func (my *Engine) Register(node Schema) error {
	if node == nil {
		return fmt.Errorf("node can't be nil")
	}

	out := reflect.TypeOf(node.Type())
	outType, err := parseType(out, "result",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	if err != nil {
		return err
	}
	obj := reflect.TypeOf(node.Host())
	objType, err := parseType(obj, "host",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	if err != nil {
		return err
	}

	host, ok := objType.(*graphql.Object)
	if !ok {
		return fmt.Errorf("invalid host %s", node.Host())
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
	if host.Name() == "mutation" {
		args["data"] = &graphql.ArgumentConfig{Type: my.types[outType.Name()+"DataInput"]}
		args["delete"] = &graphql.ArgumentConfig{Type: graphql.Boolean}
	}
	host.AddFieldConfig(node.Name(), &graphql.Field{
		Type: wrapType(out, outType), Args: args, Resolve: node.Resolve, Description: node.Description(),
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
