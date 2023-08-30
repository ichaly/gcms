package ylog

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
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

var logger Logger

type Logger struct {
	l zerolog.Logger
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

func New(out io.Writer, level Level) *Logger {
	var writers []io.Writer
	if out == nil {
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
		out = console
		//out = os.Stderr
		//out = zerolog.MultiLevelWriter(console, os.Stdout)
	}
	writers = append(writers, out)
	//zerolog.SetGlobalLevel(level)
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//zerolog.TimestampFieldName = "t"
	//zerolog.LevelFieldName = "l"
	//zerolog.MessageFieldName = "m"
	return &Logger{l: zerolog.New(zerolog.MultiLevelWriter(writers...)).With().Timestamp().Logger().Level(level)}
}

var std = New(os.Stderr, InfoLevel)

func Default() *Logger     { return std }
func SetDefault(l *Logger) { std = l }
func SetLevel(level Level) { std.SetLevel(level) }

func Trace() *zerolog.Event {
	return logger.Trace()
}
func Debug() *zerolog.Event {
	return logger.Debug()
}
func Info() *zerolog.Event {
	return logger.Info()
}
func Warn() *zerolog.Event {
	return logger.Warn()
}
func Error() *zerolog.Event {
	return logger.Error()
}
func Fatal() *zerolog.Event {
	return logger.Fatal()
}
func Panic() *zerolog.Event {
	return logger.Panic()
}
