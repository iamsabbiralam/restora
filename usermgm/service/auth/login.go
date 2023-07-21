package auth

import (
	"context"
	"errors"

	login "github.com/iamsabbiralam/restora/proto/v1/usermgm/auth"
	"github.com/iamsabbiralam/restora/usermgm/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (h *Handler) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	log := h.logger.WithField("method", "Service.Auth.Login")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	if err := h.ValidateRequestedLoginData(ctx, storage.User{
		Email:    req.Login.Email,
	}); err != nil {
		log.WithError(err).Error("validation error while login to the system")
		return nil, uErr.HandleServiceErr(err)
	}

	loginData := storage.User{
		Email:    req.Login.Email,
	}

	res, err := h.login.Login(ctx, loginData)
	if err != nil {
		errMsg := "failed to create user"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &login.LoginResponse{
		ID: res.ID,
	}, nil
}
