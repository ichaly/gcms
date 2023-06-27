package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"reflect"
)

var (
	_objectType = reflect.TypeOf((*Object)(nil)).Elem()
)

type Object interface {
	GqlObjectDescription() string
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
		fieldType, _, err := parseField(&f, my.objectFieldParsers, "object field")
		if err != nil {
			return err
		}
		if fieldType == nil {
			panic(fmt.Errorf("unsupported field type: %s", f.Type.String()))
		}
		object.AddFieldConfig(strcase.ToLowerCamel(f.Name), &graphql.Field{
			Type:        fieldType,
			Description: description(&f),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		})
	}
	return nil
}

func (my *Engine) collectObject(info *typeInfo) (*graphql.Object, error) {
	if obj, ok := my.types[info.baseType]; ok {
		return obj.(*graphql.Object), nil
	}
	prototype, ok := newPrototype(info.implType).(Object)
	name, desc := info.baseType.Name(), ""
	if prototype != nil && ok {
		desc = prototype.GqlObjectDescription()
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

func (my *Engine) asObjectField(field *reflect.StructField) (graphql.Type, *typeInfo, error) {
	isObj, info, err := implementsOf(field.Type, _objectType)
	if err != nil {
		return nil, &info, err
	}
	if !isObj {
		info, err = unwrap(field.Type)
		if err != nil {
			return nil, &info, err
		}
	}
	typ, err := my.collectObject(&info)
	if err != nil {
		return nil, nil, err
	}
	return wrapType(field, typ, info.array), &info, nil
}

func (my *Engine) registerObject(p reflect.Type) (*graphql.Object, error) {
	isObj, info, err := implementsOf(p, _objectType)
	if err != nil {
		return nil, err
	}
	if !isObj {
		info, err = unwrap(p)
		if err != nil {
			return nil, err
		}
	}
	return my.collectObject(&info)
}

func (my *Engine) RegisterObject(prototype interface{}) (*graphql.Object, error) {
	return my.registerObject(reflect.TypeOf(prototype))
}
