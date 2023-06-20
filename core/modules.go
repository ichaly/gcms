package core

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewConfig,
			fx.ParamTags(`name:"configFile"`),
		),
		NewStore,
		NewCache,
		NewRender,
		NewRouter,
		NewConnect,
		fx.Annotate(
			NewEngine,
			fx.As(new(Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
	fx.Invoke(Bootstrap),
)
