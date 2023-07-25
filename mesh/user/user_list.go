package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"github.com/mitchellh/mapstructure"
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
	return base.Query
}

func (*list) Description() string {
	return "用户列表"
}

func (my *list) Resolve(p graphql.ResolveParams) ([]*data.User, error) {
	var args core.Params[*data.User]
	err := mapstructure.WeakDecode(p.Args, &args)
	if err != nil {
		return nil, err
	}

	tx := my.db.Model(User)
	if args.Where != nil {
		core.ParseWhere(args.Where, tx)
	}

	var res []*data.User
	err = tx.Find(&res).Error
	return res, err
}
