package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

var User = &data.User{}

type users struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB, e *boot.Engine) core.Schema {
	return &users{db: db}
}

func (*users) Name() string {
	return "users"
}

func (*users) Host() interface{} {
	return boot.Query
}

func (*users) Description() string {
	return "用户列表"
}

func (my *users) Resolve(_ graphql.ResolveParams) ([]*data.User, error) {
	var res []*data.User
	err := my.db.Model(&data.User{}).Find(&res).Error
	return res, err
}
