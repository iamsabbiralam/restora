package adminUser

import (
	"net/http"
	"net/url"

	"github.com/gorilla/csrf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamsabbiralam/restora/client/handler/common"
	"github.com/iamsabbiralam/restora/client/handler/paginator"
	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
)

func (s *Svc) listUserHandler(w http.ResponseWriter, r *http.Request) {
	log := s.Logger.WithField("method", "handler.admin.user.listUserHandler")
	stu, err := url.PathUnescape(r.URL.Query().Get("status"))
	var status int32
	if err != nil || stu == "" {
		status = 0
	}

	switch stu {
	case "active":
		status = 1
	case "inactive":
		status = 2
	}

	queryString := common.GetQueryStringData(r, []string{}, false)
	if queryString == nil {
		errMsg := "error with getting query string data"
		log.WithError(err).Error(errMsg)
	}

	var sortBy int32
	if queryString.SortBy == "ASC" {
		sortBy = 1
	}

	var formErr string
	fT, lT, err := s.startDateEndDateRangeCheck(queryString.StartDate, queryString.EndDate)
	if err != nil {
		formErr = "Invalid start and end date range"
	}

	users, err := s.User.ListUsers(r.Context(), &userG.ListUsersRequest{
		SearchTerm: queryString.SearchTerm,
		Limit:      common.LimitPerPage,
		Offset:     queryString.Offset,
		SortBy:     userG.SortBy(sortBy),
		Status:     userG.Status(status),
		StartDate:  fT,
		EndDate:    lT,
	})
	if err != nil {
		errMsg := "unable to get list"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	if users == nil {
		errMsg := "unable to get user list"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	userTotal := users.Total
	res, err := s.getAdminUserData(w, r, users)
	if err != nil {
		errMsg := "unable to get admin user data"
		log.WithError(err).Error(errMsg)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	formMessage := map[string]string{}
	if queryString.SearchTerm != "" && users != nil && len(users.GetUsers()) > 0 {
		formMessage = map[string]string{"FoundMessage": "Data Found"}
	} else if queryString.SearchTerm != "" && users != nil && len(users.GetUsers()) == 0 {
		formMessage = map[string]string{"NotFoundMessage": "Data Not Found"}
	}

	data := s.adminUserData(w, r, res, formMessage, queryString, stu, formErr)
	if len(res) > 0 {
		data.PaginationData = paginator.NewPaginator(int32(queryString.CurrentPage), common.LimitPerPage, userTotal, r)
	}

	s.loadAdminUserTemplate(w, r, data, "admin-user-List.html")
}

func (s *Svc) getAdminUserData(w http.ResponseWriter, r *http.Request, users *userG.ListUsersResponse) ([]AdminUser, error) {
	s.Logger.WithField("method", "handler.admin.user.getAdminUserData")
	auList := make([]AdminUser, 0, len(users.GetUsers()))
	for _, item := range users.GetUsers() {
		userAppendData := AdminUser{
			ID:       item.ID,
			Username: item.GetUserName(),
			Email:    item.GetEmail(),
			Status:   item.GetStatus(),
		}

		auList = append(auList, userAppendData)
	}

	return auList, nil
}

func (s *Svc) adminUserData(w http.ResponseWriter, r *http.Request, auList []AdminUser, formMSG map[string]string, queryString *common.DynamicQueryString, stu string, formErr string) AdminUserTempData {
	return AdminUserTempData{
		CSRFField:     csrf.TemplateField(r),
		Data:          auList,
		FormMessage:   formMSG,
		SearchTerm:    queryString.SearchTerm,
		StartDate:     queryString.StartDate,
		SortBy:        queryString.SortBy,
		SortByColumn:  queryString.SortByColumn,
		EndDate:       queryString.EndDate,
		PramStatus:    stu,
		FilterFormErr: formErr,
	}
}

func (s *Svc) startDateEndDateRangeCheck(fromTime, toTime string) (string, string, error) {
	if fromTime != "" && toTime != "" && fromTime > toTime {
		return "", "", status.Error(codes.Unknown, "Invalid from and to date range")
	}
	return fromTime, toTime, nil
}
