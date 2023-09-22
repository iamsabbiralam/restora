package categories

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) GetCategoryByID(ctx context.Context, id string) (*storage.Category, error) {
	s.logger.WithField("method", "core.categories.GetCategoryByID")
	cat, err := s.store.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cat, nil
}
