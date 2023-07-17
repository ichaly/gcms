package mesh

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "schema",
			Target: NewUserSchema,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewContentSchema,
		},
	),
)
