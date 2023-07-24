package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

type mutation struct {
	db *gorm.DB
}

func NewMutation(db *gorm.DB) core.Schema {
	return &mutation{db: db}
}

func (*mutation) Name() string {
	return "users"
}

func (*mutation) Host() interface{} {
	return boot.Mutation
}

func (*mutation) Description() string {
	return "用户管理"
}

func (my *mutation) Resolve(_ graphql.ResolveParams) ([]*data.User, error) {
	var res []*data.User
	err := my.db.Model(&data.User{}).Find(&res).Error
	return res, err
}
