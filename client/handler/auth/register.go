package auth

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"

	"github.com/iamsabbiralam/restora/client/handler/common"
)

type Registration struct {
	UserName string
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
		s.validateEmployeeMsg(w, r, errMessage, form, nil, "employee-create.html")
		return
	}

	password := common.RandomPassword(8)
	pass, err := common.HashPassword(password)
	if err != nil {
		errMsg := "error with hash password"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}
}

func (rg *Registration) ValidateEmployee(server *Svc, r *http.Request, id string) error {
	server.Logger.WithField("method", "handler.employee.ValidateEmployee")
	vre := validation.Required.Error
	return validation.ValidateStruct(&s,
		validation.Field(&s.Image, validateEmployeeImage(server, r, id)),
		validation.Field(&s.FirstName, vre("The First name is required"), validation.Match(regexp.MustCompile("^[a-zA-Z_ ]*$")).Error("Enter characters only")),
		validation.Field(&s.LastName, vre("The Last name is required"), validation.Match(regexp.MustCompile("^[a-zA-Z_ ]*$")).Error("Enter characters only")),
		validation.Field(&s.Mobile, vre("The Mobile Number is required"), validation.Match(regexp.MustCompile(`((\+|\(|0)?\d{1,3})?((\s|\)|\-))?(\d{10})$`)).Error("The Mobile Number is invalid")),
		validation.Field(&s.Email, vre("The email is required"), is.EmailFormat.Error("The email is not valid")),
		validation.Field(&s.Address, vre("The address is required")),
		validation.Field(&s.DesignationID, vre("The Designation is required")),
		validation.Field(&s.DepartmentID, vre("The Department is required")),
		validation.Field(&s.OfficeID, vre("The office id is required")),
		validation.Field(&s.Status, vre("Please select status"), validation.Min(1).Error("Status is Invalid"), validation.Max(2).Error("Status is Invalid")),
		validation.Field(&s.EndDate, validation.Date("2006-01-02"), validation.When(s.EndDate != "", validateEmployeeEndDate(server, s.StartDate, s.EndDate))),
		validation.Field(&s.Position, validation.When(s.ID != "", vre("The position is required"))),
		validation.Field(&s.RoleIDs, vre("The Role is required")),
		validation.Field(&s.UserName, vre("The Username is required")),
		validation.Field(&s.GuardianMobile, validation.Match(regexp.MustCompile(`((\+|\(|0)?\d{1,3})?((\s|\)|\-))?(\d{10})$`)).Error("The Mobile Number is invalid")),
		validation.Field(&s.FacebookUrl, is.URL.Error("The Facebook URL is invalid")),
		validation.Field(&s.GitHubUrl, is.URL.Error("The Github URL is invalid")),
		validation.Field(&s.TwitterUrl, is.URL.Error("The Twitter URL is invalid")),
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
