package categories

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	cmsErr "github.com/iamsabbiralam/restora/client/error"
	"github.com/iamsabbiralam/restora/client/handler/common"
	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
)

func (s *Svc) editCategoryHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.category.editCategoryHandler")
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	res, err := s.Category.GetCategory(r.Context(), &catG.GetCategoryRequest{
		ID: id,
	})
	if err != nil {
		errMsg := "error to get category by ID"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	if res == nil {
		errMsg := "response is nil"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	data := CategoryTempData{
		CSRFField: csrf.TemplateField(r),
		Form: Category{
			ID:     res.ID,
			Name:   res.Name,
			Status: res.Status,
		},
		FormAction: common.DynamicUrlSwitch(common.CategoryEditPath, map[string]string{"id": id}),
		Status:     common.GetStatus(catG.Status_name),
	}

	s.loadCategoryTemplate(w, r, data, "edit-category.html")
}

func (s *Svc) updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.employee.updateEmployeeHandler")
	params := mux.Vars(r)
	id := params["id"]
	if err := r.ParseForm(); err != nil {
		errMsg := "parsing form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var form Category
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "error with decode form"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	errMessage := form.ValidateCategory(s.Server, r, form.ID)
	if errMessage != nil {
		s.validateMsg(w, r, errMessage, form, nil, "edit-category.html")
		return
	}

	_, err := s.Category.UpdateCategory(r.Context(), &catG.UpdateCategoryRequest{
		ID:        id,
		Name:      form.Name,
		Status:    form.Status,
		UpdatedBy: s.GetSessionUser(r).ID,
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validateMsg(w, r, hErr, form, hErr.ValidationErrors, "edit-category.html")
			return
		}

		errMsg := "error with update category"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.CategoryListPath, http.StatusSeeOther)
}
