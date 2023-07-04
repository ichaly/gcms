package core

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

var _enumType = reflect.TypeOf((*Enum)(nil)).Elem()

var SortDirection = graphql.NewEnum(graphql.EnumConfig{
	Name:        "SortDirection",
	Description: "The direction of result ordering",
	Values: graphql.EnumValueConfigMap{
		"ASC": &graphql.EnumValueConfig{
			Value:       "ASC",
			Description: "Ascending order",
		},
		"DESC": &graphql.EnumValueConfig{
			Value:       "DESC",
			Description: "Descending order",
		},
		"ASC_NULLS_FIRST": &graphql.EnumValueConfig{
			Value:       "ASC_NULLS_FIRST",
			Description: "Ascending nulls first order",
		},
		"DESC_NULLS_FIRST": &graphql.EnumValueConfig{
			Value:       "DESC_NULLS_FIRST",
			Description: "Descending nulls first order",
		},
		"ASC_NULLS_LAST": &graphql.EnumValueConfig{
			Value:       "ASC_NULLS_LAST",
			Description: "Ascending nulls last order",
		},
		"DESC_NULLS_LAST": &graphql.EnumValueConfig{
			Value:       "DESC_NULLS_LAST",
			Description: "Descending nulls last order",
		},
	},
})

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
