package core

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

type typeInfo struct {
	array    bool
	ptrType  reflect.Type
	implType reflect.Type
	baseType reflect.Type
}

type fieldParser func(field *reflect.StructField) (graphql.Type, error)

type (
	queryStruct        struct{}
	mutationStruct     struct{}
	subscriptionStruct struct{}
)

var (
	_queryType        = reflect.TypeOf((*queryStruct)(nil)).Elem()
	_mutationType     = reflect.TypeOf((*mutationStruct)(nil)).Elem()
	_subscriptionType = reflect.TypeOf((*subscriptionStruct)(nil)).Elem()
)

var (
	q = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query", Fields: graphql.Fields{},
	})
	m = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation", Fields: graphql.Fields{},
	})
	s = graphql.NewObject(graphql.ObjectConfig{
		Name: "Subscription", Fields: graphql.Fields{},
	})
)

var SortDirection = graphql.NewEnum(graphql.EnumConfig{
	Name:        "SortDirection",
	Description: "The direction of result ordering",
	Values: graphql.EnumValueConfigMap{
		"ASC": &graphql.EnumValueConfig{
			Value:       "ASC",
			Description: "Ascending order",
		},
		"DESC": &graphql.EnumValueConfig{
			Value:       "DESC",
			Description: "Descending order",
		},
		"ASC_NULLS_FIRST": &graphql.EnumValueConfig{
			Value:       "ASC_NULLS_FIRST",
			Description: "Ascending nulls first order",
		},
		"DESC_NULLS_FIRST": &graphql.EnumValueConfig{
			Value:       "DESC_NULLS_FIRST",
			Description: "Descending nulls first order",
		},
		"ASC_NULLS_LAST": &graphql.EnumValueConfig{
			Value:       "ASC_NULLS_LAST",
			Description: "Ascending nulls last order",
		},
		"DESC_NULLS_LAST": &graphql.EnumValueConfig{
			Value:       "DESC_NULLS_LAST",
			Description: "Descending nulls last order",
		},
	},
})
