package handler

import (
	"errors"
	"io/fs"
	"net/http"

	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/iamsabbiralam/restora/client/conn"
	dashboard "github.com/iamsabbiralam/restora/client/handler/admin/dashboard"
	urlAuth "github.com/iamsabbiralam/restora/client/handler/auth"
	"github.com/iamsabbiralam/restora/client/handler/common"
	guest "github.com/iamsabbiralam/restora/client/handler/home"
	"github.com/iamsabbiralam/restora/client/handler/user/profile"
	loginG "github.com/iamsabbiralam/restora/proto/v1/usermgm/auth"
	"github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"github.com/iamsabbiralam/restora/utility/middleware"
)

type Svc struct {
	*common.Server
}

func NewServer(
	config *viper.Viper,
	logger *logrus.Entry,
	assets fs.FS,
	decoder *schema.Decoder,
	auth *common.Authenticator,
	cookies *sessions.CookieStore,
	conn *conn.Conn,
) (*mux.Router, error) {
	cs := &common.Server{
		Config:  config,
		Logger:  logger,
		Assets:  hashfs.NewFS(assets),
		Decoder: decoder,
		Cookies: cookies,
		User:    user.NewUserServiceClient(conn.Urm),
		Login:   loginG.NewLoginServiceClient(conn.Urm),
	}

	if err := cs.ParseTemplates(); err != nil {
		return nil, err
	}

	csrfSecure := config.GetBool("csrf.secure")
	csrfSecret := config.GetString("csrf.secret")
	if csrfSecret == "" {
		return nil, errors.New("CSRF secret must not be empty")
	}

	r := mux.NewRouter()
	middleware.ChainHTTPMiddleware(r, logger,
		middleware.CSRF(logger, []byte(csrfSecret), csrf.Secure(csrfSecure), csrf.Path("/")),
	)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", common.CacheStaticFiles(http.FileServer(http.FS(cs.Assets)))))
	r, err := guest.Register(cs, r)
	if err != nil {
		return nil, err
	}	

	mw := r.NewRoute().Subrouter()
	mw.Use(cs.GetAuthMiddleware)
	mw, err = urlAuth.Register(cs, mw)
	if err != nil {
		return nil, err
	}

	mw, err = dashboard.Register(cs, mw)
	if err != nil {
		return nil, err
	}

	mw, err = profile.Register(cs, mw)
	if err != nil {
		return nil, err
	}

	r.NotFoundHandler = cs.GetErrorHandler()
	mw.NotFoundHandler = cs.GetErrorHandler()
	return r, nil
}
