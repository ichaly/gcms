package zlog

import (
	"github.com/rs/zerolog"
	"io"
)

type LoggerOption func(zerolog.Logger) zerolog.Logger

type RotateOption func(*rotate)

func WithOut(out io.Writer) LoggerOption {
	return func(l zerolog.Logger) zerolog.Logger {
		return l.Output(out)
	}
}

func WithLevel(level Level) LoggerOption {
	return func(l zerolog.Logger) zerolog.Logger {
		return l.Level(level)
	}
}

func WithLevelFunc(fn LevelFunc) RotateOption {
	return func(r *rotate) {
		r.LevelFunc = fn
	}
}

func WithFilename(filename string) RotateOption {
	return func(r *rotate) {
		r.Filename = filename
	}
}

func WithMaxAge(maxAge int) RotateOption {
	return func(r *rotate) {
		r.MaxAge = maxAge
	}
}

func WithMaxSize(maxSize int) RotateOption {
	return func(r *rotate) {
		r.MaxSize = maxSize
	}
}

func WithMaxBackups(maxBackups int) RotateOption {
	return func(r *rotate) {
		r.MaxBackups = maxBackups
	}
}

func UseCompress(compress bool) RotateOption {
	return func(r *rotate) {
		r.Compress = compress
	}
}

func UseLocalTime(localTime bool) RotateOption {
	return func(r *rotate) {
		r.LocalTime = localTime
	}
}

func UseDaily(daily bool) RotateOption {
	return func(r *rotate) {
		r.Daily = daily
	}
}
