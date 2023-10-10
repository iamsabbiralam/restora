package brands

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) CreateBrand(ctx context.Context, req storage.Brand) (string, error) {
	s.logger.WithField("method", "core.brands.CreateBrand")
	id, err := s.store.CreateBrand(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}
