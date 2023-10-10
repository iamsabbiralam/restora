package brands

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) ListBrand(ctx context.Context, req *braG.ListBrandRequest) (*braG.ListBrandResponse, error) {
	log := s.logger.WithField("method", "service.brands.ListBrand")
	startDate, endDate, err := s.startDateEndDateRangeCheck(req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return nil, uErr.HandleServiceErr(err)
	}

	brands, err := s.cb.ListBrand(ctx, storage.ListBrandFilter{
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
		log.WithError(err).Error("error with getting brands")
		return nil, uErr.HandleServiceErr(err)
	}

	list := make([]*braG.Brand, len(brands))
	if len(brands) > 0 {
		for i, val := range brands {
			list[i] = &braG.Brand{
				ID:        val.ID,
				Name:      val.Name,
				Status:    braG.Status(val.Status),
				CreatedAt: timestamppb.New(val.CreatedAt),
				UpdatedAt: timestamppb.New(val.UpdatedAt),
			}
		}
	}

	var total int32
	if len(brands) > 0 {
		total = int32(brands[0].Count)
	}

	if list == nil {
		return nil, uErr.HandleServiceErr(errors.New("error with getting users"))
	}

	return &braG.ListBrandResponse{
		Categories: list,
		Total:      total,
	}, nil
}
