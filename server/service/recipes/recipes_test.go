package recipes

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/timestamppb"

	recG "github.com/iamsabbiralam/restora/proto/v1/server/recipe"
	recCore "github.com/iamsabbiralam/restora/server/core/recipes"
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
		"service": "test_recipes",
	})

	return st, New(recCore.New(st, logger), logger)
}

func newTestStorage(tb testing.TB) *postgres.Storage {
	if testing.Short() {
		tb.Skip("skipping tests that use postgres on -short")
	}

	return _testStorage
}

type RecipeTestStruct struct {
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

func TestRecipes(t *testing.T) {
	st := newTestStorage(t)
	_, s := newTestSvc(t, st)
	jsonData := `{"onion", "garlic", "ginger"}`
	ing, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Unable to marshal period: %v", err)
	}

	tests := []RecipeTestStruct{
		{
			methodName: "CREATE_RECIPE_SUCCESS",
			desc:       "Success Create Recipe",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.CreateRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.CreateRecipeRequest{
				Title:         "test recipe",
				Description:   "test recipe description",
				Ingredrient:   string(ing),
				Image:         "1.jpg",
				CookingTime:   timestamppb.Now(),
				ServingAmount: 1,
				Status:        recG.Status_Active,
			},
			want: &recG.CreateRecipeResponse{
				ID: id,
			},
		},
		{
			methodName: "CREATE_RECIPE_FAILED",
			desc:       "Failed Create Recipe",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.CreateRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in:         &recG.CreateRecipeRequest{},
			want:       nil,
		},
		{
			methodName: "UPDATE_RECIPE_SUCCESS",
			desc:       "Success Update Recipe",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.UpdateRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.UpdateRecipeRequest{
				ID:            id,
				Title:         "test title update",
				Description:   "test recipe description update",
				Ingredrient:   string(ing),
				Image:         "1.jpg",
				CookingTime:   timestamppb.Now(),
				ServingAmount: 1,
				Status:        recG.Status_Inactive,
			},
			want: &recG.UpdateRecipeResponse{
				ID:            id,
				Title:         "test title update",
				Description:   "test recipe description update",
				Ingredrient:   string(ing),
				Image:         "1.jpg",
				CookingTime:   timestamppb.Now(),
				ServingAmount: 1,
				Status:        recG.Status_Inactive,
			},
		},
		{
			methodName: "UPDATE_RECIPE_FAILED",
			desc:       "Failed Update Recipe",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.UpdateRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in:         &recG.UpdateRecipeRequest{},
			want:       nil,
		},
		{
			methodName: "GET_RECIPE_SUCCESS",
			desc:       "Success Get Recipe",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.GetRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.GetRecipeRequest{
				ID: id,
			},
			want: &recG.GetRecipeResponse{
				ID:            id,
				Title:         "test title update",
				Description:   "test recipe description update",
				Ingredrient:   string(ing),
				Image:         "1.jpg",
				CookingTime:   timestamppb.Now(),
				ServingAmount: 1,
				Status:        recG.Status_Active,
			},
		},
		{
			methodName: "GET_RECIPE_FAILED",
			desc:       "Failed Get Recipe",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.GetRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in:         &recG.GetRecipeRequest{},
			want:       nil,
		},

		{
			methodName: "LIST_RECIPE_SUCCESS",
			desc:       "Success List Recipe",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.ListRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.ListRecipeRequest{
				SortBy:       recG.SortBy_DESC,
				SortByColumn: "",
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &recG.ListRecipeResponse{
				Recipes: []*recG.Recipe{
					{
						Title:  "test recipe",
						Status: recG.Status_Active,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_RECIPE_SUCCESS_BY_COLUMN",
			desc:       "Success List Recipe by SortByColumn",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.ListRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.ListRecipeRequest{
				SortBy:       recG.SortBy_ASC,
				SortByColumn: "",
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &recG.ListRecipeResponse{
				Recipes: []*recG.Recipe{
					{
						Title:  "test recipe",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_RECIPE_SUCCESS_BY_SEARCH",
			desc:       "Success List Recipe by SearchTerm",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.ListRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.ListRecipeRequest{
				SortBy:       recG.SortBy_ASC,
				SortByColumn: "",
				SearchTerm:   "BdNews",
				StartDate:    "",
				EndDate:      "",
			},
			want: &recG.ListRecipeResponse{
				Recipes: []*recG.Recipe{
					{
						Title:  "test recipe",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "LIST_RECIPE_SUCCESS_BY_DATE",
			desc:       "Success List Recipe by StartDate",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.ListRecipeResponse{}, recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.ListRecipeRequest{
				SortBy:       recG.SortBy_ASC,
				SortByColumn: "",
				SearchTerm:   "",
				StartDate:    "",
				EndDate:      "",
			},
			want: &recG.ListRecipeResponse{
				Recipes: []*recG.Recipe{
					{
						Title:  "test recipe",
						Status: 1,
					},
				},
				Total: 1,
			},
		},
		{
			methodName: "DELETE_RECIPE_SUCCESS",
			desc:       "Success Delete Recipe",
			wantErr:    false,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.Recipe{}, timestamppb.Timestamp{})},
			in: &recG.DeleteRecipeRequest{
				ID: id,
			},
		},
		{
			methodName: "DELETE_RECIPE_FAILED",
			desc:       "Failed Delete Recipe",
			wantErr:    true,
			tops:       cmp.Options{cmpopts.IgnoreFields(recG.Recipe{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"), cmpopts.IgnoreUnexported(recG.Recipe{}, timestamppb.Timestamp{})},
			in:         &recG.DeleteRecipeRequest{},
			want:       nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.methodName, func(t *testing.T) {
			switch test.methodName {
			case "CreateBrand":
				CreateRecipeTest(t, test, s)
			case "UpdateRecipe":
				UpdateRecipeTest(t, test, s)
			case "GetRecipe":
				GetRecipeTest(t, test, s)
			case "ListRecipe":
				ListRecipeTest(t, test, s)
			case "DeleteRecipe":
				DeleteRecipeTest(t, test, s)
			}
		})
	}
}

func CreateRecipeTest(t *testing.T, test RecipeTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*recG.CreateRecipeRequest)
	if !ok {
		t.Error("request type conversion error")
	}
	got, err := s.CreateRecipe(ctx, req)
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

func UpdateRecipeTest(t *testing.T, test RecipeTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*recG.UpdateRecipeRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	req.ID = id
	got, err := s.UpdateRecipe(ctx, req)
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

func GetRecipeTest(t *testing.T, test RecipeTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*recG.GetRecipeRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	req.ID = id
	got, err := s.GetRecipe(ctx, req)
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

func ListRecipeTest(t *testing.T, test RecipeTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*recG.ListRecipeRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	got, err := s.ListRecipe(ctx, req)
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

func DeleteRecipeTest(t *testing.T, tc RecipeTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := tc.in.(*recG.DeleteRecipeRequest)
	if !ok {
		t.Error("request type conversion error")
	}

	if !tc.wantErr {
		req.ID = id
	}

	got, err := s.DeleteRecipe(ctx, req)
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
