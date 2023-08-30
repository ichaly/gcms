package ylog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	console := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	console.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	console.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	console.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	console.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	exception := NewRotateBySizeAndTime(&RotateConfig{
		LevelEnablerFunc: func(level Level) bool {
			return level > WarnLevel
		},
		Filename: "logs/error.log",
	})
	trace := NewRotateBySizeAndTime(&RotateConfig{
		LevelEnablerFunc: func(level Level) bool {
			return level <= WarnLevel
		},
		Filename: "logs/trace.log",
	})
	out := zerolog.MultiLevelWriter(console, exception, trace)
	l := New(out, DebugLevel)
	l.Trace().Int("id", 1).Msg("Trace测试日志")
	l.Debug().Int("id", 2).Msg("Debug测试日志")
	l.Warn().Int("id", 3).Msg("Warn测试日志")
	l.Error().Int("id", 4).Msg("Error测试日志")
}
