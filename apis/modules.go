package apis

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Name:   "schema",
			Target: NewSchema,
		},
	),
)
