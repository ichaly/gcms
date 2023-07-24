package data

import (
	"github.com/ichaly/gcms/core"
)

type Channel struct {
	Name string `gorm:"size:200;comment:名称"`
	core.Entity
}

func (Channel) TableName() string {
	return "sys_channel"
}
