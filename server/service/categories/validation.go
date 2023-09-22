package categories

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) ValidateRequestedData(ctx context.Context, req storage.Category, id string) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, vre("Category name is required")),
		validation.Field(&req.Status, validation.Min(1).Error("Status is Invalid"), validation.Max(2).Error("Status is Invalid")),
	)
}

func (s *Svc) startDateEndDateRangeCheck(fromTime, toTime string) (string, string, error) {
	if fromTime != "" && toTime != "" && fromTime > toTime {
		return "", "", status.Error(codes.Unknown, "Invalid from and to date range")
	}
	return fromTime, toTime, nil
}
