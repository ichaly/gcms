package root

import (
	"github.com/ichaly/gcms/root/data"
	"github.com/ichaly/gcms/root/serv"
	"github.com/ichaly/gcms/root/view"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	data.Modules,
	serv.Modules,
	view.Modules,
)
