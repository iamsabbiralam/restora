package recipes

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	recG "github.com/iamsabbiralam/restora/proto/v1/server/recipe"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) UpdateRecipe(ctx context.Context, req *recG.UpdateRecipeRequest) (*recG.UpdateRecipeResponse, error) {
	log := s.logger.WithField("method", "service.recipes.UpdateRecipe")
	if req == nil {
		return nil, uErr.HandleServiceErr(errors.New("request is nil"))
	}

	updateData := storage.Recipe{
		ID:               req.GetID(),
		Title:            req.GetTitle(),
		Ingredient:       req.GetIngredrient(),
		Image:            req.GetImage(),
		Description:      req.GetDescription(),
		AuthorSocialLink: req.GetAuthorSocialLink(),
		ServingAmount:    int32(req.GetServingAmount()),
		CookingTime:      req.CookingTime.AsTime(),
		Status:           int32(req.GetStatus()),
	}
	if err := s.ValidateRequestedData(ctx, updateData, req.GetID()); err != nil {
		log.WithError(err).Error("validation error while updating recipe")
		return nil, uErr.HandleServiceErr(err)
	}

	res, err := s.store.UpdateRecipe(ctx, updateData)
	if err != nil {
		errMsg := "failed to update recipe"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if res == nil {
		return nil, uErr.HandleServiceErr(errors.New("update recipe response is nil"))
	}

	resp := &recG.UpdateRecipeResponse{
		ID:               res.ID,
		Title:            res.Title,
		Ingredrient:      res.Ingredient,
		Image:            res.Image,
		Description:      res.Description,
		AuthorSocialLink: res.AuthorSocialLink,
		ServingAmount:    int64(res.ServingAmount),
		CookingTime:      timestamppb.New(res.CookingTime),
		Status:           recG.Status(res.Status),
		UpdatedAt:        timestamppb.New(res.UpdatedAt),
		UpdatedBy:        res.UpdatedBy,
	}

	return resp, nil
}
