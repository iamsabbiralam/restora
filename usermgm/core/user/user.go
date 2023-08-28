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
	CreateUser(context.Context, storage.User) (string, error)
	UpdateUser(context.Context, storage.User) (*storage.User, error)
	DeleteUser(context.Context, storage.User) error
	GetUserByEmail(context.Context, string) (*storage.User, error)
	GetUserByID(context.Context, string) (*storage.User, error)
	GetUserByUsername(context.Context, string) (*storage.User, error)
	ListUsers(context.Context, storage.FilterUser) ([]storage.User, error)
	CreateUserInformation(context.Context, storage.UserInformation) (string, error)
	GetUserInformation(context.Context, string) (*storage.UserInformation, error)
	UpdateUserInformation(context.Context, storage.UserInformation) (*storage.UserInformation, error)
	DeleteUserInformation(context.Context, string, string) error
}
