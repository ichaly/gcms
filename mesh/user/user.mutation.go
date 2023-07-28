package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

type mutation struct {
	db *gorm.DB
}

func NewUserMutation(db *gorm.DB) base.Schema {
	return &mutation{db: db}
}

func (*mutation) Name() string {
	return "users"
}

func (*mutation) Description() string {
	return "用户管理"
}

func (*mutation) Host() interface{} {
	return base.Mutation
}

func (*mutation) Type() interface{} {
	return User
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return core.MutationResolver[*data.User](p, my.db)
}
