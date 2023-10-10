package core

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"reflect"
)

type Schema interface {
	Name() string
	Host() interface{}
	Type() interface{}
	Description() string
	Resolve(p graphql.ResolveParams) (interface{}, error)
}

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
		{Name: "regex", Desc: "Value matches regex pattern"},
		{Name: "notRegex", Desc: "Value not matching regex pattern"},
		{Name: "like", Desc: "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
		{Name: "notLike", Desc: "Value not matching pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"},
		//{Name: "iLike", Desc: "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
		//{Name: "notILike", Desc: "Value not matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"},
		//{Name: "iRegex", Desc: "Value matches (case-insensitive) regex pattern"},
		//{Name: "notIRegex", Desc: "Value not matching (case-insensitive) regex pattern"},
		//{Name: "similar", Desc: "Value matching regex pattern. Similar to the 'like' operator but with support for regex. Pattern must match entire value."},
		//{Name: "notSimilar", Desc: "Value not matching regex pattern. Similar to the 'like' operator but with support for regex. Pattern must not match entire value."},
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

	Json = graphql.NewScalar(
		graphql.ScalarConfig{
			Name:        "Json",
			Description: "The `Json` scalar type represents Json values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf)",
			Serialize: func(value interface{}) interface{} {
				return value
			},
			ParseValue: func(value interface{}) interface{} {
				return value
			},
			ParseLiteral: parseLiteral,
		},
	)
)

func parseLiteral(astValue ast.Value) interface{} {
	kind := astValue.GetKind()
	switch kind {
	case kinds.StringValue:
		return astValue.GetValue()
	case kinds.BooleanValue:
		return astValue.GetValue()
	case kinds.IntValue:
		return astValue.GetValue()
	case kinds.FloatValue:
		return astValue.GetValue()
	case kinds.ObjectValue:
		obj := make(map[string]interface{})
		for _, v := range astValue.GetValue().([]*ast.ObjectField) {
			obj[v.Name.Value] = parseLiteral(v.Value)
		}
		return obj
	case kinds.ListValue:
		list := make([]interface{}, 0)
		for _, v := range astValue.GetValue().([]ast.Value) {
			list = append(list, parseLiteral(v))
		}
		return list
	default:
		return nil
	}
}

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
