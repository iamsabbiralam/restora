package user

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"

	upb "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"github.com/iamsabbiralam/restora/usermgm/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (h *Handler) CreateUser(ctx context.Context, req *upb.CreateUserRequest) (*upb.CreateUserResponse, error) {
	log := h.logger.WithField("method", "Service.User.CreateUser")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	if err := h.ValidateRequestedData(ctx, storage.User{
		Username: req.GetUserName(),
		Email:    req.GetEmail(),
		Password: req.Password,
		Status:   int32(upb.Status_Active),
	}, ""); err != nil {
		log.WithError(err).Error("validation error while creating user")
		return nil, uErr.HandleServiceErr(err)
	}

	dbUser := storage.User{
		Username:  req.UserName,
		Email:     req.Email,
		Status:    int32(upb.Status_Active),
		Password:  req.GetPassword(),
		CreatedAt: req.CreatedAt.AsTime(),
		UpdatedAt: req.UpdatedAt.AsTime(),
	}

	res, err := h.usr.CreateUser(ctx, dbUser)
	if err != nil {
		errMsg := "failed to create user"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &upb.CreateUserResponse{
		ID: res,
	}, nil
}

func validateUserEmail(h *Handler, email string, id string, log *logrus.Entry) validation.Rule {
	return validation.By(func(interface{}) error {
		res, err := h.GetUserByEmail(context.Background(), &upb.GetUserByEmailRequest{
			Email: email,
		})

		if err != nil {
			return errors.New("unable to get user by email")
		}

		if res == nil {
			return errors.New("response is nil")
		}

		if id != "" && res.GetID() == id && res.GetEmail() == email {
			return nil
		}

		if res.GetEmail() == email {
			return errors.New("email already exists")
		}

		return nil
	})
}
