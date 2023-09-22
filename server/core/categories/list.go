package categories

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) ListCategories(ctx context.Context, req storage.ListCategoryFilter) ([]storage.Category, error) {
	s.logger.WithField("method", "core.categories.ListCategories")
	leaveReq, err := s.store.ListCategories(ctx, req)
	if err != nil && err != storage.NotFound {
		return nil, err
	}

	return leaveReq, nil
}
