package brands

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
	"github.com/iamsabbiralam/restora/server/storage"
)

type Svc struct {
	braG.UnimplementedBrandServiceServer
	cb     CoreBrand
	logger *logrus.Entry
}

type CoreBrand interface {
	CreateBrand(context.Context, storage.Brand) (string, error)
	GetBrandByID(context.Context, string) (*storage.Brand, error)
	UpdateBrand(context.Context, storage.Brand) (*storage.Brand, error)
	ListBrand(context.Context, storage.ListBrandFilter) ([]storage.Brand, error)
	DeleteBrand(context.Context, string, string) error
}

func New(cb CoreBrand, logger *logrus.Entry) *Svc {
	return &Svc{
		cb:     cb,
		logger: logger,
	}
}

// // RegisterService with grpc server.
func (h *Svc) RegisterSvc(srv *grpc.Server) error {
	braG.RegisterBrandServiceServer(srv, h)
	return nil
}

func Permission(ctx context.Context) map[string]storage.ResAct {
	p := map[string]storage.ResAct{
		"CreateBrand":  {Resource: "brand", Action: "Create"},
		"GetBrandByID": {Resource: "brand", Action: "Read", Public: true},
		"UpdateBrand":  {Resource: "brand", Action: "Update"},
		"DeleteBrand":  {Resource: "brand", Action: "Delete"},
		"ListBrand":    {Resource: "brand", Action: "Read"},
	}
	return p
}
