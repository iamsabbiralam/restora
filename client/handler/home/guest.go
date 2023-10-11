package guest

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
	r.HandleFunc(common.HomePath, s.GetHomeHandler).Methods("GET").Name("home")

	return r, nil
}
