package core

import "go.uber.org/fx"

type Schema interface {
}

type SchemaGroup struct {
	fx.In
	All []Schema `group:"schema"`
}
