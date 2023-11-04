package recipes

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	recG "github.com/iamsabbiralam/restora/proto/v1/server/recipe"
	"github.com/iamsabbiralam/restora/server/storage"
)

type Svc struct {
	recG.UnimplementedRecipeServiceServer
	store  CoreRecipe
	logger *logrus.Entry
}

type CoreRecipe interface {
	CreateRecipe(context.Context, storage.Recipe) (string, error)
	GetRecipeByID(context.Context, string) (*storage.Recipe, error)
	UpdateRecipe(context.Context, storage.Recipe) (*storage.Recipe, error)
	ListRecipe(context.Context, storage.ListRecipeFilter) ([]storage.Recipe, error)
	DeleteRecipe(context.Context, string, string) error
}

func New(store CoreRecipe, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  store,
		logger: logger,
	}
}

// // RegisterService with grpc server.
func (h *Svc) RegisterSvc(srv *grpc.Server) error {
	recG.RegisterRecipeServiceServer(srv, h)
	return nil
}

func Permission(ctx context.Context) map[string]storage.ResAct {
	p := map[string]storage.ResAct{
		"CreateRecipe":  {Resource: "recipe", Action: "Create"},
		"GetRecipeByID": {Resource: "recipe", Action: "Read", Public: true},
		"UpdateRecipe":  {Resource: "recipe", Action: "Update"},
		"DeleteRecipe":  {Resource: "recipe", Action: "Delete"},
		"ListRecipe":    {Resource: "recipe", Action: "Read"},
	}
	
	return p
}
