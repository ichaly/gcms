package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
	"strconv"
	"strings"
)

type typeInfo struct {
	array    bool
	ptrType  reflect.Type
	implType reflect.Type
	baseType reflect.Type
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

func unwrap(p reflect.Type) (typeInfo, error) {
	switch p.Kind() {
	case reflect.Slice, reflect.Array:
		info, err := unwrap(p.Elem())
		if err == nil {
			info.array = true
		}
		return info, err
	case reflect.Ptr:
		b := p.Elem()
		if !isBaseType(b) {
			return typeInfo{}, fmt.Errorf("'%s' is not pointed to a base type", p.String())
		}
		return typeInfo{
			ptrType:  p,
			baseType: b,
			implType: b,
		}, nil
	default:
		if isBaseType(p) {
			return typeInfo{
				baseType: p,
				ptrType:  reflect.New(p).Type(), // fixme: optimize for performance here
				implType: p,
			}, nil
		}
		return typeInfo{}, fmt.Errorf("unsupported type('%s') to unwrap", p.String())
	}
}

func isBaseType(p reflect.Type) bool {
	switch p.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool,
		reflect.String,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Struct,
		reflect.Interface:
		return true
	}
	return false
}

func newNonNull(t graphql.Type) graphql.Type {
	if _, ok := t.(*graphql.NonNull); !ok {
		t = graphql.NewNonNull(t)
	}
	return t
}

func newList(t graphql.Type) graphql.Type {
	if _, ok := t.(*graphql.List); !ok {
		t = graphql.NewList(t)
	}
	return t
}

func wrapType(field *reflect.StructField, t graphql.Type, isArray bool) graphql.Type {
	if isArray {
		if isElementRequired(field) {
			t = newNonNull(t)
		}
		t = newList(t)
		if isRequired(field) {
			t = newNonNull(t)
		}
	} else {
		if isRequired(field) {
			t = newNonNull(t)
		}
	}
	return t
}

func boolTag(field *reflect.StructField, tagName string) bool {
	v, ok := field.Tag.Lookup(tagName)
	if !ok {
		return false
	}
	if v == "" {
		return true
	}
	positive, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return positive
}

func isRequired(field *reflect.StructField) bool { return boolTag(field, "gqlRequired") }

func isElementRequired(field *reflect.StructField) bool { return boolTag(field, "gqlElementRequired") }
