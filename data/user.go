package data

import (
	"github.com/ichaly/gcms/base"
)

type User struct {
	Name     string `gorm:"size:200;comment:名称"`
	Avatar   string `gorm:"size:200;comment:头像"`
	Nickname string `gorm:"size:50;comment:昵称"`
	Source   string `gorm:"comment:来源"`
	base.Entity
}

func (User) TableName() string {
	return "sys_user"
}

func (User) GqlDescription() string {
	return `用户管理`
}
