package brands

import (
	"net/http"

	"github.com/gorilla/csrf"

	cmsErr "github.com/iamsabbiralam/restora/client/error"
	"github.com/iamsabbiralam/restora/client/handler/common"
	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
)

func (s *Svc) createBrandHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "handler.admin.brand.createBrandHandler")
	data := BrandTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       Brand{},
		FormAction: common.BrandCreatePath,
	}

	s.loadBrandTemplate(w, r, data, "create-brand.html")
}

func (s *Svc) storeBrandHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.brand.storeBrandHandler")
	if err := r.ParseForm(); err != nil {
		errMsg := "parsing form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var form Brand
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "unable to decode form data"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	errMessage := form.ValidateBrand(s.Server, r, "")
	if errMessage != nil {
		s.validateMsg(w, r, errMessage, form, nil, "create-brand.html")
		return
	}

	_, err := s.Brand.CreateBrand(r.Context(), &braG.CreateBrandRequest{
		Name:      form.Name,
		Status:    form.Status,
		CreatedBy: s.GetSessionUser(r).ID,
		UpdatedBy: s.GetSessionUser(r).ID,
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validateMsg(w, r, hErr, form, hErr.ValidationErrors, "create-brand.html")
			return
		}

		errMsg := "unable to create the brand"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.BrandListPath, http.StatusSeeOther)
}
