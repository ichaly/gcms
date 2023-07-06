package boot

import (
	"crypto/sha256"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"reflect"
)

type typeParser func(field *reflect.StructField) (graphql.Type, error)

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
