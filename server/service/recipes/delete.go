package recipes

import (
	"context"
	"database/sql"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	recG "github.com/iamsabbiralam/restora/proto/v1/server/recipe"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) DeleteRecipe(ctx context.Context, req *recG.DeleteRecipeRequest) (*emptypb.Empty, error) {
	log := s.logger.WithField("method", "service.recipe.DeleteRecipe")
	if req == nil || req.GetID() == "" {
		return nil, uErr.HandleServiceErr(errors.New("id is required"))
	}

	recipe := storage.Recipe{
		ID: req.GetID(),
		CRUDTimeDate: storage.CRUDTimeDate{
			DeletedBy: sql.NullString{
				Valid: true,
			},
		},
	}

	err := s.store.DeleteRecipe(ctx, recipe.ID, req.DeletedBy)
	if err != nil {
		errMsg := "failed to delete recipe"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &emptypb.Empty{}, nil
}
