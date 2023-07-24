package base

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"reflect"
	"runtime"
	"strings"
)

func isPrim(p reflect.Type) bool {
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

func unwrap(p reflect.Type) (reflect.Type, error) {
	switch p.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Array:
		return unwrap(p.Elem())
	default:
		if !isPrim(p) {
			return nil, fmt.Errorf("unsupported type('%s') to unwrap", p.String())
		}
		return p, nil
	}
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
		//return newNonNull(t)
		return t
	}
}

func parseType(typ reflect.Type, errString string, parsers ...typeParser) (graphql.Type, error) {
	for _, check := range parsers {
		res, err := check(typ)
		if err != nil {
			return nil, err
		}
		if res == nil {
			continue
		}
		return res, nil
	}
	return nil, fmt.Errorf("unsupported type('%s') for %s", typ.String(), errString)
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

func getFuncName(fn interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	name = name[strings.LastIndex(name, ".")+1:]
	name = strings.ReplaceAll(name, "-fm", "")
	name = strings.Replace(name, "Get", "", 1)
	name = strings.Replace(name, "Resolve", "", 1)
	name = strcase.ToLowerCamel(name)
	return name
}
