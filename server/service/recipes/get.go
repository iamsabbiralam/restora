package recipes

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	recG "github.com/iamsabbiralam/restora/proto/v1/server/recipe"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) GetRecipe(ctx context.Context, req *recG.GetRecipeRequest) (*recG.GetRecipeResponse, error) {
	log := s.logger.WithField("method", "service.recipes.GetRecipe")
	if req.GetID() == "" {
		return nil, uErr.HandleServiceErr(errors.New("ID is required"))
	}

	r, err := s.store.GetRecipeByID(ctx, req.GetID())
	if err != nil {
		errMsg := "failed to get recipe by id"
		log.WithError(err).Error(errMsg)
		return nil, uErr.HandleServiceErr(err)
	}

	if r == nil {
		log.WithError(err).Error("error while getting recipe by id")
		return nil, uErr.HandleServiceErr(err)
	}

	res := &recG.GetRecipeResponse{
		ID:               r.ID,
		Title:            r.Title,
		Ingredrient:      r.Ingredient,
		Image:            r.Image,
		Description:      r.Description,
		AuthorSocialLink: r.AuthorSocialLink,
		ServingAmount:    int64(r.ServingAmount),
		CookingTime:      timestamppb.New(r.CookingTime),
		Status:           recG.Status(r.Status),
		CreatedAt:        timestamppb.New(r.CreatedAt),
		CreatedBy:        r.CreatedBy,
		UpdatedAt:        timestamppb.New(r.UpdatedAt),
		UpdatedBy:        r.UpdatedBy,
	}

	return res, nil
}
