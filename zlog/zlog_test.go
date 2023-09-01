package zlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}
	errors := NewRotate(WithFilename("logs/error.log"), WithLevelFunc(func(l Level) bool {
		return l > WarnLevel
	}))
	traces := NewRotate(WithFilename("logs/trace.log"), WithLevelFunc(func(l Level) bool {
		return l <= WarnLevel
	}))
	out := zerolog.MultiLevelWriter(output, errors, traces)
	l := New(WithOut(out), WithLevel(DebugLevel))
	SetDefault(l)
	for i := 0; i < 100; i++ {
		Trace().Int("id", 1).Msgf("Trace测试日志%d", i)
		Debug().Int("id", 2).Msgf("Debug测试日志%d", i)
		Warn().Int("id", 3).Msgf("Warn测试日志%d", i)
		Error().Int("id", 4).Msgf("Error测试日志%d", i)
		time.Sleep(10 * time.Second)
	}
}
