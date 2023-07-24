package user

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "schema",
			Target: NewMutation,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewList,
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
