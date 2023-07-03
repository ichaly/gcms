package core

import (
	"github.com/graphql-go/graphql"
)

func (my *Engine) buildExpressionInput(t graphql.Type) graphql.Type {
	var name string
	var list []__inputValue
	var isList bool

	list = append(list, expNull...)
	if val, ok := t.(*graphql.NonNull); ok {
		t = val.OfType
	}
	if val, ok := t.(*graphql.List); ok {
		t = val.OfType
		isList = true
	}
	if val, ok := t.(*graphql.NonNull); ok {
		t = val.OfType
	}
	//if val, ok := t.(*graphql.Enum); ok {
	//	t = val
	//	isEnum = true
	//}
	if val, ok := t.(*graphql.Object); ok {
		return my.buildWhereInput(val)
	}

	if isList {
		list = append(list, expList...)
	} else {
		list = append(list, expScalar...)
	}

	if isList {
		name = t.Name() + "ListExpression"
	} else {
		name = t.Name() + "Expression"
	}
	typ, ok := my.types[name]
	if ok {
		return typ
	}

	fields := graphql.InputObjectConfigFieldMap{}
	for _, e := range list {
		if e.Type == nil {
			e.Type = t
		}
		fields[e.Name] = &graphql.InputObjectFieldConfig{Type: e.Type, Description: e.Description}
	}
	input := graphql.NewInputObject(graphql.InputObjectConfig{Name: name, Fields: fields})
	my.types[name] = input

	return input
}

func (my *Engine) buildSortInput(object *graphql.Object) graphql.Type {
	name := object.Name() + "SortInput"
	val, ok := my.types[name]
	if ok {
		return val
	}

	fields := graphql.InputObjectConfigFieldMap{}
	for k, _ := range object.Fields() {
		fields[k] = &graphql.InputObjectFieldConfig{Type: SortDirection}
	}
	input := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: name, Fields: fields,
	})

	my.types[name] = input
	return input
}

func (my *Engine) buildWhereInput(object *graphql.Object) graphql.Type {
	name := object.Name() + "WhereInput"
	val, ok := my.types[name]
	if ok {
		return val
	}
	fields := graphql.InputObjectConfigFieldMap{}
	for k, v := range object.Fields() {
		typ := my.buildExpressionInput(v.Type)
		if typ != nil {
			fields[k] = &graphql.InputObjectFieldConfig{Type: typ}
		}
	}
	input := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: name, Fields: fields,
	})

	fields["or"] = &graphql.InputObjectFieldConfig{Type: input}
	fields["and"] = &graphql.InputObjectFieldConfig{Type: input}
	fields["not"] = &graphql.InputObjectFieldConfig{Type: input}

	my.types[name] = input
	return input
}
