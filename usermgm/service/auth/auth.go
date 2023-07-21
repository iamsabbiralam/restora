package auth

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	auth "github.com/iamsabbiralam/restora/proto/v1/usermgm/auth"
	"github.com/iamsabbiralam/restora/usermgm/storage"
)

type Handler struct {
	auth.UnimplementedLoginServiceServer
	login  CoreLogin
	user   CoreUser
	logger *logrus.Entry
}

type CoreLogin interface {
	Login(context.Context, storage.User) (*storage.User, error)
}

type CoreUser interface {
	GetUserByEmail(context.Context, string) (*storage.User, error)
}

func New(login CoreLogin, user CoreUser, logger *logrus.Entry) *Handler {
	return &Handler{login: login, logger: logger, user: user}
}

// // RegisterService with grpc server.
func (h *Handler) RegisterSvc(srv *grpc.Server) error {
	auth.RegisterLoginServiceServer(srv, h)
	return nil
}

func Permission(ctx context.Context) map[string]storage.ResAct {
	p := map[string]storage.ResAct{
		"Login": {Resource: "user", Action: "Login"},
	}
	return p
}

func (h *Handler) ValidateRequestedLoginData(ctx context.Context, req storage.User) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, vre("Email is required"), is.EmailFormat.Error("The email is not valid"), h.validateLoginEmail(req.Email)),
	)
}

func (h *Handler) validateLoginEmail(email string) validation.Rule {
	return validation.By(func(interface{}) error {
		_, err := h.user.GetUserByEmail(context.Background(), email)
		if err != nil && status.Convert(err).Code().String() != "NotFound" {
			return err
		}

		return nil
	})
}
