package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/iamsabbiralam/restora/server/storage"
)

func insertTestBrand(t *testing.T, s *Storage) {
	got, err := s.CreateBrand(context.TODO(), storage.Brand{
		Name:   "test brand",
		Status: 1,
	})
	if err != nil {
		t.Fatalf("Unable to create brand: %v", err)
	}

	if got == "" {
		t.Fatalf("Unable to create brand: %v", err)
	}
}

func deleteTestBrand(t *testing.T, s *Storage, id, deletedBy string) {
	if err := s.DeleteBrand(context.TODO(), id, deletedBy); err != nil {
		t.Fatalf("Unable to delete brand: %v", err)
	}
}

func deleteAllTestBrands(t *testing.T, s *Storage) {
	if err := s.DeleteBrandsPermanently(context.TODO()); err != nil {
		t.Fatalf("Unable to delete all brands: %v", err.Error())
	}
}

func TestCreateBrand(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.Brand
		want     string
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage)
	}{
		{
			name: "CREATE_BRAND_SUCCESS",
			in: storage.Brand{
				Name:   "test brand",
				Status: 1,
			},
			wantErr:  false,
			teardown: deleteAllTestBrands,
		}, {
			name: "CREATE_BRAND_FAIL",
			in: storage.Brand{
				Name:   "",
				Status: 1,
			},
			wantErr:  true,
			setup:    insertTestBrand,
			teardown: deleteAllTestBrands,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.CreateBrand(context.TODO(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.CreateBrand() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && got == "" {
				t.Errorf("Storage.CreateBrand() want, got %v", got)
				return
			}

			if tc.teardown != nil {
				tc.teardown(t, s)
			}
		})
	}
}

var getBrandByIDTestCase = []struct {
	name     string
	id       string
	want     *storage.Brand
	wantErr  bool
	teardown func(*testing.T, *Storage, string, string)
	setup    func(*testing.T, *Storage, string)
}{
	{
		name: "GET_BRAND_SUCCESS",
		want: &storage.Brand{
			Name:   "test brand",
			Status: 1,
		},
		teardown: deleteTestBrand,
	},
	{
		name:    "nonExistentBrandID",
		id:      "nonexistentbrandid",
		wantErr: true,
	},
}

func TestGetBrandByID(t *testing.T) {
	s := newTestStorage(t)
	rID, err := s.CreateBrand(context.TODO(), storage.Brand{
		Name:   "test brand",
		Status: 1,
	})
	if err != nil {
		t.Fatalf("Unable to create brand: %v", err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Brand{},
			"ID",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range getBrandByIDTestCase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.id != "" {
				rID = tc.id
			}

			got, err := s.GetBrandByID(context.Background(), rID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetBrandByID() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.GetBrandByID() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, rID, "")
			}
		})
	}
}

var updateBrandByID = []struct {
	name     string
	in       storage.Brand
	want     *storage.Brand
	wantErr  bool
	setup    func(*testing.T, *Storage)
	teardown func(*testing.T, *Storage, string, string)
}{
	{
		name: "UPDATE_BRAND_SUCCESS",
		in: storage.Brand{
			Name:   "test brand",
			Status: 1,
		},
		want: &storage.Brand{
			Name:   "test update brand",
			Status: 2,
		},
		setup:    insertTestBrand,
		teardown: deleteTestBrand,
	},
}

func TestUpdateBrandByID(t *testing.T) {
	s := newTestStorage(t)
	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Brand{},
			"ID",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range updateBrandByID {
		t.Run(tc.name, func(t *testing.T) {
			rID, err := s.CreateBrand(context.TODO(), storage.Brand{
				Name:   "test brand",
				Status: 1,
			})
			if err != nil {
				t.Fatalf("Unable to create brand: %v", err)
			}

			cat, err := s.GetBrandByID(context.Background(), rID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetBrandByID() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			tc.in.ID = cat.ID
			got, err := s.UpdateBrand(context.Background(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.UpdateBrand() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			tc.want = got
			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.UpdateBrand() diff = + got, - want %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, got.ID, "")
			}
		})
	}
}

func TestDeleteBrand(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.Brand
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage, string, string)
	}{
		{
			name: "BRAND_DELETION_SUCCESS",
			in: storage.Brand{
				CRUDTimeDate: storage.CRUDTimeDate{
					DeletedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				},
			},
			wantErr:  false,
			teardown: deleteTestBrand,
		}, {
			name: "BRAND_DELETION_FAIL",
			in: storage.Brand{
				CRUDTimeDate: storage.CRUDTimeDate{
					DeletedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				},
			},
			wantErr: true,
		},
	}

	rID, err := s.CreateBrand(context.TODO(), storage.Brand{
		Name:   "test brand",
		Status: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	testCases[0].in.ID = rID
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			err := s.DeleteBrand(context.TODO(), tc.in.ID, "")
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.DeleteBrand() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestListBrand(t *testing.T) {
	ts := newTestStorage(t)
	fs := []storage.Brand{
		{
			Name:   "casual",
			Status: 1,
		},
		{
			Name:   "Audi",
			Status: 1,
		},
	}
	for _, f := range fs {
		if _, err := ts.CreateBrand(context.TODO(), f); err != nil {
			t.Error(err)
		}
	}

	testsList := []struct {
		name string
		in   storage.ListBrandFilter
		want []storage.Brand
	}{
		{
			name: "No Limit",
			in:   storage.ListBrandFilter{},
			want: []storage.Brand{fs[0], fs[1]},
		},
		{
			name: "First Two",
			in: storage.ListBrandFilter{
				Limit:  2,
				Offset: 0,
			},
			want: []storage.Brand{fs[0], fs[1]},
		},
		{
			name: "Last One",
			in: storage.ListBrandFilter{
				Limit:  1,
				Offset: 1,
			},
			want: []storage.Brand{fs[1]},
		},
	}
	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Brand{},
			"ID",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
			"Count",
		),
	}
	for _, test := range testsList {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := ts.ListBrand(context.TODO(), test.in)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(test.want, got, opts...) {
				t.Error("(-want +got): ", cmp.Diff(test.want, got, opts...))
			}
		})
	}
}
