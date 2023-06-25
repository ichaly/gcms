package form

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		RootQuery,
		RootMutation,
		fx.Annotate(
			UserQuery,
			fx.ParamTags(`name:"rootQuery"`),
		),
		fx.Annotate(
			ContentQuery,
			fx.ParamTags(`name:"userQuery"`),
		),
	),
)
