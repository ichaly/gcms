package data

import (
	"github.com/ichaly/gcms/base"
)

type Channel struct {
	Name string `gorm:"size:200;comment:名称"`
	base.Entity
}

func (Channel) TableName() string {
	return "sys_channel"
}
