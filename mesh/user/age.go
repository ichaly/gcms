package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"time"
)

type Age struct {
}

func NewAge() core.Schema {
	return &Age{}
}

func (*Age) Name() string {
	return "age"
}

func (*Age) Host() interface{} {
	return User
}

func (*Age) Description() string {
	return "年龄"
}

func (my *Age) Resolve(p graphql.ResolveParams) (int, error) {
	user := p.Source.(*data.User)
	if user.Birthday.IsZero() {
		return 0, nil
	}
	return time.Now().Year() - user.Birthday.Year(), nil
}
