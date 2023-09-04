package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/apps/data"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"gorm.io/gorm"
)

type mutation struct {
	db       *gorm.DB
	validate *base.Validate
}

func NewContentMutation(d *gorm.DB, v *base.Validate) core.Schema {
	return &mutation{db: d, validate: v}
}

func (*mutation) Name() string {
	return "contents"
}

func (*mutation) Description() string {
	return "内容管理"
}

func (*mutation) Host() interface{} {
	return core.Mutation
}

func (*mutation) Type() interface{} {
	return Content
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.MutationResolver[*data.Content](p, my.db, my.validate)
}
