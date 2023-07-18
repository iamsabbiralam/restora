package auth

import (
	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
)

type Svc struct {
	*common.Server
}

func Register(h *common.Server, r *mux.Router) (*mux.Router, error) {
	s := &Svc{
		Server: h,
	}

	r.HandleFunc(common.LoginInPath, s.loadLoginForm).Methods("GET").Name("loginUrl")
	r.HandleFunc(common.RegistrationPathPath, s.loadRegistrationForm).Methods("GET").Name("registrationUrl")
	// r.HandleFunc(common.LogoutPath, s.handleLogout).Methods("GET").Name("logout")
	return r, nil
}
