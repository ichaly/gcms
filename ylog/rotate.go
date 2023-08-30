package ylog

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"time"
)

type LevelEnablerFunc func(Level) bool

type RotateConfig struct {
	LevelEnablerFunc //日志级别
	// 共用配置
	Filename string // 完整文件名
	MaxAge   int    // 保留旧日志文件的最大天数

	// 按时间轮转配置
	//RotationTime time.Duration // 日志文件轮转时间

	// 按大小轮转配置
	MaxSize    int  // 日志文件最大大小（MB）
	MaxBackups int  // 保留日志文件的最大数量
	Compress   bool // 是否对日志文件进行压缩归档
	LocalTime  bool // 是否使用本地时间，默认 UTC 时间
}

type rotate struct {
	*lumberjack.Logger
	config *RotateConfig
}

func (my rotate) WriteLevel(l Level, p []byte) (n int, err error) {
	if my.config.LevelEnablerFunc(l) {
		return my.Write(p)
	}
	return len(p), nil
}

func NewRotateBySizeAndTime(cfg *RotateConfig) io.Writer {
	loggerWrite := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}
	go func() {
		for {
			now := time.Now()
			layout := "2006-01-02"
			//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
			today, _ := time.ParseInLocation(layout, now.Format(layout), time.Local)
			// 第二天零点时间戳
			next := today.AddDate(0, 0, 1)
			after := next.UnixNano() - now.UnixNano() - 1
			<-time.After(time.Duration(after) * time.Nanosecond)
			_ = loggerWrite.Rotate()
		}
	}()
	return rotate{loggerWrite, cfg}
}
