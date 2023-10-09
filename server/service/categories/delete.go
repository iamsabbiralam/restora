package categories

import (
	"context"
	"database/sql"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) DeleteCategory(ctx context.Context, req *catG.DeleteCategoryRequest) (*emptypb.Empty, error) {
	log := s.logger.WithField("method", "service.categories.DeleteCategory")
	if req == nil || req.GetID() == "" {
		return nil, uErr.HandleServiceErr(errors.New("id is required"))
	}

	cat := storage.Category{
		ID: req.GetID(),
		CRUDTimeDate: storage.CRUDTimeDate{
			DeletedBy: sql.NullString{
				Valid: true,
			},
		},
	}

	err := s.cc.DeleteCategory(ctx, cat.ID, req.DeletedBy)
	if err != nil {
		errMsg := "failed to delete category"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &emptypb.Empty{}, nil
}
