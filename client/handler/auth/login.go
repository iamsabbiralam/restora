package auth

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	
	"github.com/iamsabbiralam/restora/client/handler/common"
)

type Login struct {
	Email    string
	Password string
}

type LoginTempData struct {
	CSRFField   template.HTML
	Form        Login
	FormAction  string
	GlobalURLs  map[string]string
	FormErrors  map[string]string
	FormMessage map[string]string
}

func (s *Svc) loadLoginForm(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "handler.auth.loadLoginForm")
	data := LoginTempData{
		CSRFField:   csrf.TemplateField(r),
		FormAction:  common.LoginInPath,
	}

	s.loadLoginTemplate(w, r, data, "login.html")
}

func (s *Svc) loadLoginTemplate(w http.ResponseWriter, r *http.Request, data LoginTempData, htmlFile string) {
	log := s.Logger.WithField("method", "handler.auth.loadLoginTemplate")
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