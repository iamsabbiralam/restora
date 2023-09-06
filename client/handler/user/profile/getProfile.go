package profile

import (
	"net/http"

	"github.com/iamsabbiralam/restora/client/handler/common"
	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
)

func (s *Svc) getProfileHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.Profile.getProfileHandler")
	ctx := r.Context()
	loggedUser := s.GetSessionUser(r)
	res, err := s.User.GetUserByID(ctx, &userG.GetUserByIDRequest{
		ID: loggedUser.ID,
	})
	if err != nil {
		errMsg := "unable to get profile data"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	data := ProfileTempData{
		Form: Profile{
			FirstName: res.FirstName,
			LastName:  res.LastName,
			Mobile:    res.PhoneNumber,
			Gender:    int(res.Gender),
			Address:   res.Address,
			City:      res.City,
			Country:   res.Country,
		},
	}

	s.loadProfileTemplate(w, r, data, "profile.html")
}

func (s *Svc) editProfileHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.Profile.editProfileHandler")
	ctx := r.Context()
	loggedUser := s.GetSessionUser(r).ID
	var form Profile
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		errMsg := "unable to decode form data"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	errMessage := form.ValidateProfile(s.Server, r, loggedUser)
	if errMessage != nil {
		s.validateMsg(w, r, errMessage, form)
		return
	}

	_, err := s.User.UpdateUser(ctx, &userG.UpdateUserRequest{
		FirstName:   form.FirstName,
		LastName:    form.LastName,
		PhoneNumber: form.Mobile,
		Gender:      int64(form.Gender),
		Address:     form.Address,
		City:        form.City,
		Country:     form.Country,
	})
	if err != nil {
		errMsg := "unable to update profile data"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	http.Redirect(w, r, common.ProfilePath, http.StatusSeeOther)
}
