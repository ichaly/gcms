package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
	"strings"
)

func isPrimitive(p reflect.Type) bool {
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
		if !isPrimitive(b) {
			return typeInfo{}, fmt.Errorf("'%s' is not pointed to a base type", p.String())
		}
		return typeInfo{
			ptrType:  p,
			baseType: b,
			implType: b,
		}, nil
	default:
		if isPrimitive(p) {
			return typeInfo{
				ptrType:  reflect.New(p).Type(), // fixme: optimize for performance here
				baseType: p,
				implType: p,
			}, nil
		}
		return typeInfo{}, fmt.Errorf("unsupported type('%s') to unwrap", p.String())
	}
}

func implementsOf(p reflect.Type, intf reflect.Type) (implemented bool, info typeInfo, err error) {
	switch p.Kind() {
	case reflect.Slice, reflect.Array:
		e := p.Elem()
		if e.Kind() == reflect.Ptr || isPrimitive(e) {
			implemented, info, err = implementsOf(p.Elem(), intf)
			if err == nil {
				info.array = true
			}
		} else {
			err = fmt.Errorf("'%s' is illegal as an element of slice/array", e.String())
		}
	case reflect.Ptr:
		implemented = p.Implements(intf)
		if implemented {
			info.ptrType = p
			info.array = false
			info.implType = p
			info.baseType = p.Elem()
			if !isPrimitive(info.baseType) {
				err = fmt.Errorf("'%s' is not point to a base type", p.String())
			}
			return
		}
		b := p.Elem()
		if !isPrimitive(b) {
			err = fmt.Errorf("'%s' is not point to a base type", p.String())
			return
		}
		implemented = b.Implements(intf)
		if implemented {
			info.ptrType = p
			info.implType = b
			info.baseType = b
			info.array = false
		}
	default:
		if isPrimitive(p) {
			implemented = p.Implements(intf)
			if implemented {
				info.implType = p
				info.baseType = p
			}
			// try ptr
			pp := reflect.New(p).Type()
			info.ptrType = pp
			if implemented {
				return
			}

			implemented = pp.Implements(intf)
			if implemented {
				info.implType = pp
				info.baseType = p
			}
		}
	}
	return
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

func wrapType(p reflect.Type, t graphql.Type) graphql.Type {
	switch p.Kind() {
	case reflect.Slice, reflect.Array:
		return newList(wrapType(p.Elem(), t))
	case reflect.Ptr:
		return wrapType(p.Elem(), t)
	default:
		return newNonNull(t)
	}
}

func newPrototype(p reflect.Type) interface{} {
	elem := false
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	} else {
		elem = true
	}
	v := reflect.New(p)
	if elem {
		v = v.Elem()
	}
	return v.Interface()
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
