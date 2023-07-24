package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

var Content = &data.Content{}

type contents struct {
	db *gorm.DB
}

func NewContents(db *gorm.DB) core.Schema {
	return &contents{db: db}
}

func (my *contents) Resolve(_ graphql.ResolveParams) ([]*data.Content, error) {
	var res []*data.Content
	err := my.db.Model(&data.Content{}).Find(&res).Error
	return res, err
}

func (*contents) Name() string {
	return "contents"
}

func (*contents) Host() interface{} {
	return base.Query
}

func (*contents) Description() string {
	return "内容列表"
}
