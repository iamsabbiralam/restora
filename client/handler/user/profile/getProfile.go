package profile

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/iamsabbiralam/restora/client/handler/common"
	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		CSRFField: csrf.TemplateField(r),
		Form: Profile{
			FirstName: res.FirstName,
			LastName:  res.LastName,
			Email:     res.Email,
			UserName:  res.UserName,
			Mobile:    res.PhoneNumber,
			Gender:    int(res.Gender),
			DOB:       res.Birthday.AsTime().Format("2006-01-02"),
			Address:   res.Address,
			City:      res.City,
			Country:   res.Country,
		},
	}

	s.loadProfileTemplate(w, r, data, "profile.html")
}

func (s *Svc) updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.Profile.updateProfileHandler")
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
		ID:          loggedUser,
		FirstName:   form.FirstName,
		LastName:    form.LastName,
		PhoneNumber: form.Mobile,
		UserName:    form.UserName,
		Email:       form.Email,
		Gender:      int64(form.Gender),
		Birthday:    timestamppb.New(s.StringToDate(form.DOB)),
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
