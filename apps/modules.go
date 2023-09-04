package apps

import (
	"github.com/ichaly/gcms/apps/data"
	"github.com/ichaly/gcms/apps/serv"
	"github.com/ichaly/gcms/apps/view"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	data.Modules,
	serv.Modules,
	view.Modules,
)
