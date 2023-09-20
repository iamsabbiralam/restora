package password

import (
	"net/http"

	// "github.com/gorilla/csrf"
	cmsErr "github.com/iamsabbiralam/restora/client/error"

	"github.com/iamsabbiralam/restora/client/handler/common"
	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
)

// func (s *Svc) getChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
// 	s.Logger.WithField("method", "handler.password.getChangePasswordHandler")
// 	data := ChangePasswordTempData{
// 		CSRFField:  csrf.TemplateField(r),
// 		FormAction: common.ChangePasswordEditPath,
// 	}

// 	s.loadChangePasswordTemplate(w, r, data, "profile.html")
// }

func (s *Svc) postChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.password.postChangePasswordHandler")
	loggedUser := s.GetSessionUser(r).ID
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var form ChangePasswordForm
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "unable to decode form data"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	errMessage := form.ValidatePassword(s.Server, r, loggedUser)
	if errMessage != nil {
		s.validatePasswordMsg(w, r, errMessage, form, nil, "edit-password.html")
		return
	}

	res, err := s.User.GetUserByID(r.Context(), &userG.GetUserByIDRequest{
		ID: loggedUser,
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validatePasswordMsg(w, r, hErr, form, hErr.ValidationErrors, "edit-password.html")
			return
		}

		errMsg := "unable to get user by id"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	if res == nil {
		errMsg := "res is empty"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	if form.CurrentPassword != res.Password {
		errMsg := "the given credential is incorrect"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	if form.NewPassword != form.ConfirmPassword {
		errMsg := "password and confirm password does not match"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	pass, err := common.HashPassword(form.ConfirmPassword)
	if err != nil {
		errMsg := "error with hash password"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	_, err = s.User.UpdateUser(r.Context(), &userG.UpdateUserRequest{
		ID:       loggedUser,
		Password: pass,
	})
	if err != nil {
		hErr := cmsErr.ToHTTPError(err)
		if hErr.Code == http.StatusUnprocessableEntity {
			s.validatePasswordMsg(w, r, hErr, form, hErr.ValidationErrors, "edit-password.html")
			return
		}

		errMsg := "unable to update user password"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, common.ChangePasswordPath, http.StatusSeeOther)
}
