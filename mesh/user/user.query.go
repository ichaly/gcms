package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

var User = &data.User{}

type query struct {
	db *gorm.DB
}

func NewUserQuery(db *gorm.DB) base.Schema {
	return &query{db: db}
}

func (*query) Name() string {
	return "users"
}

func (*query) Description() string {
	return "用户列表"
}

func (*query) Host() interface{} {
	return base.Query
}

func (my *query) Type() interface{} {
	return []*data.User{}
}

func (my *query) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return core.QueryResolver[*data.User](p, my.db)
}