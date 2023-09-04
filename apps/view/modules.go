package view

import (
	"github.com/ichaly/gcms/apps/view/content"
	"github.com/ichaly/gcms/apps/view/user"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	user.Modules,
	content.Modules,
)
