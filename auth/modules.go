package auth

import (
	"github.com/ichaly/gcms/core"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		NewEnforcer,
		fx.Annotate(
			NewCasbin,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
		fx.Annotate(
			NewGraphql,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
)
