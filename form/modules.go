package form

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "query",
			Target: UserQuery,
		},
	),
)
