package core

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewConfig,
			fx.ParamTags(`name:"configFile"`),
		),
		NewCache,
		NewStore,
		NewRender,
		NewRouter,
		NewConnect,
	),
	fx.Invoke(Bootstrap),
)
