package base

import (
	"github.com/graphql-go/graphql"
	"reflect"
	"strings"
)

// defaultResolver uses reflection to attempt to resolve the result of a given field.
func defaultResolver(p graphql.ResolveParams) (interface{}, error) {
	source := p.Source
	fieldName := p.Info.FieldName
	sourceVal := reflect.ValueOf(source)
	if sourceVal.IsValid() && sourceVal.Type().Kind() == reflect.Ptr {
		sourceVal = sourceVal.Elem()
	}
	if !sourceVal.IsValid() {
		return nil, nil
	}

	// Struct
	if sourceVal.Type().Kind() == reflect.Struct {
		_, val, err := findFieldInStruct(sourceVal, fieldName)
		return val, err
	}

	// map[string]interface
	if sourceMap, ok := source.(map[string]interface{}); ok {
		property := sourceMap[fieldName]
		val := reflect.ValueOf(property)
		if val.IsValid() && val.Type().Kind() == reflect.Func {
			// try type casting the func to the most basic func signature
			// for more complex signatures, user have to define ResolveFn
			if propertyFn, ok := property.(func() interface{}); ok {
				return propertyFn(), nil
			}
		}
		return property, nil
	}

	// last resort, return nil
	return nil, nil
}

func findFieldInStruct(source reflect.Value, fieldName string) (bool, interface{}, error) {
	for i := 0; i < source.NumField(); i++ {
		fieldValue := source.Field(i)
		fieldType := source.Type().Field(i)

		if strings.EqualFold(fieldType.Name, fieldName) {
			//if fieldType.Name == strings.Title(fieldName) {
			// If ptr and value is nil return nil
			if fieldValue.Type().Kind() == reflect.Ptr && fieldValue.IsNil() {
				return true, nil, nil
			}
			return true, fieldValue.Interface(), nil
		}

		tag := fieldType.Tag
		checkTag := func(tagName string) bool {
			t := tag.Get(tagName)
			tOptions := strings.Split(t, ",")
			if len(tOptions) == 0 {
				return false
			}
			if tOptions[0] != fieldName {
				return false
			}
			return true
		}
		if checkTag("json") || checkTag("graphql") {
			return true, fieldValue.Interface(), nil
		}

		if fieldValue.Kind() == reflect.Struct && fieldType.Anonymous {
			if ok, val, err := findFieldInStruct(fieldValue, fieldName); ok {
				return ok, val, err
			}
		}
		continue
	}

	return false, nil, nil
}
