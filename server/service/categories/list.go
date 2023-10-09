package categories

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) ListCategory(ctx context.Context, req *catG.ListCategoryRequest) (*catG.ListCategoryResponse, error) {
	log := s.logger.WithField("method", "service.categories.ListCategory")
	startDate, endDate, err := s.startDateEndDateRangeCheck(req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return nil, uErr.HandleServiceErr(err)
	}

	categories, err := s.cc.ListCategories(ctx, storage.ListCategoryFilter{
		SearchTerm:   req.SearchTerm,
		Limit:        req.Limit,
		Offset:       req.Offset,
		Status:       storage.ActiveStatus(req.Status),
		SortBy:       req.GetSortBy().String(),
		SortByColumn: req.GetSortByColumn(),
		StartDate:    startDate,
		EndDate:      endDate,
	})

	if err != nil {
		log.WithError(err).Error("error with getting categories")
		return nil, uErr.HandleServiceErr(err)
	}

	list := make([]*catG.Category, len(categories))
	if len(categories) > 0 {
		for i, val := range categories {
			list[i] = &catG.Category{
				ID:        val.ID,
				Name:      val.Name,
				Status:    catG.Status(val.Status),
				CreatedAt: timestamppb.New(val.CreatedAt),
				UpdatedAt: timestamppb.New(val.UpdatedAt),
			}
		}
	}

	var total int32
	if len(categories) > 0 {
		total = int32(categories[0].Count)
	}

	if list == nil {
		return nil, uErr.HandleServiceErr(errors.New("error with getting users"))
	}

	return &catG.ListCategoryResponse{
		Categories: list,
		Total:      total,
	}, nil
}
