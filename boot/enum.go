package boot

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

var _enumType = reflect.TypeOf((*GqlEnum)(nil)).Elem()

type GqlEnum interface {
	GqlObject
	EnumValues() map[string]*struct {
		Value             interface{} `json:"value"`
		Description       string      `json:"description"`
		DeprecationReason string      `json:"deprecationReason"`
	}
}

func (my *Engine) asEnum(field *reflect.StructField) (graphql.Type, error) {
	isEnum := field.Type.Implements(_enumType)
	if !isEnum {
		return nil, nil
	}

	typ, err := my.buildEnum(field.Type)
	if err != nil {
		return nil, err
	}

	return wrapType(field.Type, typ), nil
}

func (my *Engine) buildEnum(base reflect.Type) (*graphql.Enum, error) {
	typ, err := unwrap(base)
	if err != nil {
		return nil, err
	}

	if val, ok := my.Types[typ.Name()]; ok {
		return val.(*graphql.Enum), nil
	}

	name, desc := typ.Name(), typ.(GqlEnum).Description()
	values := graphql.EnumValueConfigMap{}
	for k, v := range typ.(GqlEnum).EnumValues() {
		values[k] = &graphql.EnumValueConfig{
			Value:             v.Value,
			Description:       v.Description,
			DeprecationReason: v.DeprecationReason,
		}
	}

	e := graphql.NewEnum(graphql.EnumConfig{
		Name: name, Description: desc, Values: values,
	})

	my.Types[name] = e
	return e, nil
}
