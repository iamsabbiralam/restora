package user

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	upb "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"github.com/iamsabbiralam/restora/usermgm/storage"
)

type Handler struct {
	upb.UnimplementedUserServiceServer
	usr    CoreUser
	logger *logrus.Entry
}

type CoreUser interface {
	CreateUser(context.Context, storage.User) (string, error)
	UpdateUser(context.Context, storage.User) (*storage.User, error)
	DeleteUser(context.Context, storage.User) error
	GetUserByEmail(context.Context, string) (*storage.User, error)
	GetUserByID(context.Context, string) (*storage.User, error)
	GetUserByUsername(context.Context, string) (*storage.User, error)
	ListUsers(context.Context, storage.FilterUser) ([]storage.User, error)
	GetUserInformationByUserID(context.Context, string) (*storage.UserInformation, error)
	UpdateUserInformationByUserID(context.Context, storage.UserInformation) (*storage.UserInformation, error)
}

func New(usr CoreUser, logger *logrus.Entry) *Handler {
	return &Handler{usr: usr, logger: logger}
}

// // RegisterService with grpc server.
func (h *Handler) RegisterSvc(srv *grpc.Server) error {
	upb.RegisterUserServiceServer(srv, h)
	return nil
}

func Permission(ctx context.Context) map[string]storage.ResAct {
	p := map[string]storage.ResAct{
		"CreateUser":        {Resource: "user", Action: "Create"},
		"UpdateUser":        {Resource: "user", Action: "Update"},
		"DeleteUser":        {Resource: "user", Action: "Delete"},
		"GetUserByEmail":    {Resource: "user", Action: "Read"},
		"GetUserByID":       {Resource: "user", Action: "Read", Public: true},
		"GetUserByUsername": {Resource: "user", Action: "Read"},
		"ListUsers":         {Resource: "user", Action: "Read"},
		"InviteUser":        {Resource: "user", Action: "Create"},
		"AcceptInvitation":  {Resource: "user", Action: "Update"},
	}
	return p
}

func (h *Handler) validateUserName(value string, id string) validation.Rule {
	return validation.By(func(interface{}) error {
		res, err := h.usr.GetUserByUsername(context.Background(), value)
		if err != nil && status.Convert(err).Code().String() != "NotFound" {
			return err
		}

		if res == nil {
			return errors.New("unable to get user by username")
		}

		if id != "" && res.ID == id && res.Username == value {
			return nil
		}

		if res.Username == value {
			return errors.New("username already exists")
		}

		return nil
	})
}

func (h *Handler) validateUserEmail(email string, id string) validation.Rule {
	return validation.By(func(interface{}) error {
		res, err := h.usr.GetUserByEmail(context.Background(), email)
		if err != nil && status.Convert(err).Code().String() != "NotFound" {
			return err
		}

		if res == nil {
			return errors.New("unable to get user by email")
		}

		if id != "" && res.ID == id && res.Email == email {
			return nil
		}

		if res.Email == email {
			return errors.New(" Email already exists")
		}

		return nil
	})
}

func (h *Handler) ValidateRequestedData(ctx context.Context, req storage.User, id string) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, vre("Username is required"), h.validateUserName(req.Username, id)),
		validation.Field(&req.Email, vre("Email is required"), is.EmailFormat.Error("The email is not valid"), h.validateUserEmail(req.Email, id)),
		validation.Field(&req.Status, validation.Min(1).Error("Status is Invalid"), validation.Max(2).Error("Status is Invalid")),
	)
}
