package data

import (
	"github.com/ichaly/gcms/base"
)

type Team struct {
	Name string `gorm:"size:200;comment:名称"`
	base.Entity
}

func (Team) TableName() string {
	return "sys_team"
}
