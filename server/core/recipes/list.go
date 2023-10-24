package recipes

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) ListRecipe(ctx context.Context, req storage.ListRecipeFilter) ([]storage.Recipe, error) {
	s.logger.WithField("method", "core.recipes.ListRecipe")
	recipes, err := s.store.ListRecipe(ctx, req)
	if err != nil && err != storage.NotFound {
		return nil, err
	}

	return recipes, nil
}
