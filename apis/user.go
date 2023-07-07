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
	_ = e.AddTo(my.GetUsers, boot.Query, "", "")
	return my
}

func (my *UserApi) GetUsers(p graphql.ResolveParams) ([]*data.User, error) {
	var res []*data.User
	err := my.db.Model(&data.User{}).Find(&res).Error
	return res, err
}
