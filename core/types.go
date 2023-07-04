package core

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

const (
	Query        = "Query"
	Mutation     = "Mutation"
	Subscription = "Subscription"
)

type typeInfo struct {
	array    bool
	ptrType  reflect.Type
	implType reflect.Type
	baseType reflect.Type
}

type fieldParser func(field *reflect.StructField) (graphql.Type, error)

var (
	q = graphql.NewObject(graphql.ObjectConfig{
		Name: Query, Fields: graphql.Fields{},
	})
	m = graphql.NewObject(graphql.ObjectConfig{
		Name: Mutation, Fields: graphql.Fields{},
	})
	s = graphql.NewObject(graphql.ObjectConfig{
		Name: Subscription, Fields: graphql.Fields{},
	})
)

type __input struct {
	Name        string
	Type        graphql.Type
	Description string
}

var expNull = []__input{
	{Name: "isNull", Type: graphql.Boolean, Description: "Is value null (true) or not null (false)"},
}

var expList = []__input{
	{Name: "in", Description: "Is in list of values"},
	{Name: "notIn", Description: "Is not in list of values"},
}

var expScalar = []__input{
	{Name: "eq", Description: "Equals value"},
	{Name: "ne", Description: "Does not equal value"},
	{Name: "gt", Description: "Is greater than value"},
	{Name: "lt", Description: "Is lesser than value"},
	{Name: "ge", Description: "Is greater than or equals value"},
	{Name: "le", Description: "Is lesser than or equals value"},
	{Name: "like", Description: "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
	{Name: "notLike", Description: "Value not matching pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"},
	{Name: "iLike", Description: "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
	{Name: "notILike", Description: "Value not matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"},
	{Name: "similar", Description: "Value matching regex pattern. Similar to the 'like' operator but with support for regex. Pattern must match entire value."},
	{Name: "notSimilar", Description: "Value not matching regex pattern. Similar to the 'like' operator but with support for regex. Pattern must not match entire value."},
	{Name: "regex", Description: "Value matches regex pattern"},
	{Name: "notRegex", Description: "Value not matching regex pattern"},
	{Name: "iRegex", Description: "Value matches (case-insensitive) regex pattern"},
	{Name: "notIRegex", Description: "Value not matching (case-insensitive) regex pattern"},
}
