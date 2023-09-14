package user

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	upb "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"github.com/iamsabbiralam/restora/usermgm/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (h *Handler) UpdateUser(ctx context.Context, req *upb.UpdateUserRequest) (*upb.UpdateUserResponse, error) {
	log := h.logger.WithField("method", "Service.User.UpdateUser")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	if err := h.ValidateRequestedData(ctx, storage.User{
		ID:       req.GetID(),
		Username: req.GetUserName(),
		Email:    req.GetEmail(),
	}, req.GetID()); err != nil {
		log.WithError(err).Error("validation error while updating user")
		return nil, uErr.HandleServiceErr(err)
	}

	dbUser := storage.User{
		ID:       req.GetID(),
		Status:   int32(req.GetStatus()),
		Password: req.Password,
		Email:    req.GetEmail(),
		Username: req.GetUserName(),
	}
	res, err := h.usr.UpdateUser(ctx, dbUser)
	if err != nil {
		errMsg := "failed to update user"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if res == nil {
		return nil, uErr.HandleServiceErr(errors.New("update user response is nil"))
	}

	dbUserInfo := storage.UserInformation{
		UserID:    req.GetID(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Mobile:    req.GetPhoneNumber(),
		DOB:       req.GetBirthday().AsTime(),
		Gender:    int(req.GetGender()),
		Address:   req.GetAddress(),
		City:      req.GetCity(),
		Country:   req.GetCountry(),
	}
	userInfo, err := h.usr.UpdateUserInformationByUserID(ctx, dbUserInfo)
	if err != nil {
		errMsg := "failed to update user information by user id"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if userInfo == nil {
		return nil, uErr.HandleServiceErr(errors.New("update user information response is nil"))
	}

	var password string
	if req.GetPassword() == "" {
		password = res.Password
	} else {
		password = req.GetPassword()
	}

	return &upb.UpdateUserResponse{
		ID:        res.ID,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Email:     res.Email,
		UserName:  res.Username,
		Password:  password,
		Status:    upb.Status(res.Status),
		Address:   userInfo.Address,
		City:      userInfo.City,
		Country:   userInfo.Country,
		UpdatedAt: timestamppb.New(res.UpdatedAt),
		UpdatedBy: res.UpdatedBy.String,
	}, nil
}
