package user

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "schema",
			Target: NewUsers,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewAge,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewContents,
		},
	),
)
