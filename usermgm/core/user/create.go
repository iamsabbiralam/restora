package user

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Svc) CreateUser(ctx context.Context, user storage.User) (string, error) {
	log := s.logger.WithField("method", "Core.User.CreateUser")
	usr, err := s.store.CreateUser(ctx, user)
	if err != nil {
		errMsg := "Failed to create user storage entry"
		log.WithError(err).Error(errMsg)
		return "", err
	}

	return usr, nil
}
