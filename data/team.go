package data

import (
	"github.com/ichaly/gcms/core"
)

type Team struct {
	Name string `gorm:"size:200;comment:名称"`
	core.Entity
}

func (Team) TableName() string {
	return "sys_team"
}
