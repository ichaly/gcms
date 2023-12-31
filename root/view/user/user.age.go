package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/root/data"
	"time"
)

type age struct {
}

func NewUserAge() core.Schema {
	return &age{}
}

func (*age) Name() string {
	return "age"
}

func (*age) Description() string {
	return "年龄"
}

func (*age) Host() interface{} {
	return User
}

func (*age) Type() interface{} {
	return 0
}

func (my *age) Resolve(p graphql.ResolveParams) (interface{}, error) {
	user := p.Source.(*data.User)
	if user.Birthday == nil || user.Birthday.IsZero() {
		return nil, nil
	}
	return time.Now().Year() - user.Birthday.Year(), nil
}
