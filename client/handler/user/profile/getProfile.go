package profile

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
	"google.golang.org/protobuf/types/known/timestamppb"

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

func (s *Svc) uploadProfileImageHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.Profile.UserProfileEditHandler")
	ctx := r.Context()
	loggedUser := s.GetSessionUser(r).ID
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		errMsg := "parsing form"
		log.WithError(err).Error(errMsg)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var form Profile
	if err := s.Decoder.Decode(&form, r.PostForm); err != nil {
		log.WithError(err).Error("decoding form")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	res, err := s.User.GetUserByID(ctx, &userG.GetUserByIDRequest{
		ID: loggedUser,
	})
	if err != nil {
		errMsg := "unable to get profile info"
		log.WithError(err).Error(errMsg)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if res == nil {
		errMsg := "user not found"
		log.WithError(err).Error(errMsg)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if err := validation.Validate(validateImage(s, r, loggedUser)); err != nil {
		errMsg := "image is required"
		log.WithError(err).Error(errMsg)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	file, fileHeader, err := r.FormFile("Image")
	if err != nil {
		errMsg := "unable to get file"
		log.WithError(err).Error(errMsg)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	image, err := s.SaveImage(file, fileHeader, "assets/images/profile/")
	if err != nil {
		errMsg := "unable to save file"
		log.WithError(err).Error(errMsg)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	form.Image = image
	_, err = s.User.UpdateUser(ctx, &userG.UpdateUserRequest{
		Image: image,
	})
	if err != nil {
		errMsg := "unable to update profile image"
		log.WithError(err).Error(errMsg)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	dataImage := map[string]string{
		"Image":   image,
		"Message": "Successfully update Image",
	}
	if err := s.SessionResetLoginData(r, w); err != nil {
		log.Error("error with update session user: %+v", err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(dataImage)
}
