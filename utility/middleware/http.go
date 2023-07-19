package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func ChainHTTPMiddleware(usr *mux.Router, logger *logrus.Entry, mw ...func(http.Handler) http.Handler) {
	for _, f := range []func(http.Handler) http.Handler{
		Logger(logger),
		Gzip,
		ContentType(logger, "text/html"),
	} {
		usr.Use(f)
	}
	for _, f := range mw {
		usr.Use(f)
	}
}

func Logger(logger *logrus.Entry) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			lrw := negroni.NewResponseWriter(w)
			h.ServeHTTP(lrw, req.WithContext(req.Context()))

			// TODO: after a while log just errors, for now log everything
			status := lrw.Status()
			ss := lrw.Header()
			hostname := strings.ToLower(req.Host)
			method := req.Method
			path := req.URL.String()
			location := ss.Get("Location")
			ip := GetIP(req)

			log := logger.WithFields(logrus.Fields{
				"status":     status,
				"statusText": http.StatusText(status),
				"hostname":   hostname,
				"method":     method,
				"path":       path,
				"clientIP":   ip,
			})
			if location != "" {
				log = log.WithField("redirectLocation", location)
			}

			log.Infof("request handled")
		})
	}
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-forwarded-for")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func CSRF(logger *logrus.Entry, secret []byte, opts ...csrf.Option) func(h http.Handler) http.Handler {
	opts = append([]csrf.Option{
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logger.WithContext(r.Context())
			log.WithFields(logrus.Fields{
				"csrf_error": csrf.FailureReason(r).Error(),
				"token":      csrf.Token(r),
				"template":   csrf.TemplateField(r),
			}).Error("csrf error")
			fmt.Fprintln(w, csrf.FailureReason(r))
		})),
	}, opts...)
	return csrf.Protect(secret, opts...)
}

type contentWriter struct {
	http.ResponseWriter
	def    string
	logger *logrus.Entry
}

func ContentType(logger *logrus.Entry, typeDefault string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(contentWriter{
				ResponseWriter: w,
				def:            typeDefault,
				logger:         logger.WithContext(r.Context()),
			}, r)
		})
	}
}
