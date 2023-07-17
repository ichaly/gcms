package mesh

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

const Content = "Content"

type ContentSchema struct {
	db *gorm.DB
}

func NewContentSchema(d *gorm.DB, e *boot.Engine) core.Schema {
	my := &ContentSchema{db: d}
	e.NewQuery(my.GetContests)
	return my
}

func (my *ContentSchema) GetContests(p graphql.ResolveParams) ([]*data.Content, error) {
	var res []*data.Content
	err := my.db.Model(&data.Content{}).Find(&res).Error
	return res, err
}
