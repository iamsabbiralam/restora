package auth

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"

	"github.com/iamsabbiralam/restora/client/handler/common"
)

type Registration struct {
	Email    string
	Password string
}

type RegistrationTempData struct {
	CSRFField   template.HTML
	Form        Registration
	FormAction  string
	GlobalURLs  map[string]string
	FormErrors  map[string]string
	FormMessage map[string]string
}

func (s *Svc) loadRegistrationForm(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "handler.auth.loadRegistrationForm")
	data := RegistrationTempData{
		CSRFField:  csrf.TemplateField(r),
		FormAction: common.RegistrationPathPath,
	}

	s.loadRegistrationTemplate(w, r, data, "register.html")
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
