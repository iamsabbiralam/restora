package brands

import "context"

func (s *Svc) DeleteBrand(ctx context.Context, id, deletedBy string) error {
	s.logger.WithField("method", "core.brands.DeleteBrand")
	err := s.store.DeleteBrand(ctx, id, deletedBy)
	if err != nil {
		return err
	}

	return nil
}
