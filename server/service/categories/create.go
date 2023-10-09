package categories

import (
	"context"
	"errors"

	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) CreateCategory(ctx context.Context, req *catG.CreateCategoryRequest) (*catG.CreateCategoryResponse, error) {
	log := s.logger.WithField("method", "Service.Categories.CreateCategory")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	if err := s.ValidateRequestedData(ctx, storage.Category{
		Name:   req.GetName(),
		Status: int32(req.GetStatus()),
	}, ""); err != nil {
		log.WithError(err).Error("validation error while creating category")
		return nil, uErr.HandleServiceErr(err)
	}

	storeData := storage.Category{
		Name:   req.GetName(),
		Status: int32(req.GetStatus()),
	}

	res, err := s.cc.CreateCategory(ctx, storeData)
	if err != nil {
		errMsg := "failed to create category"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &catG.CreateCategoryResponse{
		ID: res,
	}, nil
}
