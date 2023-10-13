package postgres

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/iamsabbiralam/restora/server/storage"
	"github.com/pkg/errors"
)

func (s *Storage) CreateRecipeValidation(ctx context.Context, req storage.Recipe) error {
	if req == (storage.Recipe{}) {
		return errors.New("invalid request")
	}

	required := validation.Required
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Title, required),
	); err != nil {
		return err
	}

	return nil
}
