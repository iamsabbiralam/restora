package password

import (
	"html/template"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
)

type Svc struct {
	*common.Server
}

type ChangePasswordForm struct {
	CurrentPassword string
	NewPassword     string
	ConfirmPassword string
}

type ChangePasswordTempData struct {
	CSRFField   template.HTML
	Form        ChangePasswordForm
	FormAction  string
	FormErrors  map[string]string
	FormMessage map[string]string
}

func Register(h *common.Server, mw *mux.Router) (*mux.Router, error) {
	s := &Svc{
		Server: h,
	}

	// mw.HandleFunc(common.ChangePasswordPath, s.getChangePasswordHandler).Methods("GET").Name("change.password")
	mw.HandleFunc(common.ChangePasswordEditPath, s.postChangePasswordHandler).Methods("POST").Name("change.password.update")
	return mw, nil
}

func (s *Svc) loadChangePasswordTemplate(w http.ResponseWriter, r *http.Request, data ChangePasswordTempData, html string) {
	log := s.Logger.WithField("method", "handler.password.loadChangePasswordTemplate")
	tmpl := s.LookupTemplate(html)
	if tmpl == nil {
		log.Error("unable to load template")
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Infof("error with template execution: %+v", err)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}
}

func (cp ChangePasswordForm) ValidatePassword(server *common.Server, r *http.Request, id string) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&cp,
		validation.Field(&cp.CurrentPassword, vre("The current password field is required"), validation.Length(8, 100)),
		validation.Field(&cp.NewPassword, vre("The new password field is required"), validation.Length(8, 100)),
		validation.Field(&cp.ConfirmPassword, vre("The confirm password field is required")),
	)
}

func (s *Svc) validatePasswordMsg(w http.ResponseWriter, r *http.Request, err error, form ChangePasswordForm, errEmp map[string]string, tmpl string) error {
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

	data := ChangePasswordTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       form,
		FormErrors: vErrs,
	}

	s.loadChangePasswordTemplate(w, r, data, tmpl)
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
