package core

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
	Name() string
	Protected() bool
	Init(*gin.RouterGroup)
}
