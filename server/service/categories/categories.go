package categories

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	"github.com/iamsabbiralam/restora/server/storage"
)

type Svc struct {
	catG.UnimplementedCategoryServiceServer
	cc     CoreCategory
	logger *logrus.Entry
}

/* func New(cc CoreCategory) *Svc {
	return &Svc{
		cc: cc,
	}
} */

type CoreCategory interface {
	CreateCategory(context.Context, storage.Category) (string, error)
	GetCategoryByID(context.Context, string) (*storage.Category, error)
	UpdateCategory(context.Context, storage.Category) (*storage.Category, error)
	ListCategories(context.Context, storage.ListCategoryFilter) ([]storage.Category, error)
	DeleteCategory(context.Context, string) error
}

func New(cc CoreCategory, logger *logrus.Entry) *Svc {
	return &Svc{
		cc:     cc,
		logger: logger,
	}
}

// // RegisterService with grpc server.
func (h *Svc) RegisterSvc(srv *grpc.Server) error {
	catG.RegisterCategoryServiceServer(srv, h)
	return nil
}

func Permission(ctx context.Context) map[string]storage.ResAct {
	p := map[string]storage.ResAct{
		"CreateCategory":  {Resource: "category", Action: "Create"},
		"GetCategoryByID": {Resource: "category", Action: "Read", Public: true},
		"UpdateCategory":  {Resource: "category", Action: "Update"},
		"DeleteCategory":  {Resource: "category", Action: "Delete"},
		"ListCategories":  {Resource: "category", Action: "Read"},
	}
	return p
}
