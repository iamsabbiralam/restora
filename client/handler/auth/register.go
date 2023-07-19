package auth

import (
	"html/template"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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

	http.Redirect(w, r, common.LoginInPath, http.StatusSeeOther)
}

func (rg Registration) ValidateRegistration(server *Svc, r *http.Request, id string) error {
	server.Logger.WithField("method", "handler.register.ValidateRegistration")
	vre := validation.Required.Error
	return validation.ValidateStruct(&rg,
		validation.Field(&rg.UserName, vre("The Username is required")),
		validation.Field(&rg.Email, vre("The email is required"), is.EmailFormat.Error("The email is not valid")),
		validation.Field(&rg.Password, vre("The password is required")),
		validation.Field(&rg.ConfirmPassword, vre("The confirm password is required")),
	)
}

func (s *Svc) loadRegistrationTemplate(w http.ResponseWriter, r *http.Request, data RegistrationTempData, htmlFile string) {
	log := s.Logger.WithField("method", "handler.auth.loadRegistrationTemplate")
	tmpl := s.LookupTemplate(htmlFile)
	if tmpl == nil {
		log.Error("unable to load template")
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Errorf("unable to execute template: %s", err)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}
}

func (s *Svc) validateRegistrationMsg(w http.ResponseWriter, r *http.Request, err error, form Registration, errEmp map[string]string, temp string) error {
	s.Logger.WithField("method", "handler.register.validateRegistrationMsg")
	vErrs := map[string]string{}
	if e, ok := err.(validation.Errors); ok {
		if len(e) > 0 {
			for key, value := range e {
				vErrs[key] = value.Error()
			}
		}
	}

	if errEmp != nil {
		vErrs = errEmp
	}

	data := RegistrationTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       form,
		FormErrors: vErrs,
		FormAction: common.RegistrationPath,
	}

	s.loadRegistrationTemplate(w, r, data, temp)
	return nil
}

func getVErrs(err string) map[string]string {
	vErrs := map[string]string{}
	for _, v := range strings.Split(err, "; ") {
		val := strings.Split(v, ": ")
		vErrs[strings.Title(val[0])] = val[1]
	}

	return vErrs
}
