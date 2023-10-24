package recipes

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) GetRecipeByID(ctx context.Context, id string) (*storage.Recipe, error) {
	s.logger.WithField("method", "core.recipes.GetRecipeByID")
	rec, err := s.store.GetRecipeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return rec, nil
}
