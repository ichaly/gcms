package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

func (my *Engine) RegisterObject(prototype interface{}) (*graphql.Object, error) {
	return my.registerObject(reflect.TypeOf(prototype))
}

func (my *Engine) registerObject(p reflect.Type) (*graphql.Object, error) {
	info, err := unwrap(p)
	if err != nil {
		return nil, err
	}
	if obj, ok := my.types[info.baseType]; ok {
		return obj, nil
	}
	object := graphql.NewObject(graphql.ObjectConfig{
		Name:   info.baseType.Name(),
		Fields: graphql.Fields{},
	})
	my.types[info.baseType] = object

	err = my.unwrapObjectFields(info.baseType, object, 0)
	if err != nil {
		return nil, err
	}
	return object, nil
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
		fieldType, _, err := asBuiltinScalar(&f)
		if err != nil {
			return err
		}
		if fieldType == nil {
			panic(fmt.Errorf("unsupported field type: %s", f.Type.String()))
		}
		object.AddFieldConfig(f.Name, &graphql.Field{
			Type:        fieldType,
			Description: description(&f),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		})
	}
	return nil
}
