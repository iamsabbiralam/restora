package postgres

import (
	"context"
	"encoding/json"
	"testing"

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
