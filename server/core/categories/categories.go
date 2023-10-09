package categories

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/iamsabbiralam/restora/server/storage"
)

type Svc struct {
	store  CategoryStore
	logger *logrus.Entry
}

func New(cs CategoryStore, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  cs,
		logger: logger,
	}
}

type CategoryStore interface {
	CreateCategory(context.Context, storage.Category) (string, error)
	GetCategoryByID(context.Context, string) (*storage.Category, error)
	UpdateCategory(context.Context, storage.Category) (*storage.Category, error)
	ListCategories(context.Context, storage.ListCategoryFilter) ([]storage.Category, error)
	DeleteCategory(context.Context, string, string) error
}
