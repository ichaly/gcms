package data

import (
	"github.com/ichaly/gcms/core"
	"time"
)

type User struct {
	Name     string    `gorm:"index;size:200;comment:名称"`
	Avatar   string    `gorm:"size:200;comment:头像"`
	Nickname string    `gorm:"size:50;comment:昵称"`
	Source   string    `gorm:"comment:来源"`
	Birthday time.Time `gorm:"comment:生日"`
	core.Entity
}

func (User) TableName() string {
	return "sys_user"
}

func (User) Description() string {
	return "用户信息"
}
