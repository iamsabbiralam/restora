package brands

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	cmsErr "github.com/iamsabbiralam/restora/client/error"
	"github.com/iamsabbiralam/restora/client/handler/common"
	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
)

func (s *Svc) editBrandHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.brand.editBrandHandler")
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	res, err := s.Brand.GetBrand(r.Context(), &braG.GetBrandRequest{
		ID: id,
	})
	if err != nil {
		errMsg := "error to get brand by ID"
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

	data := BrandTempData{
		CSRFField: csrf.TemplateField(r),
		Form: Brand{
			ID:     res.ID,
			Name:   res.Name,
			Status: res.Status,
		},
		FormAction: common.DynamicUrlSwitch(common.BrandEditPath, map[string]string{"id": id}),
		Status:     common.GetStatus(braG.Status_name),
	}

	s.loadBrandTemplate(w, r, data, "edit-brand.html")
}

func (s *Svc) updateBrandHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.brand.updateBrandHandler")
	params := mux.Vars(r)
	id := params["id"]
	if err := r.ParseForm(); err != nil {
		errMsg := "parsing form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var form Brand
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "error with decode form"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	errMessage := form.ValidateBrand(s.Server, r, form.ID)
	if errMessage != nil {
		s.validateMsg(w, r, errMessage, form, nil, "edit-brand.html")
		return
	}

	_, err := s.Brand.UpdateBrand(r.Context(), &braG.UpdateBrandRequest{
		ID:        id,
		Name:      form.Name,
		Status:    form.Status,
		UpdatedBy: s.GetSessionUser(r).ID,
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validateMsg(w, r, hErr, form, hErr.ValidationErrors, "edit-brand.html")
			return
		}

		errMsg := "error with update brand"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.BrandListPath, http.StatusSeeOther)
}
