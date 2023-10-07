package view

import (
	"github.com/ichaly/gcms/root/view/content"
	"github.com/ichaly/gcms/root/view/user"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	user.Modules,
	content.Modules,
)
