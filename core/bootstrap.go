package core

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"time"
)

var (
	Version   string
	GitHash   string
	BuildTime string
)

func Bootstrap(l fx.Lifecycle, c *Config, h *chi.Mux, p PluginGroup) {
	Version = "0.0.1"
	GitHash = "Unknown"
	BuildTime = time.Now().Format("2006-01-02 15:04:05")
	h.Group(func(h chi.Router) {
		for _, p := range p.Plugins {
			p.Init(h)
		}
	})
	srv := &http.Server{Addr: fmt.Sprintf(":%v", c.App.Port), Handler: h}
	l.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Printf("当前版本:%s-%s 发布日期:%s\n", Version, GitHash, BuildTime)
				err := srv.ListenAndServe()
				if err != nil {
					fmt.Printf("%v failed to start: %v", c.App.Name, err)
					return
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