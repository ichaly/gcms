package auth

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		//Casbin鉴权
		NewEnforcer,
		//Oauth2认证
		NewOauthServer,
		NewOauthTokenStore,
		NewOauthClientStore,
		fx.Annotate(
			NewOauth,
			fx.ResultTags(`group:"plugin"`),
		),
		//Graphql鉴权中间件
		fx.Annotate(
			NewGraphql,
			fx.ResultTags(`group:"middleware"`),
		),
		//登录验证中间件
		fx.Annotate(
			NewOauthVerify,
			fx.ResultTags(`group:"middleware"`),
		),
	),
)
