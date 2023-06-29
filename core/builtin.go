package core

import (
	"crypto/sha256"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"reflect"
)

var encPrefix = "__gc:foobar:"
var decPrefix = "__gc:enc:"
var key = sha256.Sum256([]byte("123"))

var Void = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "Void",
	Description:  "void",
	Serialize:    func(value interface{}) interface{} { return "0" },
	ParseValue:   func(value interface{}) interface{} { return 0 },
	ParseLiteral: func(valueAST ast.Value) interface{} { return 0 },
})

var Cursor = graphql.NewScalar(graphql.ScalarConfig{
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

func (my *Engine) asBuiltinScalarField(field *reflect.StructField) (graphql.Type, error) {
	info, err := unwrap(field.Type)
	if err != nil {
		return nil, err
	}

	var scalar graphql.Type
	if info.baseType.PkgPath() == "" {
		// builtin
		switch info.baseType.Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Uint, reflect.Uint64, reflect.Uint32,
			reflect.Int8, reflect.Int16, reflect.Uint8, reflect.Uint16:
			scalar = graphql.Int
		case reflect.Float32, reflect.Float64:
			scalar = graphql.Float
		case reflect.Bool:
			scalar = graphql.Boolean
		case reflect.String:
			scalar = graphql.String
		default:
		}
	} else {
		switch info.baseType.String() {
		case "time.Time":
			scalar = graphql.DateTime
		}
	}

	if scalar == nil {
		return nil, nil
	}

	return wrapType(field.Type, scalar), nil
}
