package apis

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

type UserApi struct {
	db *gorm.DB
}

func NewUserApi(d *gorm.DB, e *boot.Engine) core.Schema {
	my := &UserApi{d}
	e.NewQuery(my.GetUsers).Name("users").Description("用户管理")
	return my
}

func (my *UserApi) GetUsers(p graphql.ResolveParams) ([]*data.User, error) {
	var res []*data.User
	err := my.db.Model(&data.User{}).Find(&res).Error
	return res, err
}
