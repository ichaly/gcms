package auth

import (
	"github.com/ichaly/gcms/core"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		NewEnforcer,
		fx.Annotate(
			NewCros,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewCasbin,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewGraphql,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewIndex,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
)
