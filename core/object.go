package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
)

var _objectType = reflect.TypeOf((*GqlObject)(nil)).Elem()

type GqlObject interface {
	Description() string
}

func (my *Engine) asObject(typ reflect.Type) (graphql.Type, error) {
	object, err := my.buildObject(typ)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (my *Engine) buildObject(base reflect.Type) (*graphql.Object, error) {
	typ, err := unwrap(base)
	if err != nil {
		return nil, err
	}

	obj, ok := my.types[typ.Name()]
	if ok {
		return obj.(*graphql.Object), nil
	}

	name, desc := typ.Name(), ""
	if isObject := typ.Implements(_objectType); isObject {
		desc = newPrototype(typ).(GqlObject).Description()
	}
	o := graphql.NewObject(graphql.ObjectConfig{
		Name: name, Description: desc, Fields: graphql.Fields{},
	})
	err = my.parseFields(typ, o, 0)
	if err != nil {
		return nil, err
	}
	my.types[name] = o

	my.buildDataInput(o)
	my.buildSortInput(o)
	my.buildWhereInput(o)
	return o, nil
}

func (my *Engine) parseFields(typ reflect.Type, obj *graphql.Object, dep int) error {
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if !f.IsExported() {
			continue
		}
		if f.Anonymous {
			sub, err := unwrap(f.Type)
			if err != nil {
				return err
			}
			err = my.parseFields(sub, obj, dep+1)
			if err != nil {
				return err
			}
			continue
		}

		fieldType, err := parseType(f.Type, "obj field",
			my.asBuiltinScalar,
			my.asCustomScalar,
			my.asId,
			my.asEnum,
			my.asObject,
		)
		if err != nil {
			return err
		}
		if fieldType == nil {
			panic(fmt.Errorf("unsupported field type: %s", f.Type.String()))
		}
		fieldName := strcase.ToLowerCamel(f.Name)
		obj.AddFieldConfig(fieldName, &graphql.Field{
			Type:        wrapType(f.Type, fieldType),
			Resolve:     defaultResolver,
			Description: description(&f),
		})
	}
	return nil
}

func description(field *reflect.StructField) string {
	tag := field.Tag.Get("gorm")
	tags := strings.Split(tag, ";")
	for _, t := range tags {
		if strings.HasPrefix(t, "comment:") {
			return t[8:]
		}
	}
	return ""
}
