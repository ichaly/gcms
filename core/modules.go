package core

import (
	"github.com/ichaly/gcms/base"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		base.NewEngine,
		NewConfig,
		NewServer,
		NewStorage,
		NewValidate,
		fx.Annotate(
			NewConnect,
			fx.ParamTags(``, `group:"gorm"`, `group:"entity"`),
		),
		fx.Annotated{
			Group:  "gorm",
			Target: NewSonyFlake,
		},
		fx.Annotated{
			Group:  "gorm",
			Target: NewCache,
		},
		fx.Annotate(
			NewGraphql,
			fx.As(new(Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
	fx.Invoke(Bootstrap),
)
