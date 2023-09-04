package base

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PluginGroup struct {
	fx.In
	Plugins     []Plugin `group:"plugin"`
	Middlewares []Plugin `group:"middleware"`
}

type Plugin interface {
	Base() string
	Init(gin.IRouter)
}
