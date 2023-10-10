package brands

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) GetBrand(ctx context.Context, req *braG.GetBrandRequest) (*braG.GetBrandResponse, error) {
	log := s.logger.WithField("method", "service.brands.GetBrand")
	if req.GetID() == "" {
		return nil, uErr.HandleServiceErr(errors.New("ID is required"))
	}

	r, err := s.cb.GetBrandByID(ctx, req.GetID())
	if err != nil {
		errMsg := "failed to get brand by id"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if r == nil {
		log.WithError(err).Error("error while getting brands by id")
		return nil, uErr.HandleServiceErr(err)
	}

	res := &braG.GetBrandResponse{
		ID:        r.ID,
		Name:      r.Name,
		Status:    braG.Status(r.Status),
		CreatedAt: timestamppb.New(r.CreatedAt),
		CreatedBy: r.CreatedBy,
		UpdatedAt: timestamppb.New(r.UpdatedAt),
		UpdatedBy: r.UpdatedBy,
	}

	return res, nil
}
