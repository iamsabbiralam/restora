package recipes

import (
	"context"
	"errors"

	recG "github.com/iamsabbiralam/restora/proto/v1/server/recipe"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) CreateRecipe(ctx context.Context, req *recG.CreateRecipeRequest) (*recG.CreateRecipeResponse, error) {
	log := s.logger.WithField("method", "Service.recipes.CreateRecipe")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	storeData := storage.Recipe{
		Title:            req.GetTitle(),
		Ingredient:       req.GetIngredrient(),
		Image:            req.GetImage(),
		Description:      req.GetDescription(),
		AuthorSocialLink: req.GetAuthorSocialLink(),
		ServingAmount:    int32(req.GetServingAmount()),
		CookingTime:      req.CookingTime.AsTime(),
		Status:           int32(req.GetStatus()),
	}
	if err := s.ValidateRequestedData(ctx, storeData, ""); err != nil {
		log.WithError(err).Error("validation error while creating recipe")
		return nil, uErr.HandleServiceErr(err)
	}

	res, err := s.store.CreateRecipe(ctx, storeData)
	if err != nil {
		errMsg := "unable to create recipe"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	return &recG.CreateRecipeResponse{
		ID: res,
	}, nil
}
