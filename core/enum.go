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

//type OrderDirection string
//
//const (
//	ASC              OrderDirection = "ASC"
//	DESC             OrderDirection = "DESC"
//	ASC_NULLS_FIRST  OrderDirection = "ASC_NULLS_FIRST"
//	DESC_NULLS_FIRST OrderDirection = "DESC_NULLS_FIRST"
//	ASC_NULLS_LAST   OrderDirection = "ASC_NULLS_LAST"
//	DESC_NULLS_LAST  OrderDirection = "DESC_NULLS_LAST"
//)
//
//func (OrderDirection) GqlDescription() string {
//	return "The direction of result ordering"
//}
//
//func (OrderDirection) GqlEnumValues() map[string]*graphql.EnumValueConfig {
//	return map[string]*graphql.EnumValueConfig{
//		"ASC":             {Value: ASC, Description: "Ascending order"},
//		"DESC":            {Value: DESC, Description: "Descending order"},
//		"ASC_NULL_FIRST":  {Value: ASC_NULLS_FIRST, Description: "Ascending nulls first order"},
//		"DESC_NULL_FIRST": {Value: DESC_NULLS_FIRST, Description: "Descending nulls first order"},
//		"ASC_NULL_LAST":   {Value: ASC_NULLS_LAST, Description: "Ascending nulls last order"},
//		"DESC_NULL_LAST":  {Value: DESC_NULLS_LAST, Description: "Descending nulls last order"},
//	}
//}

var OrderDirection = graphql.NewEnum(graphql.EnumConfig{
	Name:        "OrderDirection",
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

func (my *Engine) asEnumField(field *reflect.StructField) (graphql.Type, error) {
	isEnum, info, err := implementsOf(field.Type, _enumType)
	if err != nil {
		return nil, err
	}
	if !isEnum {
		return nil, nil
	}
	typ := my.parseEnum(&info)
	return wrapType(field.Type, typ), nil
}
