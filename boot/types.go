package boot

import (
	"crypto/sha256"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"reflect"
)

type typeParser func(typ reflect.Type) (graphql.Type, error)

type __input struct {
	Name        string
	Type        graphql.Type
	Description string
}

const (
	Query        = "Query"
	Mutation     = "Mutation"
	Subscription = "Subscription"
)

var (
	encPrefix = "__gc:foobar:"
	decPrefix = "__gc:enc:"
	key       = sha256.Sum256([]byte("123"))
)

var (
	Void = graphql.NewScalar(graphql.ScalarConfig{
		Name:         "Void",
		Description:  "void",
		Serialize:    func(value interface{}) interface{} { return "0" },
		ParseValue:   func(value interface{}) interface{} { return 0 },
		ParseLiteral: func(valueAST ast.Value) interface{} { return 0 },
	})

	Cursor = graphql.NewScalar(graphql.ScalarConfig{
		Name:        "Cursor",
		Description: "A `Cursor` is an encoded string use for pagination",
		Serialize: func(val interface{}) interface{} {
			//js := []byte(fmt.Sprintf(`{ me: "null", posts_cursor: "%v12345" }`, encPrefix))
			//nonce := sha256.Sum256(js)
			//out, err := encryptValues(js, []byte(encPrefix), []byte(decPrefix), nonce[:], key)
			//if err != nil {
			//	return nil
			//}
			return val
		},
		ParseValue: func(val interface{}) interface{} {
			return val
		},
		ParseLiteral: func(val ast.Value) interface{} {
			return nil
		},
	})
)

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
