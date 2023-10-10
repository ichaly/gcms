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
		//跨域中间件
		fx.Annotate(
			NewCros,
			fx.ResultTags(`group:"middleware"`),
		),
		//鉴权中间件
		//fx.Annotate(
		//	NewCasbin,
		//	fx.ResultTags(`group:"middleware"`),
		//),
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
		fx.Annotate(
			NewOauth,
			fx.ResultTags(`group:"plugin"`),
		),
		fx.Annotate(
			NewIndex,
			fx.ResultTags(`group:"plugin"`),
		),
	),
)
