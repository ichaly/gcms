package core

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"reflect"
)

var _scalarType = reflect.TypeOf((*Scalar)(nil)).Elem()

type Scalar interface {
	Object
	GqlScalarSerialize() interface{}
	GqlScalarParseValue(value interface{})
}

func (my *Engine) buildCustomScalar(info *typeInfo) *graphql.Scalar {
	name := info.baseType.Name()
	if s, ok := my.types[name]; ok {
		return s.(*graphql.Scalar)
	}
	scalar := newPrototype(info.implType).(Scalar)

	literalParsing := func(valueAST ast.Value) interface{} {
		s := newPrototype(info.implType).(Scalar)
		s.GqlScalarParseValue(valueAST.GetValue())
		return s
	}

	d := graphql.NewScalar(graphql.ScalarConfig{
		Name:        name,
		Description: scalar.GqlDescription(),
		Serialize: func(value interface{}) interface{} {
			if s, ok := value.(Scalar); ok {
				return s.GqlScalarSerialize()
			}
			return nil
		},
		ParseValue: func(value interface{}) interface{} {
			s := newPrototype(info.implType).(Scalar)
			s.GqlScalarParseValue(value)
			return s
		},
		ParseLiteral: literalParsing,
	})
	my.types[name] = d
	return d
}

func (my *Engine) asCustomScalarField(field *reflect.StructField) (graphql.Type, error) {
	isScalar, info, err := implementsOf(field.Type, _scalarType)
	if err != nil {
		return nil, err
	}
	if !isScalar {
		return nil, nil
	}
	typ := my.buildCustomScalar(&info)
	return wrapType(field.Type, typ), nil
}
