package brands

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
)

func (s *Svc) deleteBrandHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.brand.deleteBrandHandler")
	if err := r.ParseForm(); err != nil {
		errMsg := "error with parsing form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	_, err := s.Brand.DeleteBrand(r.Context(), &braG.DeleteBrandRequest{
		ID:        id,
		DeletedBy: s.GetSessionUser(r).ID,
	})
	if err != nil {
		errMsg := "error with delete brand"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.BrandListPath, http.StatusSeeOther)
}
