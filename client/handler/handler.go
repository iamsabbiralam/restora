package handler

import (
	"errors"
	"io/fs"
	"net/http"

	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/iamsabbiralam/restora/client/conn"
	"github.com/iamsabbiralam/restora/client/handler/common"
	guest "github.com/iamsabbiralam/restora/client/handler/home"
	urlAuth "github.com/iamsabbiralam/restora/client/handler/auth"
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
		Config:       config,
		Logger:       logger,
		Assets:       hashfs.NewFS(assets),
		Decoder:      decoder,
		Cookies:      cookies,
	}

	if err := cs.ParseTemplates(); err != nil {
		return nil, err
	}

	// csrfSecure := config.GetBool("csrf.secure")
	csrfSecret := config.GetString("csrf.secret")
	if csrfSecret == "" {
		return nil, errors.New("CSRF secret must not be empty")
	}
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", common.CacheStaticFiles(http.FileServer(http.FS(cs.Assets)))))

	r, err := guest.Register(cs, r)
	if err != nil {
		return nil, err
	}

	r, err = urlAuth.Register(cs, r)
	if err != nil {
		return nil, err
	}

	r.NotFoundHandler = cs.GetErrorHandler()
	return r, nil
}
