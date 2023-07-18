package mesh

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

type ContentSchema struct {
	db *gorm.DB
}

func NewContentSchema(db *gorm.DB) *ContentSchema {
	return &ContentSchema{db: db}
}

func (my *ContentSchema) ResolveContentsForQuery(_ graphql.ResolveParams) ([]*data.Content, error) {
	var res []*data.Content
	err := my.db.Model(&data.Content{}).Find(&res).Error
	return res, err
}
