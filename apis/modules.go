package apis

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "api",
			Target: NewUserApi,
		},
		fx.Annotated{
			Group:  "api",
			Target: NewContentApi,
		},
	),
)
