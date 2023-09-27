package categories

import (
	"net/http"
	"net/url"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gorilla/csrf"
	"github.com/iamsabbiralam/restora/client/handler/common"
	"github.com/iamsabbiralam/restora/client/handler/paginator"
	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
)

func (s *Svc) listCategoryHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "departmentListHandler")
	stu, err := url.PathUnescape(r.URL.Query().Get("status"))
	var status int32
	if err != nil || stu != "" {
		status = 0
	}

	if stu == "active" {
		status = 1
	}

	if stu == "inactive" {
		status = 2
	}

	queryString := common.GetQueryStringData(r, []string{}, false)
	var sortBy int32
	if queryString.SortBy == "ASC" {
		sortBy = 1
	}
	
	var formErr string
	fT, lT, err := s.startDateEndDateRangeCheck(queryString.StartDate, queryString.EndDate)
	if err != nil {
		formErr = "Invalid start and end date range"
	}

	catList, err := s.Category.ListCategory(r.Context(), &catG.ListCategoryRequest{
		SortBy:       catG.SortBy(sortBy),
		Status:       catG.Status(status),
		SortByColumn: queryString.SortByColumn,
		StartDate:    fT,
		EndDate:      lT,
		SearchTerm:   queryString.SearchTerm,
		Limit:        common.LimitPerPage,
		Offset:       queryString.Offset,
	})
	if err != nil {
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}

	categories := []Category{}
	if catList != nil {
		for _, item := range catList.GetCategories() {
			catAppendData := Category{
				ID:        item.ID,
				Name:      item.Name,
				Status:    item.Status,
				CreatedAt: item.CreatedAt.AsTime(),
				UpdatedAt: item.UpdatedAt.AsTime(),
			}
			categories = append(categories, catAppendData)
		}
	}

	formMSG := map[string]string{}
	if queryString.SearchTerm != "" && len(catList.GetCategories()) > 0 {
		formMSG = map[string]string{"FoundMessage": "Data Found"}
	} else if queryString.SearchTerm != "" && len(catList.GetCategories()) == 0 {
		formMSG = map[string]string{"NotFoundMessage": "Data Not Found"}
	}

	data := CategoryTempData{
		CSRFField:     csrf.TemplateField(r),
		Data:          categories,
		FormMessage:   formMSG,
		SearchTerm:    queryString.SearchTerm,
		StartDate:     queryString.StartDate,
		SortBy:        queryString.SortBy,
		SortByColumn:  queryString.SortByColumn,
		PramStatus:    stu,
		Status:        common.GetStatus(catG.Status_name),
		EndDate:       queryString.EndDate,
		FilterFormErr: formErr,
	}

	if len(categories) > 0 {
		data.PaginationData = paginator.NewPaginator(int32(queryString.CurrentPage), common.LimitPerPage, catList.Total, r)
	}

	s.loadCategoryTemplate(w, r, data, "category-list.html")
}

func (s *Svc) startDateEndDateRangeCheck(fromTime, toTime string) (string, string, error) {
	if fromTime != "" && toTime != "" && fromTime > toTime {
		return "", "", status.Error(codes.Unknown, "Invalid from and to date range")
	}
	return fromTime, toTime, nil
}
