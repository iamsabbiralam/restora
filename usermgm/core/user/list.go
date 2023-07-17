package user

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Svc) ListUsers(ctx context.Context, p storage.FilterUser) ([]storage.User, error) {
	log := s.logger.WithField("method", "core.user.ListUsers")
	users, err := s.store.ListUsers(ctx, p)
	if err != nil {
		errMsg := "Failed to list admin user"
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	return users, nil
}
