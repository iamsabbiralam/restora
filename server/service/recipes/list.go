package recipes

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	recG "github.com/iamsabbiralam/restora/proto/v1/server/recipe"
	"github.com/iamsabbiralam/restora/server/storage"
	uErr "github.com/iamsabbiralam/restora/utility/error/error"
)

func (s *Svc) ListRecipe(ctx context.Context, req *recG.ListRecipeRequest) (*recG.ListRecipeResponse, error) {
	log := s.logger.WithField("method", "service.recipe.ListRecipe")
	startDate, endDate, err := s.startDateEndDateRangeCheck(req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return nil, uErr.HandleServiceErr(err)
	}

	recipes, err := s.store.ListRecipe(ctx, storage.ListRecipeFilter{
		SearchTerm:   req.SearchTerm,
		Limit:        req.Limit,
		Offset:       req.Offset,
		Status:       storage.ActiveStatus(req.Status),
		SortBy:       req.GetSortBy().String(),
		SortByColumn: req.GetSortByColumn(),
		StartDate:    startDate,
		EndDate:      endDate,
	})
	if err != nil {
		log.WithError(err).Error("error with getting recipes")
		return nil, uErr.HandleServiceErr(err)
	}

	list := make([]*recG.Recipe, len(recipes))
	if len(recipes) > 0 {
		for i, val := range recipes {
			list[i] = &recG.Recipe{
				ID:               val.ID,
				Title:            val.Title,
				Image:            val.Image,
				Ingredrient:      val.Ingredient,
				Description:      val.Description,
				AuthorSocialLink: val.AuthorSocialLink,
				ServingAmount:    int64(val.ServingAmount),
				CookingTime:      timestamppb.New(val.CookingTime),
				Status:           recG.Status(val.Status),
				CreatedAt:        timestamppb.New(val.CreatedAt),
				UpdatedAt:        timestamppb.New(val.UpdatedAt),
			}
		}
	}

	var total int32
	if len(recipes) > 0 {
		total = int32(recipes[0].Count)
	}

	if list == nil {
		return nil, uErr.HandleServiceErr(errors.New("error with getting recipes"))
	}

	return &recG.ListRecipeResponse{
		Recipes: list,
		Total:   total,
	}, nil
}
