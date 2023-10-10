package brands

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) UpdateBrand(ctx context.Context, req storage.Brand) (*storage.Brand, error) {
	s.logger.WithField("method", "core.brands.UpdateBrands")
	bra, err := s.store.UpdateBrand(ctx, req)
	if err != nil {
		return nil, err
	}

	return bra, nil
}
