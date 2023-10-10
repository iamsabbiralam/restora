package brands

import (
	"context"
	"errors"

	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) CreateBrand(ctx context.Context, req *braG.CreateBrandRequest) (*braG.CreateBrandResponse, error) {
	log := s.logger.WithField("method", "Service.brands.CreateBrand")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	if err := s.ValidateRequestedData(ctx, storage.Brand{
		Name:   req.GetName(),
		Status: int32(req.GetStatus()),
	}, ""); err != nil {
		log.WithError(err).Error("validation error while creating category")
		return nil, uErr.HandleServiceErr(err)
	}

	storeData := storage.Brand{
		Name:   req.GetName(),
		Status: int32(req.GetStatus()),
	}

	res, err := s.cb.CreateBrand(ctx, storeData)
	if err != nil {
		errMsg := "failed to create brands"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &braG.CreateBrandResponse{
		ID: res,
	}, nil
}
