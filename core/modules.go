package core

import (
	"github.com/ichaly/gcms/boot"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		boot.NewEngine,
		NewConfig,
		NewStore,
		NewCache,
		NewRender,
		NewRouter,
		NewConnect,
		fx.Annotate(
			NewGraphql,
			fx.As(new(Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
	fx.Invoke(Bootstrap),
)
