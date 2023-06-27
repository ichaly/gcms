package core

import (
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/gcms/base"
	"gorm.io/gorm"
	"reflect"
)

type Engine struct {
	Query              *graphql.Object
	Mutation           *graphql.Object
	Subscription       *graphql.Object
	types              map[reflect.Type]graphql.Type
	objectFieldParsers []fieldParser
}

func NewEngine(eg base.EntityGroup, db *gorm.DB) (*Engine, error) {
	engine := &Engine{
		types: make(map[reflect.Type]graphql.Type),
	}
	engine.objectFieldParsers = []fieldParser{
		engine.asBuiltinScalar,
		engine.asIdField,
		engine.asEnumField,
		engine.asObjectField,
		engine.asCustomScalarField,
	}
	engine.Query = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Root Query",
		Fields:      graphql.Fields{},
	})
	for _, v := range eg.Entities {
		obj, err := engine.RegisterObject(v)
		if err != nil {
			return nil, err
		}
		engine.Query.AddFieldConfig(strcase.ToLowerCamel(obj.Name()), &graphql.Field{
			Type:        obj,
			Description: obj.Description(),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		})
	}
	return engine, nil
}
