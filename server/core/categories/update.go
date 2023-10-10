package categories

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) UpdateCategory(ctx context.Context, req storage.Category) (*storage.Category, error) {
	s.logger.WithField("method", "core.categories.UpdateCategory")
	cat, err := s.store.UpdateCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return cat, nil
}
