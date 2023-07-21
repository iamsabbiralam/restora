package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/iamsabbiralam/restora/client/handler/common"
	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
)

type Svc struct {
	*common.Server
}

func Register(h *common.Server, r *mux.Router) (*mux.Router, error) {
	s := &Svc{
		Server: h,
	}

	r.HandleFunc(common.RegistrationPath, s.getRegistrationHandler).Methods("GET").Name("register.get")
	r.HandleFunc(common.RegistrationPath, s.postRegistrationHandler).Methods("POST").Name("register.store")
	r.HandleFunc(common.LoginInPath, s.loadLoginForm).Methods("GET").Name("login.get")
	r.HandleFunc(common.LoginInPath, s.postLoginForm).Methods("POST").Name("login.post")
	// r.HandleFunc(common.LogoutPath, s.handleLogout).Methods("GET").Name("logout")
	return r, nil
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

func (s *Svc) validateRegistrationMsg(w http.ResponseWriter, r *http.Request, err error, regForm Registration, errEmp map[string]string, temp string) error {
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
		Form:       regForm,
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

func (l Login) ValidateLoginForm(s *Svc, r *http.Request) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, vre("The email field is required"), validation.Length(3, 100)),
		validation.Field(&l.Password, vre("The password field is required"), validatePassword(s, r, l.Email, l.Password)),
	)
}

func (s *Svc) validationLoginMsg(w http.ResponseWriter, r *http.Request, err error, loginForm Login, errEmp map[string]string, temp string) {
	s.Logger.WithField("method", "handler.auth.validationLogin")
	if err := loginForm.ValidateLoginForm(s, r); err != nil {
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

		data := LoginTempData{
			CSRFField:  csrf.TemplateField(r),
			FormErrors: vErrs,
			FormAction: common.LoginInPath,
			Form:       loginForm,
		}

		s.loadLoginTemplate(w, r, data, temp)
		return
	}
}

func validatePassword(s *Svc, r *http.Request, email, pass string) validation.Rule {
	return validation.By(func(interface{}) error {
		s.Logger.WithField("method", "handler.auth.validatePassword")
		res, _ := s.User.GetUserByEmail(r.Context(), &userG.GetUserByEmailRequest{
			Email: email,
		})

		if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(pass)); err != nil {
			return errors.New("invalid password given")
		}

		return nil
	})
}

func (s *Svc) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.Cookies.Get(r, common.SessionCookieName)
		if err != nil {
			log.Fatal(err)
		}

		authUserID := session.Values["authUserID"]
		if authUserID != nil {
			next.ServeHTTP(w, r)
		}

		http.Redirect(w, r, common.LoginInPath, http.StatusTemporaryRedirect)

	})
}

func (s *Svc) loginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.Cookies.Get(r, common.SessionCookieName)
		if err != nil {
			log.Fatal(err)
		}

		authUserID := session.Values["authUserID"]
		if authUserID != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}
