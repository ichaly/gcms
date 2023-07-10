package apis

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

const Content = "Content"

type ContentApi struct {
	db *gorm.DB
}

func NewContentApi(d *gorm.DB, e *boot.Engine) core.Schema {
	my := &ContentApi{db: d}
	e.NewQuery(my.GetContests).Name("contents")
	return my
}

func (my *ContentApi) GetContests(p graphql.ResolveParams) ([]*data.Content, error) {
	return nil, nil
}
