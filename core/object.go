package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"reflect"
)

type Object interface {
	GqlDescription() string
}

func (my *Engine) unwrapObjectFields(baseType reflect.Type, object *graphql.Object, depth int) error {
	size := baseType.NumField()
	for i := 0; i < size; i++ {
		f := baseType.Field(i)
		if !f.IsExported() {
			continue
		}
		if f.Anonymous {
			sub, err := unwrap(f.Type)
			if err != nil {
				return err
			}
			err = my.unwrapObjectFields(sub.baseType, object, depth+1)
			if err != nil {
				return err
			}
			continue
		}

		var fieldType graphql.Type
		fieldType, err := parseFieldType(&f, my.fieldParsers, "object field")
		if err != nil {
			return err
		}
		if fieldType == nil {
			panic(fmt.Errorf("unsupported field type: %s", f.Type.String()))
		}
		fieldName := strcase.ToLowerCamel(f.Name)
		object.AddFieldConfig(fieldName, &graphql.Field{
			Type: fieldType, Description: description(&f),
		})
	}
	return nil
}

func (my *Engine) buildObject(info *typeInfo) (*graphql.Object, error) {
	name, desc := info.baseType.Name(), ""
	if obj, ok := my.types[name]; ok {
		return obj.(*graphql.Object), nil
	}

	prototype, ok := newPrototype(info.implType).(Object)
	if prototype != nil && ok {
		desc = prototype.GqlDescription()
	}

	object := graphql.NewObject(graphql.ObjectConfig{
		Name: name, Description: desc, Fields: graphql.Fields{},
	})
	my.types[name] = object
	err := my.unwrapObjectFields(info.baseType, object, 0)
	if err != nil {
		return nil, err
	}

	my.buildSortInput(object)
	my.buildWhereInput(object)
	return object, nil
}

func (my *Engine) asObjectField(field *reflect.StructField) (graphql.Type, error) {
	info, err := unwrap(field.Type)
	if err != nil {
		return nil, err
	}
	typ, err := my.buildObject(&info)
	if err != nil {
		return nil, err
	}
	return wrapType(field.Type, typ), nil
}
