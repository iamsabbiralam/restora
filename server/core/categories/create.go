package categories

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) CreateCategory(ctx context.Context, req storage.Category) (string, error) {
	s.logger.WithField("method", "core.categories.CreateCategory")
	id, err := s.store.CreateCategory(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}
