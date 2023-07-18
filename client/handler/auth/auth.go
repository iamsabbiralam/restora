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

	r.HandleFunc(common.RegistrationPathPath, s.loadRegistrationForm).Methods("GET").Name("register.get")
	r.HandleFunc(common.RegistrationPathPath, s.postRegistrationHandler).Methods("POST").Name("register.store")
	r.HandleFunc(common.LoginInPath, s.loadLoginForm).Methods("GET").Name("login.get")
	// r.HandleFunc(common.LogoutPath, s.handleLogout).Methods("GET").Name("logout")
	return r, nil
}
