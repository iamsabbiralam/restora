package auth

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"

	cmsErr "github.com/iamsabbiralam/restora/client/error"
	"github.com/iamsabbiralam/restora/client/handler/common"
	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
)

type Registration struct {
	UserName        string
	Email           string
	Password        string
	ConfirmPassword string
}

type RegistrationTempData struct {
	CSRFField   template.HTML
	Form        Registration
	FormAction  string
	GlobalURLs  map[string]string
	FormErrors  map[string]string
	FormMessage map[string]string
}

func (s *Svc) getRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "handler.auth.getRegistrationHandler")
	data := RegistrationTempData{
		CSRFField:  csrf.TemplateField(r),
		FormAction: common.RegistrationPath,
	}

	s.loadRegistrationTemplate(w, r, data, "register.html")
}

func (s *Svc) postRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.auth.postRegistrationHandler")
	ctx := r.Context()
	if err := r.ParseForm(); err != nil {
		errMsg := "error with parse form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var form Registration
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "error with decode form"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	errMessage := form.ValidateRegistration(s, r, "")
	if errMessage != nil {
		s.validateRegistrationMsg(w, r, errMessage, form, nil, "register.html")
		return
	}

	pass, err := common.HashPassword(form.ConfirmPassword)
	if err != nil {
		errMsg := "error with hash password"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	_, err = s.User.CreateUser(ctx, &userG.CreateUserRequest{
		UserName: form.UserName,
		Email:    form.Email,
		Password: pass,
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validateRegistrationMsg(w, r, hErr, form, hErr.ValidationErrors, "register.html")
			return
		}

		errMsg := "failed to register user"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.LoginPath, http.StatusSeeOther)
}
