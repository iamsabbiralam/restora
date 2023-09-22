package categories

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) GetCategory(ctx context.Context, req *catG.GetCategoryRequest) (*catG.GetCategoryResponse, error) {
	log := s.logger.WithField("method", "service.categories.GetCategory")
	if req.GetID() == "" {
		return nil, uErr.HandleServiceErr(errors.New("category id is required"))
	}

	r, err := s.cc.GetCategoryByID(ctx, req.GetID())
	if err != nil {
		errMsg := "failed to get category by id"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if r == nil {
		log.WithError(err).Error("error while getting category by id")
		return nil, uErr.HandleServiceErr(err)
	}

	res := &catG.GetCategoryResponse{
		ID:        r.ID,
		Name:      r.Name,
		Status:    catG.Status(r.Status),
		CreatedAt: timestamppb.New(r.CreatedAt),
		CreatedBy: r.CreatedBy,
		UpdatedAt: timestamppb.New(r.UpdatedAt),
		UpdatedBy: r.UpdatedBy,
	}

	return res, nil
}
