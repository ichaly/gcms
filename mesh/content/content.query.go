package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

var Content = &data.Content{}

type query struct {
	db *gorm.DB
}

func NewContentQuery(db *gorm.DB) base.Schema {
	return &query{db: db}
}

func (*query) Name() string {
	return "contents"
}

func (*query) Description() string {
	return "内容列表"
}

func (*query) Host() interface{} {
	return base.Query
}

func (*query) Type() interface{} {
	return []*data.Content{}
}

func (my *query) Resolve(p graphql.ResolveParams) (interface{}, error) {
	var res []*data.Content
	err := my.db.WithContext(p.Context).Model(&data.Content{}).Find(&res).Error
	return res, err
}
