package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

var _idType = reflect.TypeOf((*GqlID)(nil)).Elem()

type GqlID interface {
	ID()
}

func (my *Engine) asId(typ reflect.Type) (graphql.Type, error) {
	isId := typ.Implements(_idType)
	if !isId {
		return nil, nil
	}

	base, err := unwrap(typ)
	if err != nil {
		return nil, err
	}

	switch base.Kind() {
	case reflect.Uint64, reflect.Uint, reflect.Uint32,
		reflect.Int64, reflect.Int, reflect.Int32,
		reflect.String:
	default:
		panic(fmt.Errorf("%s cannot be used as an ID", base.String()))
	}
	return graphql.ID, nil
}
