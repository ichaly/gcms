package main

import (
	"github.com/ichaly/gcms/main/cmd"
	"github.com/ichaly/gcms/zlog"
	"os"
)

func main() {
	defer zlog.Sync()
	out := []zlog.AdapterOption{
		{
			Out: os.Stdout,
			LevelEnablerFunc: func(level zlog.Level) bool {
				return level <= zlog.FatalLevel
			},
		},
		{
			Out: zlog.NewProductionRotateByTime("logs/error.log"),
			LevelEnablerFunc: func(level zlog.Level) bool {
				return level > zlog.WarnLevel
			},
		},
		{
			Out: zlog.NewProductionRotateByTime("logs/trace.log"),
			LevelEnablerFunc: func(level zlog.Level) bool {
				return level <= zlog.WarnLevel
			},
		},
	}
	logger := zlog.NewAdapter(out)
	zlog.ReplaceDefault(logger)
	cmd.Execute()
}
