package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/gcms/base"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
)

type Oauth struct {
	Oauth *server.Server
}

func NewOauth(o *server.Server) base.Plugin {
	return &Oauth{Oauth: o}
}

func (my *Oauth) Base() string {
	return "/oauth"
}

func (my *Oauth) Init(r gin.IRouter) {
	//授权路由
	r.Any("/token", my.tokenHandler())
	r.Any("/authorize", my.authorizeHandler())
}

func (my *Oauth) tokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := my.Oauth.HandleTokenRequest(c.Writer, c.Request); err != nil {
			c.JSON(http.StatusInternalServerError, gqlerrors.FormatErrors(err))
		}
	}
}

func (my *Oauth) authorizeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := my.Oauth.HandleAuthorizeRequest(c.Writer, c.Request); err != nil {
			c.JSON(http.StatusInternalServerError, gqlerrors.FormatErrors(err))
		}
	}
}
