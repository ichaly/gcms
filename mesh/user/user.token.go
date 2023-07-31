package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/base"
)

type token struct {
}

func NewUserToken() base.Schema {
	return &token{}
}

func (*token) Name() string {
	return "token"
}

func (*token) Description() string {
	return "令牌"
}

func (*token) Host() interface{} {
	return User
}

func (*token) Type() interface{} {
	return ""
}

func (my *token) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
