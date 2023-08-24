package postgres

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Storage) CreateUserInformationValidation(ctx context.Context, req *storage.UserInformation) error {
	if req == nil {
		return errors.New("invalid request")
	}

	required := validation.Required
	if err := validation.ValidateStruct(req,
		validation.Field(&req.FirstName, required),
		validation.Field(&req.LastName, required),
		validation.Field(&req.Image, required),
		validation.Field(&req.Mobile, required),
		validation.Field(&req.Gender, required),
		validation.Field(&req.Address, required),
		validation.Field(&req.City, required),
		validation.Field(&req.Country, required),
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
