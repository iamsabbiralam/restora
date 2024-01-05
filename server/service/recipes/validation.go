package recipes

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) ValidateRequestedData(ctx context.Context, req storage.Recipe, id string) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&req,
		validation.Field(&req.Title, vre("The title field is required")),
		validation.Field(&req.Ingredient, vre("The ingredient field is required")),
		validation.Field(&req.Image, vre("The image field is required")),
		validation.Field(&req.Description, vre("The description field is required")),
		validation.Field(&req.AuthorSocialLink, vre("The authorized social link field is required")),
		validation.Field(&req.ServingAmount, vre("The serving amount field is required")),
		validation.Field(&req.CookingTime, vre("The cooking time field is required")),
		validation.Field(&req.Status, validation.Min(1).Error("Status is Invalid"), validation.Max(2).Error("Status is Invalid")),
	)
}

func (s *Svc) startDateEndDateRangeCheck(fromTime, toTime string) (string, string, error) {
	if fromTime != "" && toTime != "" && fromTime > toTime {
		return "", "", status.Error(codes.Unknown, "Invalid from and to date range")
	}
	return fromTime, toTime, nil
}
