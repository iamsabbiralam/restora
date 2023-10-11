package brands

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

	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
	braCore "github.com/iamsabbiralam/restora/server/core/brands"
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
		"service": "test_brands",
	})

	return st, New(braCore.New(st, logger), logger)
}

func newTestStorage(tb testing.TB) *postgres.Storage {
	if testing.Short() {
		tb.Skip("skipping tests that use postgres on -short")
	}

	return _testStorage
}

type BrandTestStruct struct {
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

func TestBrands(t *testing.T) {
	st := newTestStorage(t)
	_, s := newTestSvc(t, st)
	tests := []BrandTestStruct{
		{
			methodName: "CREATE_BRAND_SUCCESS",
			desc:       "Success Create Brand",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.CreateBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.CreateBrandRequest{
				Name:   "test brand",
				Status: braG.Status_Active,
			},
			want: &braG.CreateBrandResponse{
				ID: id,
			},
		},
		{
			methodName: "CREATE_BRAND_FAILED",
			desc:       "Failed Create Brand",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.CreateBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in:         &braG.CreateBrandRequest{},
			want:       nil,
		},
		{
			methodName: "UPDATE_BRAND_SUCCESS",
			desc:       "Success Update Brand",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.UpdateBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.UpdateBrandRequest{
				ID:     id,
				Name:   "test brand update",
				Status: braG.Status_Inactive,
			},
			want: &braG.UpdateBrandResponse{
				ID:     id,
				Name:   "test brand update",
				Status: braG.Status_Inactive,
			},
		},
		{
			methodName: "UPDATE_BRAND_FAILED",
			desc:       "Failed Update Brand",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.UpdateBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in:         &braG.UpdateBrandRequest{},
			want:       nil,
		},
		{
			methodName: "GET_BRAND_SUCCESS",
			desc:       "Success Get Brand",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.GetBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.GetBrandRequest{
				ID: id,
			},
			want: &braG.GetBrandResponse{
				ID:     id,
				Name:   "test brand",
				Status: 1,
			},
		},
		{
			methodName: "GET_BRAND_FAILED",
			desc:       "Failed Get Brand",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.GetBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in:         &braG.GetBrandRequest{},
			want:       nil,
		},

		{
			methodName: "LIST_BRAND_SUCCESS",
			desc:       "Success List Brand",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.ListBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.ListBrandRequest{
				SortBy:       braG.SortBy_DESC,
				SortByColumn: "",
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &braG.ListBrandResponse{
				Brands: []*braG.Brand{
					{
						Name:   "test brand",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_BRAND_SUCCESS_BY_COLUMN",
			desc:       "Success List Brand by SortByColumn",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.ListBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.ListBrandRequest{
				SortBy:       braG.SortBy_ASC,
				SortByColumn: "",
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &braG.ListBrandResponse{
				Brands: []*braG.Brand{
					{
						Name:   "test brand",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_BRAND_SUCCESS_BY_SEARCH",
			desc:       "Success List Brand by SearchTerm",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.ListBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.ListBrandRequest{
				SortBy:       braG.SortBy_ASC,
				SortByColumn: "",
				SearchTerm:   "BdNews",
				StartDate:    "",
				EndDate:      "",
			},
			want: &braG.ListBrandResponse{
				Brands: []*braG.Brand{
					{
						Name:   "test brand",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_BRAND_SUCCESS_BY_DATE",
			desc:       "Success List Brand by StartDate",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.ListBrandResponse{}, braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.ListBrandRequest{
				SortBy:       braG.SortBy_ASC,
				SortByColumn: "",
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &braG.ListBrandResponse{
				Brands: []*braG.Brand{
					{
						Name:   "test brand",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "DELETE_BRAND_SUCCESS",
			desc:       "Success Delete Brand",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.Brand{}, timestamppb.Timestamp{})},
			in: &braG.DeleteBrandRequest{
				ID: id,
			},
		},
		{
			methodName: "DELETE_BRAND_FAILED",
			desc:       "Failed Delete Brand",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(braG.Brand{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(braG.Brand{}, timestamppb.Timestamp{})},
			in:         &braG.DeleteBrandRequest{},
			want:       nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.methodName, func(t *testing.T) {
			switch test.methodName {
			case "CreateBrand":
				CreateBrandTest(t, test, s)
			case "UpdateBrand":
				UpdateBrandTest(t, test, s)
			case "GetBrand":
				GetBrandTest(t, test, s)
			case "ListBrand":
				ListBrandTest(t, test, s)
			case "DeleteBrand":
				DeleteBrandTest(t, test, s)
			}
		})
	}
}

func CreateBrandTest(t *testing.T, test BrandTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*braG.CreateBrandRequest)
	if !ok {
		t.Error("request type conversion error")
	}
	got, err := s.CreateBrand(ctx, req)
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

func UpdateBrandTest(t *testing.T, test BrandTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*braG.UpdateBrandRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	req.ID = id
	got, err := s.UpdateBrand(ctx, req)
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

func GetBrandTest(t *testing.T, test BrandTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*braG.GetBrandRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	req.ID = id
	got, err := s.GetBrand(ctx, req)
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

func ListBrandTest(t *testing.T, test BrandTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*braG.ListBrandRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	got, err := s.ListBrand(ctx, req)
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

func DeleteBrandTest(t *testing.T, tc BrandTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := tc.in.(*braG.DeleteBrandRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	if !tc.wantErr {
		req.ID = id
	}

	got, err := s.DeleteBrand(ctx, req)
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
