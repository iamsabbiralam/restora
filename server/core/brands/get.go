package brands

import (
	"context"

	"github.com/iamsabbiralam/restora/server/storage"
)

func (s *Svc) GetBrandByID(ctx context.Context, id string) (*storage.Brand, error) {
	s.logger.WithField("method", "core.brands.GetBrandByID")
	bra, err := s.store.GetBrandByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return bra, nil
}
