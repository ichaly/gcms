package mesh

import (
	"github.com/ichaly/gcms/mesh/content"
	"github.com/ichaly/gcms/mesh/user"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	user.Modules,
	content.Modules,
)
