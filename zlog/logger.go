package zlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

type Level = zerolog.Level

const (
	// TraceLevel defines trace log level.
	TraceLevel = zerolog.TraceLevel
	// DebugLevel defines debug log level.
	DebugLevel = zerolog.DebugLevel
	// InfoLevel defines info log level.
	InfoLevel = zerolog.InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel = zerolog.WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel = zerolog.ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel = zerolog.FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel = zerolog.PanicLevel
	// NoLevel defines an absent log level.
	NoLevel = zerolog.NoLevel
	// Disabled disables the logger.
	Disabled = zerolog.Disabled
)

type Logger struct {
	l zerolog.Logger
}

func New(ops ...LoggerOption) *Logger {
	console := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	console.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	console.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}
	l := zerolog.New(console).With().Timestamp().Logger()
	for _, o := range ops {
		l = o(l)
	}
	//zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	zerolog.TimestampFieldName = "t"
	return &Logger{l: l}
}

func (my *Logger) SetLevel(level Level) {
	my.l.Level(level)
}
func (my *Logger) Trace() *zerolog.Event {
	return my.l.Trace()
}
func (my *Logger) Debug() *zerolog.Event {
	return my.l.Debug()
}
func (my *Logger) Info() *zerolog.Event {
	return my.l.Info()
}
func (my *Logger) Warn() *zerolog.Event {
	return my.l.Warn()
}
func (my *Logger) Error() *zerolog.Event {
	return my.l.Error()
}
func (my *Logger) Fatal() *zerolog.Event {
	return my.l.Fatal()
}
func (my *Logger) Panic() *zerolog.Event {
	return my.l.Panic()
}

var std = New(WithOut(os.Stderr), WithLevel(InfoLevel))

func Default() *Logger     { return std }
func SetDefault(l *Logger) { std = l }
func SetLevel(level Level) { std.SetLevel(level) }
func Trace() *zerolog.Event {
	return std.Trace()
}
func Debug() *zerolog.Event {
	return std.Debug()
}
func Info() *zerolog.Event {
	return std.Info()
}
func Warn() *zerolog.Event {
	return std.Warn()
}
func Error() *zerolog.Event {
	return std.Error()
}
func Fatal() *zerolog.Event {
	return std.Fatal()
}
func Panic() *zerolog.Event {
	return std.Panic()
}
