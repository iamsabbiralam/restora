package categories

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/timestamppb"

	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	catCore "github.com/iamsabbiralam/restora/server/core/categories"
	"github.com/iamsabbiralam/restora/server/storage/postgres"
	"github.com/iamsabbiralam/restora/utility/logging"
)

var _testStorage *postgres.Storage

func newTestSvc(t *testing.T, st *postgres.Storage) (*postgres.Storage, *Svc) {
	conn := os.Getenv("DATABASE_CONNECTION")
	if conn == "" {
		t.Skip("missing database connection")
	}

	st, cleanup := postgres.NewTestStorage(conn, filepath.Join("..", "..", "migrations"))
	t.Cleanup(cleanup)

	config := viper.New()
	logger := logging.NewLogger(config).WithFields(logrus.Fields{
		"service": "test_category",
	})

	return st, New(catCore.New(st, logger), logger)
}

func newTestStorage(tb testing.TB) *postgres.Storage {
	if testing.Short() {
		tb.Skip("skipping tests that use postgres on -short")
	}
	return _testStorage
}

type CategoryTestStruct struct {
	methodName string
	desc       string
	in         interface{}
	want       interface{}
	tops       cmp.Options
	wantErr    bool
}

var (
	id string
)

func TestCategories(t *testing.T) {
	st := newTestStorage(t)
	_, s := newTestSvc(t, st)
	tests := []CategoryTestStruct{
		{
			methodName: "CREATE_CATEGORY_SUCCESS",
			desc:       "Success Create Category",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.CreateCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.CreateCategoryRequest{
				Name:   "test category",
				Status: catG.Status_Active,
			},
			want: &catG.CreateCategoryResponse{
				ID: id,
			},
		},
		{
			methodName: "CREATE_CATEGORY_FAILED",
			desc:       "Failed Create Category",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.CreateCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in:         &catG.CreateCategoryRequest{},
			want:       nil,
		},
		{
			methodName: "UPDATE_CATEGORY_SUCCESS",
			desc:       "Success Update Category",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.UpdateCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.UpdateCategoryRequest{
				ID:     id,
				Name:   "test category update",
				Status: catG.Status_Inactive,
			},
			want: &catG.UpdateCategoryResponse{
				ID:     id,
				Name:   "test category update",
				Status: catG.Status_Inactive,
			},
		},
		{
			methodName: "UPDATE_CATEGORY_FAILED",
			desc:       "Failed Update Category",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.UpdateCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in:         &catG.UpdateCategoryRequest{},
			want:       nil,
		},
		{
			methodName: "GET_CATEGORY_SUCCESS",
			desc:       "Success Get Category",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.GetCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.GetCategoryRequest{
				ID: id,
			},
			want: &catG.GetCategoryResponse{
				ID:     id,
				Name:   "test category",
				Status: 1,
			},
		},
		{
			methodName: "GET_CATEGORY_FAILED",
			desc:       "Failed Get Category",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.GetCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in:         &catG.GetCategoryRequest{},
			want:       nil,
		},

		{
			methodName: "LIST_CATEGORY_SUCCESS",
			desc:       "Success List Category",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.ListCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.ListCategoryRequest{
				SortBy:       catG.SortBy_DESC,
				SortByColumn: catG.SortByColumn_Name,
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &catG.ListCategoryResponse{
				Categories: []*catG.Category{
					{
						Name:   "test category",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_CATEGORY_SUCCESS_BY_COLUMN",
			desc:       "Success List Category by SortByColumn",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.ListCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.ListCategoryRequest{
				SortBy:       catG.SortBy_ASC,
				SortByColumn: catG.SortByColumn_Name,
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &catG.ListCategoryResponse{
				Categories: []*catG.Category{
					{
						Name:   "test category",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_CATEGORY_SUCCESS_BY_SEARCH",
			desc:       "Success List Category by SearchTerm",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.ListCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.ListCategoryRequest{
				SortBy:       catG.SortBy_ASC,
				SortByColumn: catG.SortByColumn_Name,
				SearchTerm:   "BdNews",
				StartDate:    "",
				EndDate:      "",
			},
			want: &catG.ListCategoryResponse{
				Categories: []*catG.Category{
					{
						Name:   "test category",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_CATEGORY_SUCCESS_BY_DATE",
			desc:       "Success List Category by StartDate",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.ListCategoryResponse{}, catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.ListCategoryRequest{
				SortBy:       catG.SortBy_ASC,
				SortByColumn: catG.SortByColumn_Name,
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &catG.ListCategoryResponse{
				Categories: []*catG.Category{
					{
						Name:   "test category",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "DELETE_CATEGORY_SUCCESS",
			desc:       "Success Delete Category",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.Category{}, timestamppb.Timestamp{})},
			in: &catG.DeleteCategoryRequest{
				ID: id,
			},
		},
		{
			methodName: "DELETE_CATEGORY_FAILED",
			desc:       "Failed Delete Category",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(catG.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(catG.Category{}, timestamppb.Timestamp{})},
			in:         &catG.DeleteCategoryRequest{},
			want:       nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.methodName, func(t *testing.T) {
			switch test.methodName {
			case "CreateCategory":
				CreateCategoryTest(t, test, s)
			case "UpdateCategory":
				UpdateCategoryTest(t, test, s)
			case "GetCategory":
				GetCategoryTest(t, test, s)
			case "ListCategory":
				ListCategoryTest(t, test, s)
			case "DeleteCategory":
				deleteCategoryTest(t, test, s)
			}
		})
	}
}

func CreateCategoryTest(t *testing.T, test CategoryTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*catG.CreateCategoryRequest)
	if !ok {
		t.Error("request type conversion error")
	}
	got, err := s.CreateCategory(ctx, req)
	if got != nil && got.ID != "" && !test.wantErr {
		id = got.ID
	}

	if err != nil && !test.wantErr {
		t.Fatal(err)
	}
	o := test.tops
	if !test.wantErr {
		if !cmp.Equal(test.want, got, o) {
			t.Error("(-want +got): ", cmp.Diff(test.want, got, o))
		}
	}
}

func UpdateCategoryTest(t *testing.T, test CategoryTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*catG.UpdateCategoryRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	req.ID = id
	got, err := s.UpdateCategory(ctx, req)
	if err != nil && !test.wantErr {
		t.Fatal(err)
	}

	if !test.wantErr {
		o := test.tops
		if !cmp.Equal(test.want, got, o) {
			t.Error("(-want +got): ", cmp.Diff(test.want, got, o))
		}
	}
}

func GetCategoryTest(t *testing.T, test CategoryTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*catG.GetCategoryRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	req.ID = id
	got, err := s.GetCategory(ctx, req)
	if err != nil && !test.wantErr {
		t.Fatal(err)
	}

	if !test.wantErr {
		o := test.tops
		if !cmp.Equal(test.want, got, o) {
			t.Error("(-want +got): ", cmp.Diff(test.want, got, o))
		}
	}
}

func ListCategoryTest(t *testing.T, test CategoryTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*catG.ListCategoryRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	got, err := s.ListCategory(ctx, req)
	if err != nil && !test.wantErr {
		t.Fatal(err)
	}
	if !test.wantErr {
		o := test.tops
		if !cmp.Equal(test.want, got, o) {
			t.Error("(-want +got): ", cmp.Diff(test.want, got, o))
		}
	}
}

func deleteCategoryTest(t *testing.T, tc CategoryTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := tc.in.(*catG.DeleteCategoryRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	if !tc.wantErr {
		req.ID = id
	}

	got, err := s.DeleteCategory(ctx, req)
	if err != nil && !tc.wantErr {
		t.Fatal(err)
	}

	if !tc.wantErr {
		o := tc.tops
		if !cmp.Equal(tc.want, got, o) {
			t.Error("(-want +got): ", cmp.Diff(tc.want, got, o))
		}
	}
}
