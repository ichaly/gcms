package data

import (
	"github.com/ichaly/gcms/base"
)

type Model struct {
	base.Entity
}

func (Model) TableName() string {
	return "sys_model"
}
