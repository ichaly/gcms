package zlog

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LevelEnablerFunc func(Level) bool

type AdapterOption struct {
	Out io.Writer
	LevelEnablerFunc
}

// NewAdapter 根据日志级别写入多个输出
// https://pkg.go.dev/go.uber.org/zap#example-package-AdvancedConfiguration
func NewAdapter(adapters []AdapterOption, opts ...Option) *Logger {
	var cores []zapcore.Core
	for _, a := range adapters {
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			zapcore.AddSync(a.Out),
			zap.LevelEnablerFunc(a.LevelEnablerFunc),
		)
		cores = append(cores, core)
	}
	return &Logger{l: zap.New(zapcore.NewTee(cores...), opts...)}
}
