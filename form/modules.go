package form

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Name:   "query",
			Target: RootQuery,
		},
		fx.Annotated{
			Name:   "mutation",
			Target: RootMutation,
		},
		fx.Annotate(
			UserQuery,
			fx.ParamTags(`name:"query"`),
			fx.ResultTags(`group:"query"`),
		),
	),
)
