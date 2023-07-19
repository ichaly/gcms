package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

var Content = &data.Content{}

type Contents struct {
	db *gorm.DB
}

func NewContents(db *gorm.DB) core.Schema {
	return &Contents{db: db}
}

func (my *Contents) Resolve(_ graphql.ResolveParams) ([]*data.Content, error) {
	var res []*data.Content
	err := my.db.Model(&data.Content{}).Find(&res).Error
	return res, err
}

func (*Contents) Name() string {
	return "contents"
}

func (*Contents) Host() interface{} {
	return boot.Query
}

func (*Contents) Description() string {
	return "内容列表"
}
