package user

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Svc) UpdateUser(ctx context.Context, user storage.User) (*storage.User, error) {
	log := s.logger.WithField("method", "Core.User.UpdateUser")
	usr, err := s.store.UpdateUser(ctx, user)
	if err != nil {
		errMsg := "Failed to update user storage entry"
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	return usr, nil
}
