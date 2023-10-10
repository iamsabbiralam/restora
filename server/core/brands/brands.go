package brands

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/iamsabbiralam/restora/server/storage"
)

type Svc struct {
	store  BrandStore
	logger *logrus.Entry
}

func New(store BrandStore, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  store,
		logger: logger,
	}
}

type BrandStore interface {
	CreateBrand(context.Context, storage.Brand) (string, error)
	GetBrandByID(context.Context, string) (*storage.Brand, error)
	UpdateBrand(context.Context, storage.Brand) (*storage.Brand, error)
	ListBrand(context.Context, storage.ListBrandFilter) ([]storage.Brand, error)
	DeleteBrand(context.Context, string, string) error
}
