package user

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	upb "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (h *Handler) GetUserByEmail(ctx context.Context, req *upb.GetUserByEmailRequest) (*upb.GetUserByEmailResponse, error) {
	log := h.logger.WithField("method", "service.user.GetUserByEmail")
	if req.GetEmail() == "" {
		return nil, uErr.HandleServiceErr(errors.New("Email is required"))
	}

	r, err := h.usr.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		errMsg := "failed to get user by email"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if r == nil {
		log.WithError(err).Error("error while getting user by email")
		return nil, uErr.HandleServiceErr(err)
	}

	resp := &upb.GetUserByEmailResponse{
		ID:        r.ID,
		Email:     r.Email,
		UserName:  r.Username,
		Password:  r.Password,
		Status:    upb.Status(r.Status),
		CreatedAt: timestamppb.New(r.CreatedAt),
		CreatedBy: r.CreatedBy,
		UpdatedAt: timestamppb.New(r.UpdatedAt),
		UpdatedBy: r.UpdatedBy.String,
	}

	return resp, nil
}

func (h *Handler) GetUserByID(ctx context.Context, req *upb.GetUserByIDRequest) (*upb.GetUserByIDResponse, error) {
	log := h.logger.WithField("method", "service.user.GetUserByID")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is required"))
	}

	r, err := h.usr.GetUserByID(ctx, req.GetID())
	if err != nil {
		errMsg := "failed to get user by id"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if r == nil {
		return nil, uErr.HandleServiceErr(errors.New("error with getting user by id"))
	}

	userInfo, err := h.usr.GetUserInformationByUserID(ctx, r.ID)
	if err != nil {
		errMsg := "failed to get user information by user id"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if userInfo == nil {
		return nil, uErr.HandleServiceErr(errors.New("error with getting user information by user id"))
	}

	resp := &upb.GetUserByIDResponse{
		ID:          r.ID,
		FirstName:   userInfo.FirstName,
		LastName:    userInfo.LastName,
		Email:       r.Email,
		UserName:    r.Username,
		Password:    r.Password,
		PhoneNumber: userInfo.Mobile,
		Gender:      int64(userInfo.Gender),
		Address:     userInfo.Address,
		City:        userInfo.City,
		Country:     userInfo.Country,
		Status:      upb.Status(r.Status),
		CreatedAt:   timestamppb.New(r.CreatedAt),
		CreatedBy:   r.CreatedBy,
		UpdatedAt:   timestamppb.New(r.UpdatedAt),
		UpdatedBy:   r.UpdatedBy.String,
	}

	return resp, nil
}

func (h *Handler) GetUserByUsername(ctx context.Context, req *upb.GetUserByUsernameRequest) (*upb.GetUserByUsernameResponse, error) {
	log := h.logger.WithField("method", "service.user.GetUserByUsername")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is required"))
	}

	r, err := h.usr.GetUserByUsername(ctx, req.GetUserName())
	if err != nil {
		errMsg := "failed to get user by username"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if r == nil {
		log.WithError(err).Error("error while getting username")
		return nil, uErr.HandleServiceErr(err)
	}

	resp := &upb.GetUserByUsernameResponse{
		ID:        r.ID,
		Email:     r.Email,
		UserName:  r.Username,
		Password:  r.Password,
		Status:    upb.Status(r.Status),
		CreatedAt: timestamppb.New(r.CreatedAt),
		CreatedBy: r.CreatedBy,
		UpdatedAt: timestamppb.New(r.UpdatedAt),
		UpdatedBy: r.UpdatedBy.String,
	}

	return resp, nil
}
