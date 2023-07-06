package boot

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

var _idType = reflect.TypeOf((*GqlID)(nil)).Elem()

type GqlID interface {
	ID()
}

func (my *Engine) asId(field *reflect.StructField) (graphql.Type, error) {
	isId := field.Type.Implements(_idType)
	if !isId {
		return nil, nil
	}

	typ, err := unwrap(field.Type)
	if err != nil {
		return nil, err
	}

	switch typ.Kind() {
	case reflect.Uint64, reflect.Uint, reflect.Uint32,
		reflect.Int64, reflect.Int, reflect.Int32,
		reflect.String:
	default:
		panic(fmt.Errorf("%s cannot be used as an ID", typ.String()))
	}
	return wrapType(field.Type, graphql.ID), nil
}
