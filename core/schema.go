package core

import (
	"go.uber.org/fx"
)

type Schema interface {
	Name() string
	Host() interface{}
}

type SchemaGroup struct {
	fx.In
	All []Schema `group:"schema"`
}
