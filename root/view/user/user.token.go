package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/root/data"
	"time"
)

const sign_key = "hello jwt"

type token struct{}

type myCustomClaims struct {
	UserID     uint64
	Username   string
	GrantScope string
	jwt.RegisteredClaims
}

func NewUserToken() core.Schema {
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
	user := p.Source.(*data.User)
	claim := myCustomClaims{
		UserID:     uint64(user.ID),
		Username:   user.Username,
		GrantScope: "read_user_info",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   //签发者
			Subject:   "Tom",                                           //签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(sign_key))
	return token, err
}

func parseTokenHs256(tokenString string) (*myCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &myCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sign_key), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*myCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
