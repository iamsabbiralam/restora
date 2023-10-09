package categories

import "context"

func (s *Svc) DeleteCategory(ctx context.Context, id, deletedBy string) error {
	s.logger.WithField("method", "core.categories.DeleteCategory")
	err := s.store.DeleteCategory(ctx, id, deletedBy)
	if err != nil {
		return err
	}

	return nil
}
