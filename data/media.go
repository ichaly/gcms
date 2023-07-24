package data

import (
	"github.com/ichaly/gcms/core"
)

type Media struct {
	ContentId   int64  `gorm:"comment:内容ID"`
	Url         string `gorm:"size:200;comment:路径"`
	Name        string `gorm:"size:200;comment:名称"`
	Size        int64  `gorm:"comment:大小"`
	Width       int64  `gorm:"comment:宽"`
	Height      int64  `gorm:"comment:高"`
	Description string `gorm:"type:text;comment:描述"`
	core.Entity
}

func (Media) TableName() string {
	return "sys_media"
}
