package brands

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) ListBrand(ctx context.Context, req storage.ListBrandFilter) ([]storage.Brand, error) {
	s.logger.WithField("method", "core.brands.ListBrand")
	brands, err := s.store.ListBrand(ctx, req)
	if err != nil && err != storage.NotFound {
		return nil, err
	}

	return brands, nil
}
