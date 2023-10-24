package recipes

import "context"

func (s *Svc) DeleteRecipe(ctx context.Context, id, deletedBy string) error {
	s.logger.WithField("method", "core.recipes.DeleteRecipe")
	err := s.store.DeleteRecipe(ctx, id, deletedBy)
	if err != nil {
		return err
	}

	return nil
}
