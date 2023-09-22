package categories

import "context"

func (s *Svc) DeleteCategory(ctx context.Context, id string) error {
	s.logger.WithField("method", "core.categories.DeleteCategory")
	err := s.store.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
