package zlog

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"time"
)

type LevelFunc func(Level) bool

type rotate struct {
	LevelFunc
	Daily bool
	*lumberjack.Logger
}

func (my rotate) WriteLevel(l Level, p []byte) (n int, err error) {
	if my.LevelFunc(l) {
		return my.Write(p)
	}
	return len(p), nil
}

func NewRotate(ops ...RotateOption) io.Writer {
	r := &rotate{Logger: &lumberjack.Logger{}}
	for _, o := range ops {
		o(r)
	}
	if r.Daily {
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
				_ = r.Rotate()
			}
		}()
	}
	return r
}
