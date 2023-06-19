package data

import (
	"github.com/ichaly/gcms/base"
)

type Comment struct {
	Content string `gorm:"type:text;comment:内容"`
	base.Entity
}

func (Comment) TableName() string {
	return "sys_comment"
}
