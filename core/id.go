package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

var _idType = reflect.TypeOf((*ID)(nil)).Elem()

type ID interface {
	GqlID()
}

func (my *Engine) asIdField(field *reflect.StructField) (graphql.Type, error) {
	isId, info, err := implementsOf(field.Type, _idType)
	if err != nil {
		return nil, err
	}
	if !isId {
		return nil, nil
	}
	switch info.baseType.Kind() {
	case reflect.Uint64, reflect.Uint, reflect.Uint32,
		reflect.Int64, reflect.Int, reflect.Int32,
		reflect.String:
	default:
		panic(fmt.Errorf("%s cannot be used as an ID", info.baseType.String()))
	}
	return wrapType(field.Type, graphql.ID), nil
}
