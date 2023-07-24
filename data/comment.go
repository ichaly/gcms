package data

import (
	"github.com/ichaly/gcms/core"
)

type Comment struct {
	Content string `gorm:"type:text;comment:内容"`
	core.Entity
}

func (Comment) TableName() string {
	return "sys_comment"
}
