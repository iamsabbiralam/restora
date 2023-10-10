package postgres

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Storage) CreateBrandValidation(ctx context.Context, req storage.Brand) error {
	if req == (storage.Brand{}) {
		fmt.Println("CreateBrandValidation")
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
