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

func (my *Engine) parseCustomScalar(info *typeInfo) *graphql.Scalar {
	if s, ok := my.types[info.baseType]; ok {
		return s.(*graphql.Scalar)
	}

	scalar := newPrototype(info.implType).(Scalar)

	name := info.baseType.Name()

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
	my.types[info.baseType] = d
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
	typ := my.parseCustomScalar(&info)
	return wrapType(field.Type, typ), nil
}
