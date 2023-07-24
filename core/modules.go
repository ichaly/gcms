package core

import (
	"github.com/ichaly/gcms/base"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		base.NewEngine,
		NewConfig,
		NewStorage,
		NewConnect,
		NewCache,
		NewRender,
		NewRouter,
		fx.Annotate(
			NewGraphql,
			fx.As(new(Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
	fx.Invoke(Bootstrap),
)
