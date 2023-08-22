package core

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"time"
)

var (
	Version   string
	GitHash   string
	BuildTime string
	routers   = make(map[string]gin.IRouter)
)

func Bootstrap(l fx.Lifecycle, c *Config, e *gin.Engine, g PluginGroup) {
	if Version == "" {
		Version = "v0.0.0"
	}
	if GitHash == "" {
		GitHash = "Unknown"
	}
	if BuildTime == "" {
		BuildTime = time.Now().Format("2006-01-02 15:04:05")
	}
	all := append(g.Middlewares, g.Plugins...)
	for _, p := range all {
		r, ok := routers[p.Base()]
		if p.Base() == "" {
			r = e
		} else if !ok {
			r = e.Group(p.Base())
			routers[p.Base()] = r
		}
		p.Init(r)
	}
	srv := &http.Server{Addr: fmt.Sprintf(":%v", c.App.Port), Handler: e}
	l.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Printf("当前版本:%s-%s 发布日期:%s\n", Version, GitHash, BuildTime)
				err := srv.ListenAndServe()
				if err != nil {
					fmt.Printf("%v failed to start: %v", c.App.Name, err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Printf("%v shutdown complete", c.App.Name)
			return srv.Shutdown(ctx)
		},
	})
}
