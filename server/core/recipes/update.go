package recipes

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) UpdateRecipe(ctx context.Context, req storage.Recipe) (*storage.Recipe, error) {
	s.logger.WithField("method", "core.recipes.UpdateRecipe")
	rec, err := s.store.UpdateRecipe(ctx, req)
	if err != nil {
		return nil, err
	}

	return rec, nil
}
