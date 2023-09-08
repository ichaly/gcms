package auth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ichaly/gcms/apps/data"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	JWTKey = "openapi_oauth2_token_secret"
)

func NewOauthServer(db *gorm.DB, ts oauth2.TokenStore, cs oauth2.ClientStore) *server.Server {
	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(ts)
	manager.MapClientStorage(cs)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate(JWTKey, []byte("pibigstar"), jwt.SigningMethodHS512))

	s := server.NewDefaultServer(manager)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(ClientInfoHandler(cs))
	s.SetPasswordAuthorizationHandler(PasswordAuthorizationHandler(db))
	s.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		return errors.NewResponse(err, http.StatusInternalServerError)
	})

	return s
}

func PasswordAuthorizationHandler(db *gorm.DB) func(context.Context, string, string, string) (string, error) {
	return func(ctx context.Context, clientID, username, password string) (string, error) {
		user := data.User{}
		err := db.Model(&user).Where("username = ?", username).First(&user).Error
		if err != nil {
			return "", err
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return "", err
		}
		return strconv.FormatUint(uint64(user.ID), 10), nil
	}
}

func ClientInfoHandler(cs oauth2.ClientStore) func(*http.Request) (string, string, error) {
	return func(r *http.Request) (string, string, error) {
		clientID, clientSecret, ok := r.BasicAuth()
		if !ok {
			clientID = r.Form.Get("client_id")
			clientSecret = r.Form.Get("client_secret")
		}
		if clientID == "" {
			return "", "", errors.ErrInvalidClient
		}

		client, err := cs.GetByID(r.Context(), clientID)
		if err != nil {
			return "", "", err
		}

		// single-page auth flow, per https://www.oauth.com/oauth2-servers/single-page-apps/
		return client.GetSecret(), clientSecret, nil
	}
}
