package auth

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"

	cmsErr "github.com/iamsabbiralam/restora/client/error"
	"github.com/iamsabbiralam/restora/client/handler/common"
	"github.com/iamsabbiralam/restora/proto/v1/usermgm/auth"
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
		CSRFField:  csrf.TemplateField(r),
		FormAction: common.LoginInPath,
	}

	s.loadLoginTemplate(w, r, data, "login.html")
}

func (s *Svc) postLoginForm(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.auth.postLoginForm")
	ctx := r.Context()
	if err := r.ParseForm(); err != nil {
		errMsg := "error with parse form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var form Login
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "error with decode form"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	errMessage := form.ValidateLoginForm(s, r)
	if errMessage != nil {
		s.validationLoginMsg(w, r, errMessage, form, nil, "login.html")
		return
	}

	res, err := s.Login.Login(ctx, &auth.LoginRequest{
		Login: &auth.Login{
			Email: form.Email,
		},
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validationLoginMsg(w, r, hErr, form, hErr.ValidationErrors, "login.html")
			return
		}

		errMsg := "failed to login user"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	session, err := s.Cookies.Get(r, common.SessionCookieName)
	if err != nil {
		log.Fatal(err)
	}

	session.Options.HttpOnly = true
	session.Values["authUserID"] = res.ID
	if err := session.Save(r, w); err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, common.DashboardPath, http.StatusSeeOther)
}
