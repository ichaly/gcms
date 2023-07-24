package data

import (
	"github.com/ichaly/gcms/core"
)

type Model struct {
	core.Entity
}

func (Model) TableName() string {
	return "sys_model"
}
