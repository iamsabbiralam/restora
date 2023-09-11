package user

import (
	"context"
	"database/sql"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

var errMsg = "failed to get user"

func (s *Svc) GetUserByEmail(ctx context.Context, email string) (*storage.User, error) {
	log := s.logger.WithField("method", "Core.User.GetUser")
	res, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	return res, nil
}

func (s *Svc) GetUserByID(ctx context.Context, id string) (*storage.User, error) {
	log := s.logger.WithField("method", "Core.User.GetUserByID")
	res, err := s.store.GetUserByID(ctx, id)
	if err != nil {
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, storage.NotFound
	}

	return res, nil
}

func (s *Svc) GetUserByUsername(ctx context.Context, username string) (*storage.User, error) {
	log := s.logger.WithField("method", "Core.User.GetUserByUsername")
	res, err := s.store.GetUserByUsername(ctx, username)
	if err != nil {
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	return res, nil
}

func (s *Svc) GetUserInformationByUserID(ctx context.Context, id string) (*storage.UserInformation, error) {
	log := s.logger.WithField("method", "Core.Profile.GetUserInformationByUserID")
	res, err := s.store.GetUserByID(ctx, id)
	if err != nil {
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	info, err := s.store.GetUserInformation(ctx, res.ID)
	if err != nil {
		log.WithError(err).Error(errMsg)
		return nil, err
	}

	return info, nil
}
