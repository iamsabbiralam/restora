package postgres

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Storage) CreateCategoryValidation(ctx context.Context, req storage.Category) error {
	if req == (storage.Category{}) {
		return errors.New("invalid request")
	}

	required := validation.Required
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, required),
	); err != nil {
		return err
	}

	return nil
}
