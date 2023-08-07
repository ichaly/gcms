package auth

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"net/http"
)

func NewOauthServer(ts oauth2.TokenStore, cs oauth2.ClientStore) *server.Server {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(ts, nil)
	manager.MapClientStorage(cs)

	s := server.NewDefaultServer(manager)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(server.ClientFormHandler)
	//s.SetResponseTokenHandler(func(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {
	//	code := 200
	//	if len(statusCode) > 0 {
	//		code = statusCode[0]
	//	}
	//	if code == 200 {
	//		return render.JSON(w, core.OK.WithData(data))
	//	} else {
	//		var err error
	//		if v, e := data["error"]; e {
	//			err = errors.New(fmt.Sprintf("%v", v))
	//		}
	//		return render.JSON(w, core.NewResult(code).WithError(err))
	//	}
	//})
	s.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		return errors.NewResponse(err, http.StatusInternalServerError)
	})

	return s
}
