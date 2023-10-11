package user

import (
	"context"
	"database/sql"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	upb "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"github.com/iamsabbiralam/restora/usermgm/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (h *Handler) DeleteUser(ctx context.Context, req *upb.DeleteUserRequest) (*emptypb.Empty, error) {
	log := h.logger.WithField("method", "Service.User.DeleteUser")
	if req == nil || req.GetID() == "" {
		return nil, uErr.HandleServiceErr(errors.New("user id is required"))
	}

	dbUser := storage.User{
		ID: req.GetID(),
		DeletedBy: sql.NullString{
			// String: hydra.GetUserID(ctx),
			Valid: true,
		},
	}

	err := h.usr.DeleteUser(ctx, dbUser)
	if err != nil {
		errMsg := "failed to delete user"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &emptypb.Empty{}, nil
}
