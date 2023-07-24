package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

var User = &data.User{}

type list struct {
	db *gorm.DB
}

func NewList(db *gorm.DB) core.Schema {
	return &list{db: db}
}

func (*list) Name() string {
	return "users"
}

func (*list) Host() interface{} {
	return boot.Query
}

func (*list) Description() string {
	return "用户列表"
}

func (my *list) Resolve(_ graphql.ResolveParams) ([]*data.User, error) {
	var res []*data.User
	err := my.db.Model(&data.User{}).Find(&res).Error
	return res, err
}
