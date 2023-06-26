package serv

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
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
