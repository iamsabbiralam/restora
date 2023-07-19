package auth

import (
	"context"

	"github.com/iamsabbiralam/restora/usermgm/storage"
	"github.com/sirupsen/logrus"
)

type Svc struct {
	store  Login
	logger *logrus.Entry
}

func New(rs Login, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  rs,
		logger: logger,
	}
}

type Login interface {
	Login(context.Context, storage.User) (*storage.User, error)
}
