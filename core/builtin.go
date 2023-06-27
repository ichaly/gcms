package core

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

func (my *Engine) asBuiltinScalar(field *reflect.StructField) (graphql.Type, *typeInfo, error) {
	info, err := unwrap(field.Type)
	if err != nil {
		return nil, &info, err
	}

	var scalar graphql.Type
	if info.baseType.PkgPath() == "" {
		// builtin
		switch info.baseType.Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Uint, reflect.Uint64, reflect.Uint32:
			scalar = graphql.Int
		case reflect.Int8, reflect.Int16, reflect.Uint8, reflect.Uint16:
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
		return nil, &info, nil
	}

	scalar = wrapType(field, scalar, info.array)
	return scalar, &info, nil
}
