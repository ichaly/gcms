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

func (my *Engine) parseEnum(info *typeInfo) *graphql.Enum {
	if d, ok := my.types[info.baseType]; ok {
		return d.(*graphql.Enum)
	}
	enum := newPrototype(info.implType).(Enum)

	name := info.baseType.Name()

	d := graphql.NewEnum(graphql.EnumConfig{
		Name:        name,
		Description: enum.GqlDescription(),
		Values:      enum.GqlEnumValues(),
	})
	my.types[info.baseType] = d
	return d
}

func (my *Engine) asEnumField(field *reflect.StructField) (graphql.Type, *typeInfo, error) {
	isEnum, info, err := implementsOf(field.Type, _enumType)
	if err != nil {
		return nil, &info, err
	}
	if !isEnum {
		return nil, &info, nil
	}
	typ := my.parseEnum(&info)
	return wrapType(field, typ, info.array), &info, nil
}
