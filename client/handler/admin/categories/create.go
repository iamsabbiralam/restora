package categories

import (
	"net/http"

	"github.com/gorilla/csrf"

	cmsErr "github.com/iamsabbiralam/restora/client/error"
	"github.com/iamsabbiralam/restora/client/handler/common"
	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
)

func (s *Svc) createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "handler.admin.category.createCategoryHandler")
	data := CategoryTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       Category{},
		FormAction: common.CategoryCreatePath,
	}

	s.loadCategoryTemplate(w, r, data, "create-category.html")
}

func (s *Svc) storeCategoryHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.category.storeCategoryHandler")
	if err := r.ParseForm(); err != nil {
		errMsg := "parsing form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var form Category
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "unable to decode form data"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	errMessage := form.ValidateCategory(s.Server, r, form.ID)
	if errMessage != nil {
		s.validateMsg(w, r, errMessage, form, nil, "create-category.html")
		return
	}

	_, err := s.Category.CreateCategory(r.Context(), &catG.CreateCategoryRequest{
		Name:      form.Name,
		Status:    form.Status,
		CreatedBy: s.GetSessionUser(r).ID,
		UpdatedBy: s.GetSessionUser(r).ID,
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validateMsg(w, r, hErr, form, hErr.ValidationErrors, "create-category.html")
			return
		}

		errMsg := "unable to update user password"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.CategoryListPath, http.StatusSeeOther)
}
