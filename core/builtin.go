package core

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

func (my *Engine) asBuiltinScalar(typ reflect.Type) (graphql.Type, error) {
	base, err := unwrap(typ)
	if err != nil {
		return nil, err
	}

	var scalar graphql.Type
	if base.Kind() == reflect.Map {
		scalar = Json
	} else if base.PkgPath() == "" {
		switch base.Kind() {
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
		switch base.String() {
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

	return scalar, nil
}
