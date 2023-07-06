package boot

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

func (my *Engine) asBuiltinScalar(field *reflect.StructField) (graphql.Type, error) {
	typ, err := unwrap(field.Type)
	if err != nil {
		return nil, err
	}

	var scalar graphql.Type

	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Uint, reflect.Uint64, reflect.Uint32,
			reflect.Int8, reflect.Int16, reflect.Uint8, reflect.Uint16:
			scalar = graphql.Int
		case reflect.Float32, reflect.Float64:
			scalar = graphql.Float
		case reflect.Bool:
			scalar = graphql.Boolean
		case reflect.String:
			scalar = graphql.String
		}
	} else {
		switch typ.String() {
		case "time.Time", "gorm.DeletedAt":
			scalar = graphql.DateTime
		case "boot.Void":
			scalar = Void
		case "boot.Cursor":
			scalar = Cursor
		}
	}

	if scalar == nil {
		return nil, nil
	}

	return wrapType(field.Type, scalar), nil
}
