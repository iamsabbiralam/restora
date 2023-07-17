package user

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Svc) DeleteUser(ctx context.Context, user storage.User) error {
	log := s.logger.WithField("method", "Core.User.DeleteUser")
	err := s.store.DeleteUser(ctx, user)
	if err != nil {
		errMsg := "Failed to delete user storage entry"
		log.WithError(err).Error(errMsg)
		return nil
	}

	return nil
}
