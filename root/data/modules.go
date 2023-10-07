package data

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &User{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Media{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Content{}
			},
		},
	),
)
