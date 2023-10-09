package categories

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
)

func (s *Svc) deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.category.deleteCategoryHandler")
	if err := r.ParseForm(); err != nil {
		errMsg := "error with parsing form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	_, err := s.Category.DeleteCategory(r.Context(), &catG.DeleteCategoryRequest{
		ID:        id,
		DeletedBy: s.GetSessionUser(r).ID,
	})
	if err != nil {
		errMsg := "error with Delete category"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.CategoryListPath, http.StatusSeeOther)
}
