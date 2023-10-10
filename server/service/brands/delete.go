package brands

import (
	"context"
	"database/sql"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) DeleteBrand(ctx context.Context, req *braG.DeleteBrandRequest) (*emptypb.Empty, error) {
	log := s.logger.WithField("method", "service.brands.DeleteBrand")
	if req == nil || req.GetID() == "" {
		return nil, uErr.HandleServiceErr(errors.New("id is required"))
	}

	cat := storage.Brand{
		ID: req.GetID(),
		CRUDTimeDate: storage.CRUDTimeDate{
			DeletedBy: sql.NullString{
				Valid: true,
			},
		},
	}

	err := s.cb.DeleteBrand(ctx, cat.ID, req.DeletedBy)
	if err != nil {
		errMsg := "failed to delete brand"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &emptypb.Empty{}, nil
}
