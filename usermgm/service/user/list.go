package user

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	"github.com/iamsabbiralam/restora/usermgm/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (h *Handler) ListUsers(ctx context.Context, req *userG.ListUsersRequest) (*userG.ListUsersResponse, error) {
	log := h.logger.WithField("method", "Service.User.ListUsers")
	fT, lT, err := h.startDateEndDateRangeCheck(req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return nil, uErr.HandleServiceErr(err)
	}

	usrs, err := h.usr.ListUsers(ctx, storage.FilterUser{
		SearchTerm:   req.SearchTerm,
		Limit:        req.Limit,
		Offset:       req.Offset,
		Status:       storage.ActiveStatus(req.Status),
		SortBy:       req.GetSortBy().String(),
		SortByColumn: req.GetSortByColumn().String(),
		StartDate:    fT,
		EndDate:      lT,
	})

	if err != nil {
		log.WithError(err).Error("error with getting users")
		return nil, uErr.HandleServiceErr(err)
	}

	list := make([]*userG.User, len(usrs))
	if len(usrs) > 0 {
		for i, u := range usrs {
			list[i] = &userG.User{
				ID:        u.ID,
				UserName:  u.Username,
				Email:     u.Email,
				Status:    userG.Status(u.Status),
				CreatedAt: timestamppb.New(u.CreatedAt),
			}
		}
	}

	var total int32
	if len(usrs) > 0 {
		total = int32(usrs[0].Count)
	}

	if list == nil {
		return nil, uErr.HandleServiceErr(errors.New("error with getting users"))
	}

	return &userG.ListUsersResponse{
		Users: list,
		Total: total,
	}, nil
}

func (h *Handler) startDateEndDateRangeCheck(fromTime, toTime string) (string, string, error) {
	if fromTime != "" && toTime != "" && fromTime > toTime {
		return "", "", status.Error(codes.Unknown, "Invalid from and to date range")
	}
	return fromTime, toTime, nil
}
