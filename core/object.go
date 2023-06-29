package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
)

type Object interface {
	GqlDescription() string
}

func parseFieldType(field *reflect.StructField, parsers []fieldParser, errString string) (graphql.Type, error) {
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

func (my *Engine) parseObject(info *typeInfo) (*graphql.Object, error) {
	if obj, ok := my.types[info.baseType]; ok {
		return obj.(*graphql.Object), nil
	}
	name, desc := info.baseType.Name(), ""

	prototype, ok := newPrototype(info.implType).(Object)
	if prototype != nil && ok {
		desc = prototype.GqlDescription()
	}

	object := graphql.NewObject(graphql.ObjectConfig{
		Name: name, Description: desc, Fields: graphql.Fields{},
	})
	my.types[info.baseType] = object
	err := my.unwrapObjectFields(info.baseType, object, 0)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (my *Engine) asObjectField(field *reflect.StructField) (graphql.Type, error) {
	info, err := unwrap(field.Type)
	if err != nil {
		return nil, err
	}
	typ, err := my.parseObject(&info)
	if err != nil {
		return nil, err
	}
	return wrapType(field.Type, typ), nil
}
