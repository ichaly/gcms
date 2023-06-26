package core

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"reflect"
)

type Engine struct {
	Query        *graphql.Object
	Mutation     *graphql.Object
	Subscription *graphql.Object
	types        map[reflect.Type]*graphql.Object
}

func NewEngine(eg base.EntityGroup) (*Engine, error) {
	e := &Engine{
		types: make(map[reflect.Type]*graphql.Object),
	}
	e.Query = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Root Query",
		Fields:      graphql.Fields{},
	})
	for _, v := range eg.Entities {
		obj, err := e.RegisterObject(v)
		if err != nil {
			return nil, err
		}
		e.Query.AddFieldConfig(obj.Name(), &graphql.Field{
			Type:        obj,
			Name:        obj.Name(),
			Description: obj.Description(),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		})
	}
	return e, nil
}
