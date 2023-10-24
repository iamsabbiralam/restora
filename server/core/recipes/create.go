package recipes

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) CreateRecipe(ctx context.Context, req storage.Recipe) (string, error) {
	s.logger.WithField("method", "core.recipes.CreateRecipe")
	id, err := s.store.CreateRecipe(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}
