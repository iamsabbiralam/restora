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

func insertTestCategory(t *testing.T, s *Storage) {
	got, err := s.CreateCategory(context.TODO(), storage.Category{
		Name:   "test role",
		Status: 1,
	})
	if err != nil {
		t.Fatalf("Unable to create category: %v", err)
	}

	if got == "" {
		t.Fatalf("Unable to create category: %v", err)
	}
}

func deleteTestCategory(t *testing.T, s *Storage, id, deletedBy string) {
	if err := s.DeleteCategory(context.TODO(), id, deletedBy); err != nil {
		t.Fatalf("Unable to delete category: %v", err)
	}
}

func deleteAllTestCategories(t *testing.T, s *Storage) {
	if err := s.deleteCategoriesPermanently(context.TODO()); err != nil {
		t.Fatalf("Unable to delete all categories: %v", err)
	}
}

func TestCreateCategory(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.Category
		want     string
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage)
	}{
		{
			name: "CREATE_CATEGORY_SUCCESS",
			in: storage.Category{
				Name:   "test category",
				Status: 1,
			},
			wantErr:  false,
			teardown: deleteAllTestCategories,
		}, {
			name: "CREATE_CATEGORY_FAIL",
			in: storage.Category{
				Name:   "",
				Status: 1,
			},
			wantErr:  true,
			setup:    insertTestCategory,
			teardown: deleteAllTestCategories,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.CreateCategory(context.TODO(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.CreateCategory() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && got == "" {
				t.Errorf("Storage.CreateCategory() want, got %v", got)
				return
			}

			if tc.teardown != nil {
				tc.teardown(t, s)
			}
		})
	}
}

var getCategoryByIDTestCase = []struct {
	name     string
	id       string
	want     *storage.Category
	wantErr  bool
	teardown func(*testing.T, *Storage, string, string)
	setup    func(*testing.T, *Storage, string)
}{
	{
		name: "GET_CATEGORY_SUCCESS",
		want: &storage.Category{
			Name:   "test category",
			Status: 1,
		},
		teardown: deleteTestCategory,
	},
	{
		name:    "nonExistentCategoryID",
		id:      "nonexistentcategoryid",
		wantErr: true,
	},
}

func TestGetCategoryByID(t *testing.T) {
	s := newTestStorage(t)
	rID, err := s.CreateCategory(context.TODO(), storage.Category{
		Name:   "test category",
		Status: 1,
	})
	if err != nil {
		t.Fatalf("Unable to create category: %v", err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Category{},
			"ID",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range getCategoryByIDTestCase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.id != "" {
				rID = tc.id
			}

			got, err := s.GetCategoryByID(context.Background(), rID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetCategoryByID() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.GetCategoryByID() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, rID, "")
			}
		})
	}
}

var updateCategoryByID = []struct {
	name     string
	in       storage.Category
	want     *storage.Category
	wantErr  bool
	setup    func(*testing.T, *Storage)
	teardown func(*testing.T, *Storage, string, string)
}{
	{
		name: "UPDATE_CATEGORY_SUCCESS",
		in: storage.Category{
			Name:   "test category",
			Status: 1,
		},
		want: &storage.Category{
			Name:   "test update category",
			Status: 2,
		},
		setup:    insertTestCategory,
		teardown: deleteTestCategory,
	},
}

func TestUpdateRoleByID(t *testing.T) {
	s := newTestStorage(t)
	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Category{},
			"ID",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range updateCategoryByID {
		t.Run(tc.name, func(t *testing.T) {
			rID, err := s.CreateCategory(context.TODO(), storage.Category{
				Name:   "test category",
				Status: 1,
			})
			if err != nil {
				t.Fatalf("Unable to create category: %v", err)
			}

			cat, err := s.GetCategoryByID(context.Background(), rID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetCategoryByID() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			tc.in.ID = cat.ID
			got, err := s.UpdateCategory(context.Background(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.UpdateCategory() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			tc.want = got
			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.UpdateCategory() diff = + got, - want %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, got.ID, "")
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	s := newTestStorage(t)
	testCases := []struct {
		name     string
		in       storage.Category
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage, string, string)
	}{
		{
			name: "CATEGORY_DELETATION_SUCCESS",
			in: storage.Category{
				CRUDTimeDate: storage.CRUDTimeDate{
					DeletedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				},
			},
			wantErr:  false,
			teardown: deleteTestCategory,
		}, {
			name: "CATEGORY_DELETATION_FAIL",
			in: storage.Category{
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

	rID, err := s.CreateCategory(context.TODO(), storage.Category{
		Name:   "test category",
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

			err := s.DeleteCategory(context.TODO(), tc.in.ID, "")
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.DeleteCategory() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
