package postgres

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Storage) CreateUserInformationValidation(ctx context.Context, req storage.UserInformation) error {
	if req == (storage.UserInformation{}) {
		return errors.New("invalid request")
	}

	required := validation.Required
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.UserID, required),
	); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteUserInformationValidation(ctx context.Context, req string) error {
	if req == "" {
		return errors.New("invalid request")
	}

	return nil
}
