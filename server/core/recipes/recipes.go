package recipes

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/iamsabbiralam/restora/server/storage"
)

type Svc struct {
	store  RecipeStore
	logger *logrus.Entry
}

func New(store RecipeStore, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  store,
		logger: logger,
	}
}

type RecipeStore interface {
	CreateRecipe(context.Context, storage.Recipe) (string, error)
	GetRecipeByID(context.Context, string) (*storage.Recipe, error)
	UpdateRecipe(context.Context, storage.Recipe) (*storage.Recipe, error)
	ListRecipe(context.Context, storage.ListRecipeFilter) ([]storage.Recipe, error)
	DeleteRecipe(context.Context, string, string) error
}
