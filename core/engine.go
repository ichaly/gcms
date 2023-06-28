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
	my := &Engine{
		types: make(map[reflect.Type]graphql.Type),
	}
	my.objectFieldParsers = []fieldParser{
		my.asBuiltinScalar,
		my.asIdField,
		my.asEnumField,
		my.asObjectField,
		my.asCustomScalarField,
	}
	my.Query = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query", Description: "Root Query", Fields: graphql.Fields{},
	})
	for _, v := range eg.Entities {
		obj, err := my.RegisterObject(v)
		if err != nil {
			return nil, err
		}
		name := strcase.ToLowerCamel(obj.Name())
		my.Query.AddFieldConfig(name, &graphql.Field{
			Type:        obj,
			Args:        map[string]*graphql.ArgumentConfig{},
			Description: obj.Description(),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		})
	}
	return my, nil
}
