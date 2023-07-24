package base

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"reflect"
)

type typeParser func(typ reflect.Type) (graphql.Type, error)

type (
	query        struct{}
	mutation     struct{}
	subscription struct{}
	input        struct {
		Name string
		Desc string
		Type graphql.Type
	}
)

var (
	Query        = &query{}
	Mutation     = &mutation{}
	Subscription = &subscription{}
)

var (
	expNull = []input{
		{Name: "isNull", Type: graphql.Boolean, Desc: "Is value null (true) or not null (false)"},
	}
	expList = []input{
		{Name: "in", Desc: "Is in list of values"},
		{Name: "notIn", Desc: "Is not in list of values"},
	}
	expBase = []input{
		{Name: "eq", Desc: "Equals value"},
		{Name: "ne", Desc: "Does not equal value"},
		{Name: "gt", Desc: "Is greater than value"},
		{Name: "lt", Desc: "Is lesser than value"},
		{Name: "ge", Desc: "Is greater than or equals value"},
		{Name: "le", Desc: "Is lesser than or equals value"},
		{Name: "like", Desc: "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
		{Name: "notLike", Desc: "Value not matching pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"},
		{Name: "iLike", Desc: "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
		{Name: "notILike", Desc: "Value not matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"},
		{Name: "similar", Desc: "Value matching regex pattern. Similar to the 'like' operator but with support for regex. Pattern must match entire value."},
		{Name: "notSimilar", Desc: "Value not matching regex pattern. Similar to the 'like' operator but with support for regex. Pattern must not match entire value."},
		{Name: "regex", Desc: "Value matches regex pattern"},
		{Name: "notRegex", Desc: "Value not matching regex pattern"},
		{Name: "iRegex", Desc: "Value matches (case-insensitive) regex pattern"},
		{Name: "notIRegex", Desc: "Value not matching (case-insensitive) regex pattern"},
	}
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
