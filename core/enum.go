package core

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

var _enumType = reflect.TypeOf((*Enum)(nil)).Elem()

type EnumValue struct {
	Value       interface{}
	Description string
}

type EnumValueMapping map[string]EnumValue

type Enum interface {
	GqlEnumDescription() string
	GqlEnumValues() EnumValueMapping
}

func (my *Engine) collectEnum(info *typeInfo) *graphql.Enum {
	if d, ok := my.types[info.baseType]; ok {
		return d.(*graphql.Enum)
	}
	enum := newPrototype(info.implType).(Enum)

	values := graphql.EnumValueConfigMap{}
	for valName, def := range enum.GqlEnumValues() {
		values[valName] = &graphql.EnumValueConfig{
			Value:       def.Value,
			Description: def.Description,
		}
	}

	name := info.baseType.Name()

	d := graphql.NewEnum(graphql.EnumConfig{
		Name:        name,
		Description: enum.GqlEnumDescription(),
		Values:      values,
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
	typ := my.collectEnum(&info)
	return wrapType(field, typ, info.array), &info, nil
}
