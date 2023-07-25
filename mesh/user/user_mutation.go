package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return base.Mutation
}

func (*mutation) Description() string {
	return "用户管理"
}

func (my *mutation) Resolve(p graphql.ResolveParams) (*data.User, error) {
	var user *data.User
	err := mapstructure.WeakDecode(p.Args["data"], &user)
	if err != nil {
		return nil, err
	}
	err = my.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "nickname"}),
	}).Create(&user).Error
	return user, err
}
