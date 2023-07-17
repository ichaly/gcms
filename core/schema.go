package core

import (
	"github.com/graphql-go/graphql"
	"go.uber.org/fx"
)

type Schema[T any] interface {
	For() interface{}
	Name() string
	Description() string
	Resolve(params graphql.ResolveParams) (T, error)
}

type SchemaGroup struct {
	fx.In
	All []Schema[any] `group:"schema"`
}
