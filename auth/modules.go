package auth

import (
	"github.com/ichaly/gcms/core"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewCros,
			fx.ResultTags(`group:"middleware"`),
		),
		//Oauth2认证
		NewOauthServer,
		NewOauthTokenStore,
		NewOauthClientStore,
		fx.Annotate(
			NewOauthVerify,
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewOauth,
			fx.ResultTags(`group:"plugin"`),
		),
		//Casbin鉴权
		NewEnforcer,
		fx.Annotate(
			NewCasbin,
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewGraphql,
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewIndex,
			fx.As(new(core.Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
)
