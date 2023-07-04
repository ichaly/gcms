package data

import "github.com/ichaly/gcms/base"

type Content struct {
	Title       string `gorm:"size:200;comment:标题"`
	Description string `gorm:"size:200;comment:简介"`
	Content     string `gorm:"type:text;comment:内容"`
	Source      string `gorm:"comment:来源"`
	Kind        Kind   `gorm:"comment:类型"`
	Media       []Media
	base.DeleteEntity
}

func (Content) TableName() string {
	return "sys_content"
}
