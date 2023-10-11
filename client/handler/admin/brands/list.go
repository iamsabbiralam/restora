package brands

import (
	"net/http"
	"net/url"

	"github.com/gorilla/csrf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamsabbiralam/restora/client/handler/common"
	"github.com/iamsabbiralam/restora/client/handler/paginator"
	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
)

func (s *Svc) listBrandHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.WithField("method", "handler.admin.brand.listBrandHandler")
	stu, err := url.PathUnescape(r.URL.Query().Get("status"))
	var status int32
	if err != nil || stu != "" {
		status = 0
	}

	switch stu {
	case "active":
		status = 1
	case "inactive":
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

	braList, err := s.Brand.ListBrand(r.Context(), &braG.ListBrandRequest{
		SortBy:       braG.SortBy(sortBy),
		Status:       braG.Status(status),
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

	brands := []Brand{}
	if braList != nil {
		for _, item := range braList.GetBrands() {
			braAppendData := Brand{
				ID:        item.ID,
				Name:      item.Name,
				Status:    item.Status,
				CreatedAt: item.CreatedAt.AsTime(),
				UpdatedAt: item.UpdatedAt.AsTime(),
			}
			brands = append(brands, braAppendData)
		}
	}

	formMSG := map[string]string{}
	if queryString.SearchTerm != "" && len(braList.GetBrands()) > 0 {
		formMSG = map[string]string{"FoundMessage": "Data Found"}
	} else if queryString.SearchTerm != "" && len(braList.GetBrands()) == 0 {
		formMSG = map[string]string{"NotFoundMessage": "Data Not Found"}
	}

	data := BrandTempData{
		CSRFField:     csrf.TemplateField(r),
		Data:          brands,
		FormMessage:   formMSG,
		SearchTerm:    queryString.SearchTerm,
		StartDate:     queryString.StartDate,
		SortBy:        queryString.SortBy,
		SortByColumn:  queryString.SortByColumn,
		PramStatus:    stu,
		Status:        common.GetStatus(braG.Status_name),
		EndDate:       queryString.EndDate,
		FilterFormErr: formErr,
	}

	if len(brands) > 0 {
		data.PaginationData = paginator.NewPaginator(int32(queryString.CurrentPage), common.LimitPerPage, braList.Total, r)
	}

	s.loadBrandTemplate(w, r, data, "brand-list.html")
}

func (s *Svc) startDateEndDateRangeCheck(fromTime, toTime string) (string, string, error) {
	if fromTime != "" && toTime != "" && fromTime > toTime {
		return "", "", status.Error(codes.Unknown, "Invalid from and to date range")
	}
	return fromTime, toTime, nil
}
