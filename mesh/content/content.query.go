package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

var Content = &data.Content{}

type query struct {
	db *gorm.DB
}

func NewContentQuery(db *gorm.DB) core.Schema {
	return &query{db: db}
}

func (my *query) Resolve(_ graphql.ResolveParams) ([]*data.Content, error) {
	var res []*data.Content
	err := my.db.Model(&data.Content{}).Find(&res).Error
	return res, err
}

func (*query) Name() string {
	return "query"
}

func (*query) Host() interface{} {
	return base.Query
}

func (*query) Description() string {
	return "内容列表"
}
