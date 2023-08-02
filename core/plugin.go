package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PluginGroup struct {
	fx.In
	All []Plugin `group:"plugin"`
}

type Plugin interface {
	Name() string
	Init(*gin.RouterGroup)
}
