package dashboard

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
)

type Svc struct {
	*common.Server
}

type DashboardTempData struct {
	CSRFField   template.HTML
	FormAction  string
	GlobalURLs  map[string]string
	FormErrors  map[string]string
	FormMessage map[string]string
}

func Register(h *common.Server, m *mux.Router) (*mux.Router, error) {
	s := &Svc{
		Server: h,
	}

	m.HandleFunc(common.DashboardPath, s.getDashboardHandler).Methods("GET").Name("dashboard")
	return m, nil
}

func (s *Svc) loadDashboardTemplate(w http.ResponseWriter, r *http.Request, data DashboardTempData, htmlFile string) {
	log := s.Logger.WithField("method", "handler.admin.dashboard.loadDashboardTemplate")
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

func (s *Svc) getDashboardHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "handler.admin.dashboard.getDashboardHandler")
	data := DashboardTempData{
		CSRFField:  csrf.TemplateField(r),
		FormAction: common.RegistrationPath,
	}

	s.loadDashboardTemplate(w, r, data, "dashboard.html")
}
