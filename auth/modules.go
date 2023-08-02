package auth

import (
	"github.com/ichaly/gcms/core"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewCasbin,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
)
