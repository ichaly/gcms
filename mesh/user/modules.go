package user

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "schema",
			Target: NewUserQuery,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewUserMutation,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewUserAge,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewUserToken,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewUserContents,
		},
	),
)
