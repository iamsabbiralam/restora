package categories

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) UpdateCategory(ctx context.Context, req *catG.UpdateCategoryRequest) (*catG.UpdateCategoryResponse, error) {
	log := s.logger.WithField("method", "service.categories.UpdateCategory")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	if err := s.ValidateRequestedData(ctx, storage.Category{
		ID:     req.GetID(),
		Name:   req.GetName(),
		Status: int32(catG.Status_Active),
	}, req.GetID()); err != nil {
		log.WithError(err).Error("validation error while updating category")
		return nil, uErr.HandleServiceErr(err)
	}

	updateData := storage.Category{
		ID:     req.GetID(),
		Name:   req.GetName(),
		Status: int32(req.GetStatus()),
	}
	res, err := s.cc.UpdateCategory(ctx, updateData)
	if err != nil {
		errMsg := "failed to update category"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if res == nil {
		return nil, uErr.HandleServiceErr(errors.New("update category response is nil"))
	}

	resp := &catG.UpdateCategoryResponse{
		ID:        res.ID,
		Name:      res.Name,
		Status:    catG.Status(res.Status),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
		UpdatedBy: res.UpdatedBy,
	}

	return resp, nil
}
