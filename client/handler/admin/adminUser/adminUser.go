package adminUser

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
	"github.com/iamsabbiralam/restora/client/handler/paginator"
	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
)

type Svc struct {
	*common.Server
}

type AdminUser struct {
	ID            string
	Username      string
	Email         string
	Status        userG.Status
	RoleID        string
	DepartmentID  string
	DesignationID string
	NewPass       string
	ConPass       string
	Image         string
	RoleIDs       []string
	RoleNames     string
}

type AdminUserTempData struct {
	CSRFField      template.HTML
	Form           AdminUser
	Data           []AdminUser
	FormAction     string
	FormErrors     map[string]string
	FormMessage    map[string]string
	SearchTerm     string
	PaginationData paginator.Paginator
	GlobalURLs     map[string]string
	Status         []common.Status
	PramStatus     string
	SortByColumn   string
	StartDate      string
	EndDate        string
	SortBy         string
	FilterFormErr  string
}

func Register(h *common.Server, mw *mux.Router) (*mux.Router, error) {
	s := &Svc{
		Server: h,
	}

	mw.HandleFunc(common.UserListPath, s.listUserHandler).Methods("GET").Name("user-list")
	return mw, nil
}

func (s *Svc) loadAdminUserTemplate(w http.ResponseWriter, r *http.Request, data AdminUserTempData, htmlFile string) {
	log := s.Logger.WithField("method", "handler.admin.user.loadAdminUserTemplate")
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
