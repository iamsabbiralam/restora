package auth

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	upb "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"github.com/iamsabbiralam/restora/usermgm/storage"
)

type Handler struct {
	upb.UnimplementedUserServiceServer
	login  CoreLogin
	logger *logrus.Entry
}

type CoreLogin interface {
	Login(context.Context, storage.User) (*storage.User, error)
}

func New(login CoreLogin, logger *logrus.Entry) *Handler {
	return &Handler{login: login, logger: logger}
}

// // RegisterService with grpc server.
func (h *Handler) RegisterSvc(srv *grpc.Server) error {
	upb.RegisterUserServiceServer(srv, h)
	return nil
}

func Permission(ctx context.Context) map[string]storage.ResAct {
	p := map[string]storage.ResAct{
		"Login": {Resource: "user", Action: "Login"},
	}
	return p
}
