package user

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/iamsabbiralam/restora/usermgm/storage"
)

type Svc struct {
	store  UserStore
	logger *logrus.Entry
}

func New(rs UserStore, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  rs,
		logger: logger,
	}
}

type UserStore interface {
	CreateUser(ctx context.Context, user storage.User) (string, error)
	UpdateUser(ctx context.Context, user storage.User) (*storage.User, error)
	DeleteUser(ctx context.Context, user storage.User) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*storage.User, error)
	GetUserByID(ctx context.Context, id string) (*storage.User, error)
	GetUserByUsername(ctx context.Context, username string) (*storage.User, error)
	ListUsers(ctx context.Context, p storage.FilterUser) ([]storage.User, error)
}
