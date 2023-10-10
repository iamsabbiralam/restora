package brands

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) UpdateBrand(ctx context.Context, req *braG.UpdateBrandRequest) (*braG.UpdateBrandResponse, error) {
	log := s.logger.WithField("method", "service.brands.UpdateBrand")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	if err := s.ValidateRequestedData(ctx, storage.Brand{
		ID:     req.GetID(),
		Name:   req.GetName(),
		Status: int32(req.GetStatus()),
	}, req.GetID()); err != nil {
		log.WithError(err).Error("validation error while updating brand")
		return nil, uErr.HandleServiceErr(err)
	}

	updateData := storage.Brand{
		ID:     req.GetID(),
		Name:   req.GetName(),
		Status: int32(req.GetStatus()),
	}
	res, err := s.cb.UpdateBrand(ctx, updateData)
	if err != nil {
		errMsg := "failed to update brand"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if res == nil {
		return nil, uErr.HandleServiceErr(errors.New("update brand response is nil"))
	}

	resp := &braG.UpdateBrandResponse{
		ID:        res.ID,
		Name:      res.Name,
		Status:    braG.Status(res.Status),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
		UpdatedBy: res.UpdatedBy,
	}

	return resp, nil
}
