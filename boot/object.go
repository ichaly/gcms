package boot

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"reflect"
)

var _objectType = reflect.TypeOf((*GqlObject)(nil)).Elem()

type GqlObject interface {
	Description() string
}

func (my *Engine) asObject(field *reflect.StructField) (graphql.Type, error) {
	typ, err := my.buildObject(field.Type)
	if err != nil {
		return nil, err
	}
	return wrapType(field.Type, typ), nil
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

	desc, name := "", typ.Name()
	if ptr, ok := newPrototype(typ).(GqlObject); ok {
		desc = ptr.Description()
	}
	o := graphql.NewObject(graphql.ObjectConfig{
		Name: name, Description: desc, Fields: graphql.Fields{},
	})
	err = my.parseFields(typ, o, 0)
	if err != nil {
		return nil, err
	}
	my.types[name] = o

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

		fieldType, err := parseType(
			&f, "obj field",
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
			Type: fieldType,
			//Description: description(&f),
		})
	}
	return nil
}

func parseType(field *reflect.StructField, errString string, parsers ...typeParser) (graphql.Type, error) {
	for _, check := range parsers {
		typ, err := check(field)
		if err != nil {
			return nil, err
		}
		if typ == nil {
			continue
		}
		return typ, nil
	}
	return nil, fmt.Errorf("unsupported type('%s') for %s '%s'", field.Type.String(), errString, field.Name)
}
