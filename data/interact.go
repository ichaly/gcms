package data

import (
	"github.com/ichaly/gcms/core"
)

type Interact struct {
	core.Primary
	ContentId int64 `gorm:"comment:内容ID"`
	View      int   `gorm:"comment:阅读量"`
	Like      int   `gorm:"comment:点赞量"`
	Share     int   `gorm:"comment:分享量"`
	Comment   int   `gorm:"comment:评论量"`
}

func (Interact) TableName() string {
	return "sys_interact"
}
