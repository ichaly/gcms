package core

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

var _enumType = reflect.TypeOf((*Enum)(nil)).Elem()

type Enum interface {
	Object
	GqlEnumValues() map[string]*graphql.EnumValueConfig
}

func (my *Engine) buildEnum(info *typeInfo) *graphql.Enum {
	name := info.baseType.Name()
	if d, ok := my.types[name]; ok {
		return d.(*graphql.Enum)
	}
	enum := newPrototype(info.implType).(Enum)
	d := graphql.NewEnum(graphql.EnumConfig{
		Name:        name,
		Description: enum.GqlDescription(),
		Values:      enum.GqlEnumValues(),
	})
	my.types[name] = d
	return d
}

func (my *Engine) asEnumField(field *reflect.StructField) (graphql.Type, error) {
	isEnum, info, err := implementsOf(field.Type, _enumType)
	if err != nil {
		return nil, err
	}
	if !isEnum {
		return nil, nil
	}
	typ := my.buildEnum(&info)
	return wrapType(field.Type, typ), nil
}
