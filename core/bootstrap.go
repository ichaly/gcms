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
	Version   = "v0.0.0"
	GitHash   = "Unknown"
	BuildTime = time.Now().Format("2006-01-02 15:04:05")
	routers   = make(map[string]gin.IRouter)
)

func Bootstrap(l fx.Lifecycle, c *Config, e *gin.Engine, g PluginGroup) {
	all := append(g.Middlewares, g.Plugins...)
	for _, p := range all {
		r, ok := routers[p.Base()]
		if !ok {
			r = e.Group(p.Base())
			routers[p.Base()] = r
		}
		p.Init(r)
	}
	srv := &http.Server{Addr: fmt.Sprintf(":%v", c.App.Port), Handler: e}
	l.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go startServer(srv, c)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return stopServer(srv, c)
		},
	})
}

func startServer(srv *http.Server, c *Config) {
	fmt.Printf("当前版本:%s-%s 发布日期:%s\n", Version, GitHash, BuildTime)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("%v failed to start: %v", c.App.Name, err)
	}
}

func stopServer(srv *http.Server, c *Config) error {
	fmt.Printf("%v shutdown complete", c.App.Name)
	return srv.Shutdown(context.Background())
}
