package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/iamsabbiralam/restora/server/storage"
)

func insertTestRecipe(t *testing.T, s *Storage) {
	jsonData := `{"Cinnamon", "spices"}`
	data, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Unable to marshal data: %v", err)
	}

	got, err := s.CreateRecipe(context.TODO(), storage.Recipe{
		Title:            "test recipe",
		Ingredient:       string(data),
		Image:            "default.jpg",
		Description:      "This is description",
		UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
		AuthorSocialLink: "https://github.com/iamsabbiralam",
		ReadCount:        10,
		ServingAmount:    10,
		IsUsed:           1,
		Status:           1,
	})
	if err != nil {
		t.Fatalf("Unable to create recipe: %v", err)
	}

	if got == "" {
		t.Fatalf("Unable to create recipe: %v", err)
	}
}

func deleteTestRecipe(t *testing.T, s *Storage, id, deletedBy string) {
	if err := s.DeleteRecipe(context.TODO(), id, deletedBy); err != nil {
		t.Fatalf("Unable to delete recipe: %v", err)
	}
}

func deleteAllRecipePermanently(t *testing.T, s *Storage) {
	if err := s.DeleteAllRecipePermanently(context.TODO()); err != nil {
		t.Fatalf("Unable to delete all recipe: %v", err)
	}
}

func TestCreateRecipe(t *testing.T) {
	s := newTestStorage(t)
	jsonData := `{"Cinnamon", "spices"}`
	data, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Unable to marshal data: %v", err)
	}

	testCases := []struct {
		name     string
		in       storage.Recipe
		want     string
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage)
	}{
		{
			name: "CREATE_RECIPE_SUCCESS",
			in: storage.Recipe{
				Title:            "test recipe",
				Ingredient:       string(data),
				Image:            "default.jpg",
				Description:      "This is description",
				UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
				AuthorSocialLink: "https://github.com/iamsabbiralam",
				ReadCount:        10,
				ServingAmount:    10,
				IsUsed:           1,
				Status:           1,
			},
			wantErr:  false,
			teardown: deleteAllRecipePermanently,
		}, {
			name: "CREATE_RECIPE_FAIL",
			in: storage.Recipe{
				Title:            "",
				Ingredient:       "test ingredient",
				Image:            "default.jpg",
				Description:      "This is description",
				UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
				AuthorSocialLink: "https://github.com/iamsabbiralam",
				ReadCount:        10,
				ServingAmount:    10,
				IsUsed:           1,
				Status:           1,
			},
			wantErr:  true,
			setup:    insertTestRecipe,
			teardown: deleteAllRecipePermanently,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t, s)
			}

			got, err := s.CreateRecipe(context.TODO(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.CreateRecipe() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && got == "" {
				t.Errorf("Storage.CreateRecipe() want, got %v", got)
				return
			}

			if tc.teardown != nil {
				tc.teardown(t, s)
			}
		})
	}
}

func TestGetRecipeByID(t *testing.T) {
	s := newTestStorage(t)
	jsonData := `{"Cinnamon", "spices"}`
	data, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Unable to marshal data: %v", err)
	}

	var getRecipeByIDTestCase = []struct {
		name     string
		id       string
		want     *storage.Recipe
		wantErr  bool
		teardown func(*testing.T, *Storage, string, string)
		setup    func(*testing.T, *Storage, string)
	}{
		{
			name: "GET_RECIPE_SUCCESS",
			want: &storage.Recipe{
				Title:            "test recipe",
				Ingredient:       string(data),
				Image:            "default.jpg",
				Description:      "This is description",
				UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
				AuthorSocialLink: "https://github.com/iamsabbiralam",
				ReadCount:        10,
				ServingAmount:    10,
				IsUsed:           1,
				Status:           1,
			},
			teardown: deleteTestRecipe,
		},
		{
			name:    "nonExistentRecipeID",
			id:      "nonexistentrecipeid",
			wantErr: true,
		},
	}

	rID, err := s.CreateRecipe(context.TODO(), storage.Recipe{
		Title:            "test recipe",
		Ingredient:       string(data),
		Image:            "default.jpg",
		Description:      "This is description",
		UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
		AuthorSocialLink: "https://github.com/iamsabbiralam",
		ReadCount:        10,
		ServingAmount:    10,
		IsUsed:           1,
		Status:           1,
	})
	if err != nil {
		t.Fatalf("Unable to create recipe: %v", err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Recipe{},
			"ID",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range getRecipeByIDTestCase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.id != "" {
				rID = tc.id
			}

			got, err := s.GetRecipeByID(context.Background(), rID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetRecipeByID() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.GetRecipeByID() diff = %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, rID, "")
			}
		})
	}
}

func TestUpdateRecipeByID(t *testing.T) {
	s := newTestStorage(t)
	jsonData := `{"Cinnamon", "spices"}`
	data, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Unable to marshal data: %v", err)
	}

	var updateRecipeByID = []struct {
		name     string
		in       storage.Recipe
		want     *storage.Recipe
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage, string, string)
	}{
		{
			name: "UPDATE_RECIPE_SUCCESS",
			in: storage.Recipe{
				Title:            "test recipe update",
				Ingredient:       string(data),
				Image:            "1.jpg",
				Description:      "This is description update",
				UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
				AuthorSocialLink: "https://github.com/sabbiralam",
				ReadCount:        20,
				ServingAmount:    20,
				IsUsed:           2,
				Status:           2,
			},
			want: &storage.Recipe{
				Title:            "test recipe update",
				Ingredient:       string(data),
				Image:            "1.jpg",
				Description:      "This is description update",
				UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
				AuthorSocialLink: "https://github.com/sabbiralam",
				ReadCount:        20,
				ServingAmount:    20,
				IsUsed:           2,
				Status:           2,
			},
			setup:    insertTestRecipe,
			teardown: deleteTestRecipe,
		},
	}

	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Recipe{},
			"ID",
			"CreatedAt",
			"CreatedBy",
			"UpdatedAt",
			"UpdatedBy",
			"DeletedAt",
			"DeletedBy",
		),
	}

	for _, tc := range updateRecipeByID {
		t.Run(tc.name, func(t *testing.T) {
			rID, err := s.CreateRecipe(context.TODO(), storage.Recipe{
				Title:            "test recipe",
				Ingredient:       string(data),
				Image:            "default.jpg",
				Description:      "This is description",
				UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
				AuthorSocialLink: "https://github.com/iamsabbiralam",
				ReadCount:        10,
				ServingAmount:    10,
				IsUsed:           1,
				Status:           1,
			})
			if err != nil {
				t.Fatalf("Unable to create recipe: %v", err)
			}

			rec, err := s.GetRecipeByID(context.Background(), rID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.GetRecipeByID() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			tc.in.ID = rec.ID
			got, err := s.UpdateRecipe(context.Background(), tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.UpdateRecipe() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}

			tc.want = got
			if !cmp.Equal(got, tc.want, opts...) {
				t.Errorf("Storage.UpdateRecipe() diff = + got, - want %v", cmp.Diff(got, tc.want, opts...))
			}

			if tc.teardown != nil {
				tc.teardown(t, s, got.ID, "")
			}
		})
	}
}

func TestDeleteRecipe(t *testing.T) {
	s := newTestStorage(t)
	jsonData := `{"Cinnamon", "spices"}`
	data, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Unable to marshal data: %v", err)
	}

	testCases := []struct {
		name     string
		in       storage.Recipe
		wantErr  bool
		setup    func(*testing.T, *Storage)
		teardown func(*testing.T, *Storage, string, string)
	}{
		{
			name: "RECIPE_DELETATION_SUCCESS",
			in: storage.Recipe{
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
			name: "RECIPE_DELETATION_FAIL",
			in: storage.Recipe{
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

	rID, err := s.CreateRecipe(context.TODO(), storage.Recipe{
		Title:            "test recipe",
		Ingredient:       string(data),
		Image:            "default.jpg",
		Description:      "This is description",
		UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
		AuthorSocialLink: "https://github.com/iamsabbiralam",
		ReadCount:        10,
		ServingAmount:    10,
		IsUsed:           1,
		Status:           1,
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

			err := s.DeleteRecipe(context.TODO(), tc.in.ID, "")
			if (err != nil) != tc.wantErr {
				t.Errorf("Storage.DeleteRecipe() gotErr = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestListRecipe(t *testing.T) {
	ts := newTestStorage(t)
	jsonData := `{"Cinnamon", "spices"}`
	data, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Unable to marshal data: %v", err)
	}

	fs := []storage.Recipe{
		{
			Title:            "test recipe",
			Ingredient:       string(data),
			Image:            "default.jpg",
			Description:      "This is description",
			UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
			AuthorSocialLink: "https://github.com/iamsabbiralam",
			ReadCount:        10,
			ServingAmount:    10,
			IsUsed:           1,
			Status:           1,
		},
		{
			Title:            "pasta",
			Ingredient:       string(data),
			Image:            "1.jpg",
			Description:      "This is pasta",
			UserID:           "b6ddbe32-3d7e-4828-b2d7-da9927846e6b",
			AuthorSocialLink: "https://github.com/pasta",
			ReadCount:        10,
			ServingAmount:    10,
			IsUsed:           1,
			Status:           1,
		},
	}
	for _, f := range fs {
		if _, err := ts.CreateRecipe(context.TODO(), f); err != nil {
			t.Error(err)
		}
	}

	testsList := []struct {
		name string
		in   storage.ListRecipeFilter
		want []storage.Recipe
	}{
		{
			name: "No Limit",
			in:   storage.ListRecipeFilter{},
			want: []storage.Recipe{fs[0], fs[1]},
		},
		{
			name: "First Two",
			in: storage.ListRecipeFilter{
				Limit:  2,
				Offset: 0,
			},
			want: []storage.Recipe{fs[0], fs[1]},
		},
		{
			name: "Last One",
			in: storage.ListRecipeFilter{
				Limit:  1,
				Offset: 1,
			},
			want: []storage.Recipe{fs[1]},
		},
	}
	opts := cmp.Options{
		cmpopts.IgnoreFields(
			storage.Recipe{},
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
			got, err := ts.ListRecipe(context.TODO(), test.in)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(test.want, got, opts...) {
				t.Error("(-want +got): ", cmp.Diff(test.want, got, opts...))
			}
		})
	}
}
