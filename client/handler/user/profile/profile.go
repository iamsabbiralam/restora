package profile

import (
	"html/template"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
)

type Svc struct {
	*common.Server
}

type Profile struct {
	UserID    string
	FirstName string
	LastName  string
	Email     string
	Image     string
	Mobile    string
	Gender    int
	DOB       string
	Address   string
	City      string
	Country   string
}

type ProfileTempData struct {
	CSRFField        template.HTML
	Form             Profile
	Data             []Profile
	FormAction       string
	FormErrors       map[string]string
	FormMessage      map[string]string
	PresetPermission map[string]map[string]bool
	URLs             map[string]string
	GlobalURLs       map[string]string
	GlobalURLsUser   map[string]string
	DefaultPair      string
	SearchTerm       string
	Sorting          []string
}

func Register(h *common.Server, mw *mux.Router) (*mux.Router, error) {
	s := &Svc{
		Server: h,
	}

	mw.HandleFunc(common.ProfilePath, s.getProfileHandler).Methods("GET").Name("profile")
	mw.HandleFunc(common.ProfileEditPath, s.updateProfileHandler).Methods("POST").Name("profile.update")
	return mw, nil
}

func (s *Svc) loadProfileTemplate(w http.ResponseWriter, r *http.Request, data ProfileTempData, html string) {
	log := s.Logger.WithField("method", "handler.profile.loadProfileTemplate")
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

func (p Profile) ValidateProfile(server *common.Server, r *http.Request, id string) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&p,
		validation.Field(&p.FirstName, vre("The First name is required"), validation.Length(3, 100), validation.Match(regexp.MustCompile(common.TextValidation)).Error("First name is invalid, only text is allowed.")),
		validation.Field(&p.LastName, vre("The Last name is required"), validation.Length(3, 100), validation.Match(regexp.MustCompile(common.TextValidation)).Error("Last name is invalid, only text is allowed.")),
		validation.Field(&p.Mobile, vre("The Mobile Number is required"), validation.Match(regexp.MustCompile(`((\+|\(|0)?\d{1,3})?((\s|\)|\-))?(\d{10})$`)).Error("The Mobile Number is invalid")),
	)
}

func (s *Svc) validateMsg(w http.ResponseWriter, r *http.Request, err error, form Profile) error {
	vErrs := map[string]string{}
	if e, ok := err.(validation.Errors); ok {
		if len(e) > 0 {
			for key, value := range e {
				vErrs[key] = value.Error()
			}
		}
	}

	data := ProfileTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       form,
		FormErrors: vErrs,
	}

	s.loadProfileTemplate(w, r, data, "profile-edit.html")
	return nil
}
