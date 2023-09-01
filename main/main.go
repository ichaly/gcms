package main

import (
	"fmt"
	"github.com/ichaly/gcms/main/cmd"
	"github.com/ichaly/gcms/zlog"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}
	errors := zlog.NewRotate(
		zlog.WithFilename("logs/error.log"),
		zlog.WithLevelFunc(func(l zlog.Level) bool { return l > zlog.WarnLevel }),
	)
	traces := zlog.NewRotate(
		zlog.WithFilename("logs/trace.log"),
		zlog.WithLevelFunc(func(l zlog.Level) bool { return l <= zlog.WarnLevel }),
	)
	out := zerolog.MultiLevelWriter(output, errors, traces)
	l := zlog.New(zlog.WithOut(out), zlog.WithLevel(zlog.DebugLevel))
	zlog.SetDefault(l)

	cmd.Execute()
}
