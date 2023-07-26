package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type mutation struct {
	db *gorm.DB
}

func NewUserMutation(db *gorm.DB) core.Schema {
	return &mutation{db: db}
}

func (*mutation) Name() string {
	return "users"
}

func (*mutation) Host() interface{} {
	return base.Mutation
}

func (*mutation) Description() string {
	return "用户管理"
}

func (my *mutation) Resolve(p graphql.ResolveParams) (*data.User, error) {
	var args core.Params[*data.User]
	err := mapstructure.WeakDecode(p.Args, &args)
	if err != nil {
		return nil, err
	}
	tx := my.db.Model(User)
	if args.Where != nil {
		core.ParseWhere(args.Where, tx)
	}
	if args.Delete {
		err = tx.Delete(&args.Data).Error
		return nil, err
	}
	if args.Data == nil {
		return nil, nil
	}
	if args.Data.ID > 0 {
		err = tx.Updates(&args.Data).Error
	} else {
		err = tx.Create(&args.Data).Error
	}
	if err != nil {
		return nil, err
	}
	return args.Data, nil
}
