package data

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotate(
			func() interface{} {
				return &User{}
			},
			fx.ResultTags(`group:"entity"`),
		),
	),
)
