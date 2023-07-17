package user

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Svc) DeleteUser(ctx context.Context, user storage.User) (string, error) {
	log := s.logger.WithField("method", "Core.User.DeleteUser")
	id, err := s.store.DeleteUser(ctx, user)
	if err != nil {
		errMsg := "Failed to delete user storage entry"
		log.WithError(err).Error(errMsg)
		return "", err
	}

	return id, nil
}
