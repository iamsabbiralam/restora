package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"github.com/yookoala/realpath"

	"github.com/iamsabbiralam/restora/client/conn"
	"github.com/iamsabbiralam/restora/client/handler"
	"github.com/iamsabbiralam/restora/utility"
	"github.com/iamsabbiralam/restora/utility/logging"
)

const (
	svcName = "client"
	version = "1.0.0"
)

func main() {
	cfg, err := utility.NewConfig("env/config")
	if err != nil {
		log.Fatal(err)
	}

	env := cfg.GetString("runtime.environment")
	logger := logging.NewLogger(cfg).WithFields(logrus.Fields{
		"environment": env,
		"service":     svcName,
		"version":     version,
	})
	conns := conn.NewConns(logger, cfg)
	defer conns.Close()
	server, err := setupServer(logger, cfg, conns)
	if err != nil {
		logger.WithError(err).Error("failed to run setup server")
	}

	server.Use(func(h http.Handler) http.Handler {
		recov := negroni.NewRecovery()
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recov.ServeHTTP(w, r, h.ServeHTTP)
		})
	})

	l, err := net.Listen("tcp", ":"+cfg.GetString("server.port"))
	if err != nil {
		logger.WithError(err).Error("failed to listen server")
	}

	log.Printf("starting %s service on %s", svcName, l.Addr())
	if err := http.Serve(l, server); err != nil {
		logger.WithError(err).Error("failed to serve server")
	}
}

func setupServer(logger *logrus.Entry, cfg *viper.Viper, conn *conn.Conn) (*mux.Router, error) {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	cookie := sessions.NewCookieStore([]byte(cfg.GetString("session.key")))
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	assetPath, err := realpath.Realpath(filepath.Join(wd, "assets"))
	if err != nil {
		return nil, err
	}

	asst := afero.NewIOFS(afero.NewBasePathFs(afero.NewOsFs(), assetPath))
	srv, err := handler.NewServer(cfg, logger, asst, decoder, nil, cookie, conn)
	if err != nil {
		return nil, err
	}

	return srv, nil
}
