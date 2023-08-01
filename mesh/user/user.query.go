package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"github.com/ichaly/gcms/util"
	"golang.org/x/crypto/bcrypt"
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
	password := util.EraseMap(p.Args, "where.password")
	res, err := core.QueryResolver[*data.User](p, my.db)
	if err != nil {
		return nil, err
	}
	if password == nil {
		return res, nil
	}
	if users, ok := res.([]*data.User); ok {
		for _, u := range users {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password.(map[string]interface{})["eq"].(string)))
			if err == nil {
				return []*data.User{u}, nil
			}
		}
	}
	return nil, nil
}
