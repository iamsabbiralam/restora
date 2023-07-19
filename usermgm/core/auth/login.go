package auth

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

func (s *Svc) Login(ctx context.Context, user storage.User) (*storage.User, error) {
	log := s.logger.WithField("method", "Core.Auth.Login")
	usr, err := s.store.Login(ctx, user)
	if err != nil {
		errMsg := "Failed to login"
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	return usr, nil
}
